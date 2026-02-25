package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shrin00/go_microservice/handler"
)

func main() {
	l := log.New(log.Writer(), "HelloHandler: ", log.LstdFlags)
	ph := handler.NewProduct(l)

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })
	// http.Handle("/hello", hw)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:    ":9090",
		Handler: sm,

		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// http.ListenAndServe(":9090", sm)

	go func() {
		log.Printf("Starting the server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChaan := make(chan os.Signal, 1)
	signal.Notify(sigChaan, os.Interrupt)
	signal.Notify(sigChaan, syscall.SIGTERM)

	sig := <-sigChaan
	log.Printf("Received a signal: %s", sig.String())

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
