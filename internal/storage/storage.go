package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alex-bogatiuk/wb_l0/internal/cache"
	"github.com/alex-bogatiuk/wb_l0/internal/models"
	valid "github.com/alex-bogatiuk/wb_l0/internal/validator"
	repo "github.com/alex-bogatiuk/wb_l0/pkg/repository"
	"github.com/gookit/slog"
	"github.com/jackc/pgx/v5"
)

type DBConn struct {
	db *pgx.Conn
}

//func NewDB(database *pgx.Conn) *DBConn {
//	return &DBConn{db: database}
//}

type OrderStorageService struct {
	cache cache.OrderCacheStorage
	db    DBConn
}

func NewPostgresConn(cfg *repo.Config) (*DBConn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	DBConn := DBConn{}

	var err error
	DBConn.db, err = pgx.Connect(context.Background(), "postgres://"+cfg.Username+":"+cfg.Password+"@"+cfg.Host+"/"+cfg.Basename)
	if err != nil {
		return nil, err
	}

	err = DBConn.db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return &DBConn, nil
}

func OrderStorageInit(cache cache.OrderCacheStorage, db DBConn) *OrderStorageService {
	OrderStorageService := OrderStorageService{
		cache: cache,
		db:    db,
	}
	return &OrderStorageService
}

