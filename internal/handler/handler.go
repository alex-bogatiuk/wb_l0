package handler

import (
	"encoding/json"
	"github.com/alex-bogatiuk/wb_l0/internal/cache"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"net/http"
	"strings"
)

type Handler struct {
	cache *cache.OrderCacheStorage
}

func NewHandler(cache *cache.OrderCacheStorage) *Handler {
	return &Handler{cache: cache}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("internal/html/getOrderIDForm.tmpl")

	orders := router.Group("/getOrder")
	{
		// Player page(id needed)
		orders.GET("/", h.getOrderByID)
	}

	return router
}

func (h *Handler) getOrderByID(c *gin.Context) {

	if strings.Contains(c.Request.RequestURI, "=") {
		requestURIArr := strings.Split(c.Request.RequestURI, "=")
		orderID := requestURIArr[1]
		order := h.cache.GetOrderFromCache(orderID)

		if order.OrderUID != "" {
			orderJSON, err := json.Marshal(order)
			if err != nil {
				slog.Error("json Marshal error:", err)
				return
			}

			c.HTML(http.StatusOK, "getOrderIDForm.tmpl", gin.H{
				"orderID":    orderID,
				"jsonAnswer": string(orderJSON),
			})

			return
		}

		c.HTML(http.StatusOK, "getOrderIDForm.tmpl", gin.H{
			"orderID":    orderID,
			"jsonAnswer": "Заказ с таким ID не найден",
		})
	} else {
		c.HTML(http.StatusOK, "getOrderIDForm.tmpl", gin.H{
			"orderID":    "",
			"jsonAnswer": "",
		})
	}

}

func (h *Handler) getOrder(c *gin.Context) {
	c.JSON(http.StatusOK, h.cache.GetOrderFromCache(c.Params[0].Value))
}
