package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type ServerHandler struct{}

func (sh *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server request entry point\n"))
}

func StartServer() {

	var msh *ServerHandler

	server := &http.Server{
		Addr:    ":5555",
		Handler: msh,
	}

	// Channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Run server in a goroutine
	go func() {
		fmt.Println("Starting server on : " + server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe error: %v\n", err)
		}
	}()

	// Block until a signal is received
	<-stop
	fmt.Println("\nShutting down server...")

	// cleanup process can be done now

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
		return
	}

	fmt.Println("Server exited cleanly")

}
