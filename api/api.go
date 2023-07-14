package api

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const gracefullShutDownTimeout = 5 * time.Second

type Api struct {
	server *http.Server
}

type ApiConfig struct {
	Port string
	Host string
}

func New(routes http.Handler, config *ApiConfig) *Api {

	server := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: routes,
	}

	return &Api{
		server: server,
	}
}

func (api *Api) Start() {
	go func() {
		// service connections
		if err := api.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), gracefullShutDownTimeout)
	defer cancel()

	if err := api.server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server gracefully stopped")

}
