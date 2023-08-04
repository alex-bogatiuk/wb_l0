package cache

import (
	"github.com/alex-bogatiuk/wb_l0/internal/models"
)

type OrderCache map[string]models.Order

type OrderCacheStorage struct {
	cache OrderCache
}

func OrderCacheInit() *OrderCacheStorage {
	cache := make(OrderCache)
	OrderCacheStorage := OrderCacheStorage{
		cache: cache,
	}
	return &OrderCacheStorage
}

func (OrderCacheStorage *OrderCacheStorage) AddToCache(data models.Order) {
	OrderCacheStorage.cache[data.OrderUID] = data
}

func (OrderCacheStorage *OrderCacheStorage) GetOrderFromCache(orderUID string) models.Order {
	return OrderCacheStorage.cache[orderUID]
}
