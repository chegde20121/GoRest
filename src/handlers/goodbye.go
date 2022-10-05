package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	logger *log.Logger
}

func NewGoodByeHandler(log *log.Logger) *GoodBye {
	return &GoodBye{logger: log}
}

func (gb *GoodBye) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	gb.logger.Println("GoodBye Handler")
	rw.Write([]byte("Good Bye GoGeek"))
}
