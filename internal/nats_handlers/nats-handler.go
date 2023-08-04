package natsHandlers

import (
	"encoding/json"
	mdl "github.com/alex-bogatiuk/wb_l0/internal/models"
	"github.com/gookit/slog"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/stan.go"
)

type NatsHandlerService struct {
	MsgHandler func(msg *stan.Msg)
	Conn       *pgx.Conn
}

//var NHS NatsHandlerService

func NatsHandlerServiceInit(Conn *pgx.Conn) *NatsHandlerService {
	return &NatsHandlerService{MsgHandler: func(msg *stan.Msg) {}, Conn: Conn}
}

func CreateOrder(msg *stan.Msg) {
	slog.Info("Received message:", string(msg.Data))

	var mod mdl.Order

	err := json.Unmarshal(msg.Data, &mod)
	if err != nil {
		slog.Error("JSON unmarshal error:", err)
	}

	//err = valid.ValidateStruct(mod)
	//if err != nil {
	//	slog.Error("JSON validator error:", err)
	//}
	//
	//// Check
	//var isNull string
	//err = Conn.QueryRow(context.Background(), "select order_uid from orders").Scan(&isNull)
	//if err != nil {
	//	slog.Info("Check select from base error:", err)
	//}
	//
	//fmt.Println(isNull)
	/////////////////////////////////////////////////
	//rows, err := NHS.Conn.Query(context.Background(), "select order_uid from orders where order_uid = $1", "b563feb7b2b84b6test")
	//
	//// Ensure rows is closed. It is safe to close rows multiple times.
	//defer rows.Close()
	//
	//// Iterate through the result set
	//if rows.Next() {
	//	fmt.Println("dfdfdf")
	//}
	//
	//// rows is closed automatically when rows.Next() returns false so it is not necessary to manually close rows.
	//
	//// The first error encountered by the original Query call, rows.Next or rows.Scan will be returned here.
	//if rows.Err() != nil {
	//	fmt.Printf("rows error: %v", rows.Err())
	//	return
	//}

}
