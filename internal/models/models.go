package models

import "time"

type Order struct {
	SmID              int       `json:"sm_id"`
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	OofShard          string    `json:"oof_shard"`
	DateCreated       time.Time `json:"date_created"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
}
type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}
type Payment struct {
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Bank         string `json:"bank"`
}
type Item struct {
	ChrtID      int    `json:"chrt_id"`
	Price       int    `json:"price"`
	Sale        int    `json:"sale"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Status      int    `json:"status"`
	TrackNumber string `json:"track_number"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Size        string `json:"size"`
	Brand       string `json:"brand"`
}
