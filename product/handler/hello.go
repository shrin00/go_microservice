package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


type HelloHandler struct {
	l *log.Logger
}


func NewHelloHandler(l *log.Logger) *HelloHandler {
	return &HelloHandler{l: l}
}


func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World!!")
	// read the data
	d, er := io.ReadAll(r.Body)
	if er != nil { 
		h.l.Printf("failed to read the request data: %s", er.Error())
		http.Error(w, "Oooops", http.StatusBadRequest)
		return
	}

	h.l.Printf("Data: %s\n", d)
	fmt.Fprintf(w, "Data: %s\n", d)
}