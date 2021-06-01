package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/roaires/fullcycle-arquitetura-hexagonal-golang/adapters/web/handler"
	"github.com/roaires/fullcycle-arquitetura-hexagonal-golang/application"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Server() {
	//Ajudar no tratamento das rotas
	r := mux.NewRouter()

	//Midware de log
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	//Configuração do web server
	server := &http.Server{
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
