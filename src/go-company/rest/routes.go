package rest

import (
	//"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/getJWT",GetJWT).Methods("GET")
	router.HandleFunc("/company/", GetCompany).Queries("name", "{name}").Methods("GET")

	//router.Path("/company/").Queries("name", "{name}").HandlerFunc(GetCompany).Methods("GET")
	router.HandleFunc("/company/", CreateCompany).Methods("POST")
	router.HandleFunc("/company/", DeleteCompany).Queries("name", "{name}").Methods("DELETE")
	router.HandleFunc("/company/", UpdateCompany).Methods("PATCH")

	return router

}
