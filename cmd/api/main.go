package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sonymuhamad/todo-app/pkg"
	"github.com/sonymuhamad/todo-app/providerwire"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := providerwire.InitializeHttpServer()
	cfg := pkg.LoadEnvConfig()

	go func() {
		err := server.Start(fmt.Sprintf(":%d", cfg.ServerPort))
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			server.Logger.Fatal(err)
		}
	}()

	<-ctx.Done()
	// terminating, graceful shutdown
	log.Println("Shutting Down HTTP Server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
