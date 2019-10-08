package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	h "../../api/handlers"
	mdb "../../api/mdb"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	r := mux.NewRouter()
	c := mdb.Connect()

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err := c.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("No database connection.\n", err)
	}

	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/c", h.CasesHandler(c)).Methods("GET")
	api.HandleFunc("/c", h.CasesHandlerPost(c)).Methods("POST")
	api.HandleFunc("/cu", h.CasesByUser(c)).Methods("POST")
	api.HandleFunc("/e", h.EqHandler(c)).Methods("GET")
	api.HandleFunc("/h", h.HistoricHandler(c)).Methods("GET")
	api.HandleFunc("/u", h.UsersHandler(c)).Methods("GET")
	api.HandleFunc("/u", h.UsersHandlerPost(c)).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/v/dist/static/"))))
	r.PathPrefix("/").HandlerFunc(h.IndexHandler)

	http.ListenAndServe(":81", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Content-Length", "Origin"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.MaxAge(600))(r))

}
