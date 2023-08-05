package cache

import "github.com/alex-bogatiuk/wb_l0/internal/models"

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

func (orderCacheStorage *OrderCacheStorage) AddToCache(data models.Order) {
	orderCacheStorage.cache[data.OrderUID] = data
}

func (orderCacheStorage *OrderCacheStorage) GetOrderFromCache(orderUID string) models.Order {
	return orderCacheStorage.cache[orderUID]
}
