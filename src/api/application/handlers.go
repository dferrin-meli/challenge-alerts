package application

import (
	"challenge/alerts/src/api/application/common"
	"challenge/alerts/src/api/application/conf"
	"challenge/alerts/src/api/application/registry"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (r *Server) AddHandlers(cfg *conf.Data) *Server {
	alertsHandler := registry.CreateAlertsHandler(cfg)
	// Ping endpoint
	r.GET("/ping", ping)

	//Alerts
	r.GET(fmt.Sprint(r.URLPrefix, "/alerts"), AdaptHandler(alertsHandler.GetAll))
	r.POST(fmt.Sprint(r.URLPrefix, "/alerts"), AdaptHandler(alertsHandler.Create))
	r.GET(fmt.Sprint(r.URLPrefix, "/alerts/search"), AdaptHandler(alertsHandler.Search))
	return r
}

func AdaptHandler(handler func(c *gin.Context) common.ApiError) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			fmt.Printf("Error process request. Error: %v", err.Cause()...)
			c.JSON(err.Status(), err)
		}
	}
}
