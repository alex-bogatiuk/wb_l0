package main

import (
	"github.com/alex-bogatiuk/wb_l0/internal/cache"
	"github.com/alex-bogatiuk/wb_l0/internal/storage"
	"github.com/alex-bogatiuk/wb_l0/internal/sub"
	repo "github.com/alex-bogatiuk/wb_l0/pkg/repository"
	"github.com/gookit/slog"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"math/rand"
	"time"
)

func main() {

	// Read configuration file. Port, host, basename, username, password... etc
	cfg, err := repo.InitConfig()
	if err != nil {
		slog.Fatal("error initializing cfg file:", err)
		panic(err)
	}

	// New SQL conn
	DBConn, err := storage.NewPostgresConn(cfg)
	if err != nil {
		slog.Fatal("failed initialize db:", err)
		panic(err)
	}

	// Cache
	OrderCache := cache.OrderCacheInit()
	OrderStoreService := storage.OrderStorageInit(*OrderCache, *DBConn)
	err = OrderStoreService.FillOrderStoreCache()
	if err != nil {
		slog.Fatal("restoring cache error:", err)
		panic(err)
	}

	// Connecting to NATS Streaming serv
	natsSub := sub.CreateSub(*OrderStoreService)
	err = natsSub.Connect("test-cluster", "client-2", nats.DefaultURL)
	//natsSub.nc, err := stan.Connect("test-cluster", "client-2", stan.NatsURL(nats.DefaultURL))
	if err != nil {
		slog.Fatal("error connecting to NATS Streaming:", err)
		panic(err)
	}
	//defer natsSub.Close()

	_, err = natsSub.Subscribe("wb-orders", stan.StartWithLastReceived())

	// Here API service

	select {}
	////////////////////
	// Unsubscribe
	//sub.Unsubscribe()
	// Close connection
	//sc.Close()
}

func waiting(sec int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(sec) * time.Second)
}
