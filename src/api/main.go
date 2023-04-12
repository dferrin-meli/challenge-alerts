package main

import (
	"challenge/alerts/src/api/application"
	"challenge/alerts/src/api/application/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := conf.GetData()
	application.NewServer(gin.Default()).AddHandlers(cfg).Run(cfg)
}