func (oss *OrderStorageService) SaveOrder(data []byte) error {
	newOrder := new(models.Order)
	err := json.Unmarshal(data, &newOrder)
	if err != nil {
		slog.Error("JSON unmarshal error:", err)
		return err
	}

	err = valid.ValidateOrderStruct(newOrder)
	if err != nil {
		slog.Error("JSON validator error:", err)
		return err
	}

	// Insert
	tx, err := oss.db.db.Begin(context.Background())
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(context.Background())

	// Main order
	_, err = tx.Exec(context.Background(), `insert into orders(order_uid, track_number, entry, locale, 
					internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) 
					VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		newOrder.OrderUID, newOrder.TrackNumber, newOrder.Entry, newOrder.Locale, newOrder.InternalSignature, newOrder.CustomerID,
		newOrder.DeliveryService, newOrder.Shardkey, newOrder.SmID, newOrder.DateCreated, newOrder.OofShard)
	if err != nil {
		return err
	}

	// Delivery
	_, err = tx.Exec(context.Background(), `insert into order_delivery(order_uid, name, phone, zip, city, address, 
                    region, email) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		newOrder.OrderUID, newOrder.Delivery.Name, newOrder.Delivery.Phone, newOrder.Delivery.Zip, newOrder.Delivery.City,
		newOrder.Delivery.Address, newOrder.Delivery.Region, newOrder.Delivery.Email)
	if err != nil {
		return err
	}

	// Payment
	_, err = tx.Exec(context.Background(), `insert into order_payment(order_uid, transaction, request_id, currency, 
					provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
					VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		newOrder.OrderUID, newOrder.Payment.Transaction, newOrder.Payment.RequestID, newOrder.Payment.Currency,
		newOrder.Payment.Provider, newOrder.Payment.Amount, newOrder.Payment.PaymentDt, newOrder.Payment.Bank,
		newOrder.Payment.DeliveryCost, newOrder.Payment.GoodsTotal, newOrder.Payment.CustomFee)
	if err != nil {
		return err
	}

	for _, item := range newOrder.Items {
		_, err = tx.Exec(context.Background(), `insert into order_items(order_uid, chrt_id, track_number, price, 
                        rid, name, sale, size, total_price, nm_id, brand, status) 
						VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9, $10,$11,$12)`,
			newOrder.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice,
			item.NmID, item.Brand, item.Status)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	oss.cache.AddToCache(*newOrder)

	return err
}

//func (dbService *DBService) SaveOrderDB(jsonData *model.DataItem) (sql.Result, error) {
//	result, err := dbService.db.Exec(`insert into orders (id, orderdata) values ($1, $2)`, jsonData.ID, jsonData.OrderData)
//	if err != nil {
//		log.Println("New data in database stored: ", jsonData)
//	}
//	return result, err
//}

func (oss *OrderStorageService) FillOrderStoreCache() error {

	var count int
	err := oss.db.db.QueryRow(context.Background(), "select COUNT(*) FROM orders").Scan(&count)
	if err != nil {
		return err
	}

	rows, err := oss.db.db.Query(context.Background(), `select
			orders.order_uid,
			orders.track_number,
			entry,
			locale,
			internal_signature,
			customer_id,
			delivery_service,
			shardkey,
			sm_id,
			date_created,
			oof_shard,
			order_delivery.name,
			phone,
			zip,
			city,
			address,
			region,
			email,
			chrt_id,
			order_items.track_number,
			price,
			rid,
			order_items.name,
			sale,
			size,
			total_price,
			nm_id,
			brand,
			status,
			transaction,
			request_id,
			currency,
			provider,
			amount,
			payment_dt,
			bank,
			delivery_cost,
			goods_total,
			custom_fee
		FROM orders
		INNER JOIN order_delivery ON orders.order_uid = order_delivery.order_uid
		INNER JOIN order_items ON orders.order_uid = order_items.order_uid
		INNER JOIN order_payment ON orders.order_uid = order_payment.order_uid
		ORDER BY orders.order_uid`)
	// Ensure rows is closed. It is safe to close rows multiple times.
	defer rows.Close()

	// Iterate through the result set
	var orderRow models.Order

	for rows.Next() {
		var orderRowCurrent models.Order
		var itemRow models.Item

		err := rows.Scan(&orderRowCurrent.OrderUID, &orderRowCurrent.TrackNumber, &orderRowCurrent.Entry, &orderRowCurrent.Locale, &orderRowCurrent.InternalSignature, &orderRowCurrent.CustomerID, &orderRowCurrent.DeliveryService, &orderRowCurrent.Shardkey, &orderRowCurrent.SmID, &orderRowCurrent.DateCreated, &orderRowCurrent.OofShard, &orderRowCurrent.Delivery.Name, &orderRowCurrent.Delivery.Phone, &orderRowCurrent.Delivery.Zip, &orderRowCurrent.Delivery.City, &orderRowCurrent.Delivery.Address, &orderRowCurrent.Delivery.Region, &orderRowCurrent.Delivery.Email, &itemRow.ChrtID, &itemRow.TrackNumber, &itemRow.Price, &itemRow.Rid, &itemRow.Name, &itemRow.Sale, &itemRow.Size, &itemRow.TotalPrice, &itemRow.NmID, &itemRow.Brand, &itemRow.Status, &orderRowCurrent.Payment.Transaction, &orderRowCurrent.Payment.RequestID, &orderRowCurrent.Payment.Currency, &orderRowCurrent.Payment.Provider, &orderRowCurrent.Payment.Amount, &orderRowCurrent.Payment.PaymentDt, &orderRowCurrent.Payment.Bank, &orderRowCurrent.Payment.DeliveryCost, &orderRowCurrent.Payment.GoodsTotal, &orderRowCurrent.Payment.CustomFee)
		if err != nil {
			return err
		}

		// First iterate
		if orderRow.OrderUID == "" {
			orderRow = orderRowCurrent
			orderRow.Items = append(orderRow.Items, itemRow)
			continue
		}

		if orderRow.OrderUID != orderRowCurrent.OrderUID {
			oss.cache.AddToCache(orderRow)
			orderRow = orderRowCurrent
			orderRow.Items = append(orderRow.Items, itemRow)
		} else {
			orderRow.Items = append(orderRow.Items, itemRow)
		}

	}

	// Add last row
	oss.cache.AddToCache(orderRow)

	// rows is closed automatically when rows.Next() returns false so it is not necessary to manually close rows.
	// The first error encountered by the original Query call, rows.Next or rows.Scan will be returned here.
	if rows.Err() != nil {
		fmt.Printf("rows error: %v", rows.Err())
		return rows.Err()
	}

	return err

}
