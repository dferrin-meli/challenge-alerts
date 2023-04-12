// Package application contains settings for running the api
package application

import (
	"challenge/alerts/src/api/application/conf"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Server type
type Server struct {
	*gin.Engine
	URLPrefix               string
	MiddleEndPrefix         string
	EfficiencyMonitorPrefix string
}

// NewServer Returns gin engine instance
func NewServer(g *gin.Engine) *Server {
	return &Server{
		Engine:    g,
		URLPrefix: "challenge",
	}
}

// Run server
func (r *Server) Run(cfg *conf.Data, options ...interface{}) {
	r.RedirectTrailingSlash = true
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}

	go func() {
		scopeData := " as PRODUCTION scope."
		if cfg.TestScope {
			scopeData = " as TEST scope."
		}
		fmt.Printf("[package:config][function:Server.Run] Running server at port : %s %s", cfg.Port, scopeData)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Print("[package:config][function:Server.Run] Quit server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("[package:config][function:Server.Run] Shutdown server: %s", err)
	}
}
