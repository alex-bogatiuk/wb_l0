package handler

import (
	"github.com/alex-bogatiuk/wb_l0/internal/cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	cache *cache.OrderCacheStorage
}

func NewHandler(cache *cache.OrderCacheStorage) *Handler {
	return &Handler{cache: cache}
}

//func NewHandler(db *storage.DBConn) *Handler {
//	return &Handler{db: db}
//}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	orders := router.Group("/api/getOrderByID")
	{
		// Player page(id needed)
		orders.GET("/:id", h.getOrderByID)
	}

	return router
}

func (h *Handler) getOrderByID(c *gin.Context) {
	c.JSON(http.StatusOK, h.cache.GetOrderFromCache(c.Params[0].Value))
}
