package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHelloHandler(log *log.Logger) *Hello {
	return &Hello{logger: log}
}
func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.logger.Println("Hello Rest")

	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello Rest %s", data)
}
