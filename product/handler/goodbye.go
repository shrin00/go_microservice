package handler

import (
	"fmt"
	"log"
	"net/http"
)

type GoodbyeHandler struct {
	l *log.Logger
}

func NewGoodbyeHandler(l *log.Logger) *GoodbyeHandler {
	return &GoodbyeHandler{l: l}
}

func (h *GoodbyeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Goodbye World!!")
	fmt.Fprintf(w, "Goodbye World!!\n")
}
