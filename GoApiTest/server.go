package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kkamdooong/go-restful-api-example/handler"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/companies").HandlerFunc(handler.GetCompanies)
	sub.Methods("POST").Path("/companies").HandlerFunc(handler.SaveCompany)
	sub.Methods("GET").Path("/companies/{name}").HandlerFunc(handler.GetCompany)
	sub.Methods("PUT").Path("/companies/{name}").HandlerFunc(handler.UpdateCompany)
	sub.Methods("DELETE").Path("/companies/{name}").HandlerFunc(handler.DeleteCompany)
	sub.Methods("GET").Path("/financings").HandlerFunc(handler.GetFinancings)
	sub.Methods("POST").Path("/financings").HandlerFunc(handler.SaveFinancings)

	log.Fatal(http.ListenAndServe(":3000", router))
}
