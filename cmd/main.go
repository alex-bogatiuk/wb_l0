package main

import (
	"encoding/json"
	"fmt"
	"github.com/alex-bogatiuk/wb_l0/internal/models"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/slog"
	"github.com/nats-io/nats.go"
	"time"
)

var goodJSON = `
{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
`

var Validator = validator.New()

func main() {

	var c models.Order

	err := json.Unmarshal([]byte(goodJSON), &c)
	if err != nil {
		slog.Info("json UnmarshalError:", err)
	}

	err = validateStruct(c)
	if err != nil {
		slog.Info(err)
	}

	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	// Simple Async Subscriber
	c.Subscribe("foo", func(s string) {
		fmt.Printf("Received a message: %s\n", s)
	})

	// Simple Publisher
	c.Publish("foo", "Hello World")

	// EncodedConn can Publish any raw Go type using the registered Encoder
	type person struct {
		Name    string
		Address string
		Age     int
	}

	// Go type Subscriber
	c.Subscribe("hello", func(p *person) {
		fmt.Printf("Received a person: %+v\n", p)
	})

	me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street, San Francisco, CA"}

	// Go type Publisher
	c.Publish("hello", me)

	// Unsubscribe
	sub, err := c.Subscribe("foo", nil)
	// ...
	sub.Unsubscribe()

	// Requests
	var response string
	err = c.Request("help", "help me", &response, 10*time.Millisecond)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	}

	// Replying
	c.Subscribe("help", func(subj, reply string, msg string) {
		c.Publish(reply, "I can help!")
	})

	// Close connection
	c.Close()

	/*	nc, _ := nats.Connect(nats.DefaultURL)

		defer nc.Drain()

		nc.Publish("greet.joe", []byte("hello"))

		sub, _ := nc.SubscribeSync("greet.*")

		msg, _ := sub.NextMsg(10 * time.Millisecond)
		fmt.Println("subscribed after a publish...")
		fmt.Printf("msg is nil? %v\n", msg == nil)

		nc.Publish("greet.joe", []byte("hello"))
		nc.Publish("greet.pam", []byte("hello"))

		msg, _ = sub.NextMsg(10 * time.Millisecond)
		fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

		msg, _ = sub.NextMsg(10 * time.Millisecond)
		fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

		nc.Publish("greet.bob", []byte("hello"))

		msg, _ = sub.NextMsg(10 * time.Millisecond)
		fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)*/
}

func validateStruct(c models.Order) error {

	err := Validator.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
