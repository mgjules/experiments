// Simple http server with context implemented
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jobc := make(chan struct{})
	defer close(jobc)

	go doStuff(ctx, jobc) // Do some heavy lifting

	select {
	case <-jobc:
		fmt.Fprintln(w, "Hello World")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func doStuff(ctx context.Context, jobc chan<- struct{}) {
	select {
	case <-time.After(5 * time.Second): // Simulating heavy lifting
		jobc <- struct{}{} // Signal that we are done
	case <-ctx.Done():
		log.Print("doStuff stopped")
		return
	}
}
