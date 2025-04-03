package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AdemarTellecher/pos-web-go/core/services"
	"github.com/AdemarTellecher/pos-web-go/dbase"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/negroni"
)

func main() {
	db, err := dbase.DbConnection()

	service := services.NewService(db)

	r := mux.NewRouter()

	//Hendlers
	n := negroni.New(
		negroni.NewLogger(),
	)
	defer db.Close()

	r.Handle("/v1/beer", n.With(
		negroni.Wrap(hello(service)),
	)).Methods("GET", "OPTIONS")
	http.Handle("/", r)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":4000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
func hello(service services.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		all, _ := service.GetAll()
		for _, i := range all {
			fmt.Println(i)
		}
	})
}
