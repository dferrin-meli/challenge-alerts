package main

import (
	"challenge/alerts/src/api/application"
	"challenge/alerts/src/api/application/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.LoadYMLConfiguration()
	cfg := conf.GetData()
	application.NewServer(gin.Default()).AddHandlers(cfg).Run(cfg)
}
