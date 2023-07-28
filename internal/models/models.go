package models

import "time"

type Order struct {
	SmID              int       `json:"sm_id" validate:"required"`
	OrderUID          string    `json:"order_uid" validate:"required"`
	TrackNumber       string    `json:"track_number" validate:"required"`
	Entry             string    `json:"entry" validate:"required"`
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required"`
}
type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required"`
}
type Payment struct {
	Amount       int    `json:"amount" validate:"required"`
	PaymentDt    int    `json:"payment_dt" validate:"required" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"required"`
	GoodsTotal   int    `json:"goods_total" validate:"required"`
	CustomFee    int    `json:"custom_fee"`
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
}
type Item struct {
	ChrtID      int    `json:"chrt_id" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Sale        int    `json:"sale" validate:"required"`
	TotalPrice  int    `json:"total_price" validate:"required"`
	NmID        int    `json:"nm_id" validate:"required"`
	Status      int    `json:"status" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Rid         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Size        string `json:"size" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
}
