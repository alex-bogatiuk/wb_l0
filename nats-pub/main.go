package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
	"math/rand"
	"time"
)

type Order struct {
	SmID              int       `json:"sm_id" fake:"{number:1,100000}"`
	OrderUID          string    `json:"order_uid" fake:"{uuid}"`
	TrackNumber       string    `json:"track_number" fake:"{uuid}"`
	Entry             string    `json:"entry" fake:"{firstname}"`
	Locale            string    `json:"locale" fake:"{stateabr}"`
	InternalSignature string    `json:"internal_signature" fake:"{state}"`
	CustomerID        string    `json:"customer_id" fake:"{uuid}"`
	DeliveryService   string    `json:"delivery_service" fake:"{word}"`
	Shardkey          string    `json:"shardkey" fake:"{word}"`
	OofShard          string    `json:"oof_shard" fake:"{word}"`
	DateCreated       time.Time `json:"date_created" fake:"{date}"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items" fakesize:"1"`
}
type Delivery struct {
	Name    string `json:"name" fake:"{firstname}"`
	Phone   string `json:"phone" fake:"{phone}"`
	Zip     string `json:"zip" fake:"{zip}"`
	City    string `json:"city" fake:"{city}"`
	Address string `json:"address" fake:"{streetname}"`
	Region  string `json:"region" fake:"{state}"`
	Email   string `json:"email" fake:"{email}"`
}
type Payment struct {
	Amount       int    `json:"amount" fake:"{number:1,1000000}"`
	PaymentDt    int    `json:"payment_dt" fake:"{number:1,100}"`
	DeliveryCost int    `json:"delivery_cost" fake:"{number:1,10}"`
	GoodsTotal   int    `json:"goods_total" fake:"{number:1,10}"`
	CustomFee    int    `json:"custom_fee" fake:"{number:1,100}"`
	Transaction  string `json:"transaction" fake:"{uuid}"`
	RequestID    string `json:"request_id" fake:"{uuid}"`
	Currency     string `json:"currency" fake:"{currencyshort}"`
	Provider     string `json:"provider" fake:"{word}"`
	Bank         string `json:"bank" fake:"{word}"`
}
type Item struct {
	ChrtID      int    `json:"chrt_id" fake:"{number:1,1000000}"`
	Price       int    `json:"price" fake:"{number:1,10000}"`
	Sale        int    `json:"sale" fake:"{number:1,100}"`
	TotalPrice  int    `json:"total_price" fake:"{number:1,1000}"`
	NmID        int    `json:"nm_id" fake:"{number:1,1000000}"`
	Status      int    `json:"status" fake:"{number:0,4}"`
	TrackNumber string `json:"track_number" fake:"{uuid}"`
	Rid         string `json:"rid" fake:"{uuid}"`
	Name        string `json:"name" fake:"{beername}"`
	Size        string `json:"size" fake:"{letter}"`
	Brand       string `json:"brand" fake:"{word}"`
}

func main() {

	// Connecting к NATS Streaming server
	sc, err := stan.Connect("test-cluster", "client-1", stan.NatsURL(nats.DefaultURL))
	if err != nil {
		fmt.Println("Error connecting to NATS Streaming:", err)
		return
	}
	defer sc.Close()

	// Simple Publisher 10m or 600 orders
	for i := 0; i < 600; i++ {

		waiting(1)

		var ord Order
		var itm Item
		// Order gen
		gofakeit.Struct(&ord)

		//item count
		rand.Seed(time.Now().Unix())
		for itmCount := rand.Intn(4); itmCount > 0; itmCount-- {
			gofakeit.Struct(&itm)
			ord.Items = append(ord.Items, itm)
		}

		// add trash data
		wrongDataIndicator := rand.Intn(4)

		if wrongDataIndicator > 1 {
			ordJSON, err := json.Marshal(ord)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = sc.Publish("wb-orders", ordJSON)
			if err != nil {
				fmt.Println("Error publishing:", err)
				return
			}
		} else {
			ordWrongJSON := []byte(fmt.Sprint(ord))

			err = sc.Publish("wb-orders", ordWrongJSON)
			if err != nil {
				fmt.Println("Error publishing:", err)
				return
			}
		}
	}
}

func waiting(sec int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(sec) * time.Second)
}
