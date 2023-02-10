package rest

import (
	"encoding/json"
	"fmt"
	bu "go-company/middleware"
	"go-company/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

//var val *validator.Validate

type Company struct {
	ID          string
	Name        string `json:"name" validate:"required"`
	Description string
	Amount      uint   `json:"amount" validate:"required,numeric,min=0"`
	Registered  bool   `json:"registered" validate:"required,boolean"`
	Type        string `json:"type" validate:"required,oneof=Corporations NonProfit Cooperative Sole Proprietorship"`
}

type httpError struct {
	Code    uint
	Message string
}

func returnResponse(w http.ResponseWriter, code int, response interface{}) {
	payload, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func CreateCompany(w http.ResponseWriter, req *http.Request) {

	if err := validateJWT(w, req); err != nil {
		//returnResponse(w, http.StatusUnauthorized, httpError{Code: 10, Message: "Auth error"})
		return
	}
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var inputCompany *Company
	if err := dec.Decode(&inputCompany); err != nil {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: "payload validation error"})
		return
	}

	v := validator.New()
	if err := v.Struct(inputCompany); err != nil {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: "payload validation error"})
		return
	}

	companyModel := models.Company{}
	companyModel.Amount = inputCompany.Amount
	companyModel.Description = inputCompany.Description
	companyModel.Name = inputCompany.Name
	companyModel.Registered = inputCompany.Registered
	companyModel.Type = inputCompany.Type

	dbCompany, err := bu.CreateCompany(&companyModel)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: err.Error()})
		return
	}

	returnResponse(w, http.StatusOK, dbCompany)

	return

}

func GetCompany(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryParams := mux.Vars(req)
	fmt.Println("=====2222", queryParams)
	if _, ok := queryParams["name"]; ok == false {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: "payload validation error"})
		return

	}
	fmt.Println("iddd", queryParams)

	//return

	dbCompany, err := bu.GetCompany(queryParams["name"])
	if err != nil {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: err.Error()})
		return
	}

	returnResponse(w, http.StatusOK, dbCompany)

}

func DeleteCompany(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := validateJWT(w, req); err != nil {
		returnResponse(w, http.StatusUnauthorized, httpError{Code: 10, Message: "Auth error"})
		return
	}

	queryParams := mux.Vars(req)
	if _, ok := queryParams["name"]; ok == false {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: "payload validation error"})
		return

	}

	err := bu.DeleteCompany(queryParams["name"])
	if err != nil {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: err.Error()})
		return
	}

	returnResponse(w, http.StatusOK, "Company deleted")

}

func UpdateCompany(w http.ResponseWriter, req *http.Request) {

	if err := validateJWT(w, req); err != nil {
		returnResponse(w, http.StatusUnauthorized, httpError{Code: 10, Message: "Auth error"})
		return
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var inputCompany *Company
	if err := dec.Decode(&inputCompany); err != nil {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: "payload validation error"})
		return
	}

	v := validator.New()
	if err := v.Struct(inputCompany); err != nil {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: "payload validation error"})
		return
	}

	companyModel := models.Company{}
	companyModel.Amount = inputCompany.Amount
	companyModel.Description = inputCompany.Description
	companyModel.Name = inputCompany.Name
	companyModel.Registered = inputCompany.Registered
	companyModel.Type = inputCompany.Type

	dbCompany, err := bu.UpadteCompany(&companyModel)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, httpError{Code: 10, Message: err.Error()})
		return
	}

	returnResponse(w, http.StatusOK, dbCompany)

	return

}
