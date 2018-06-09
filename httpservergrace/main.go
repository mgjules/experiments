// Http server with graceful exit
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler(),
		ReadTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill, syscall.SIGQUIT)

		<-sigquit
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("Could not stop server: %v", err)
		}
		log.Println("Exited gracefully")
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v", err)
	}
}

func handler() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/", indexHandler)
	return h
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
