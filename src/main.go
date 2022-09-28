package main

import (
	"handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	hh := handlers.NewHelloHandler(log.New(os.Stdout, "rest", log.LstdFlags))
	http.ListenAndServe(":9090", nil)
}
