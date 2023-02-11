package test

import (
	"bytes"
	"encoding/json"
	kafkaconfig "go-company/kafkaConfig"
	"go-company/models"
	"go-company/rest"

	"net/http"
	"net/http/httptest"
	"os"

	//"strings"

	"testing"

	"github.com/gorilla/mux"
)

// var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()

	writer = httptest.NewRecorder()
	code := m.Run()
	//TestHandleGetAuthToken(&testing.T{})

	os.Exit(code)
}

func setUp() {
	models.SetupDatabase()
	kafkaconfig.SetupKafkaProducer()

}

func TestHandleGet(t *testing.T) {

	request, _ := http.NewRequest("GET", "/company/", nil)
	request = mux.SetURLVars(request, map[string]string{"name": "a294"})

	w := httptest.NewRecorder()
	rest.GetCompany(w, request)
	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}

	var m models.Company
	json.NewDecoder(w.Result().Body).Decode(&m)
	if m.Name == "" {
		t.Errorf("Failed to get company details")
	}
}

func TestHandleGetInvalidParam(t *testing.T) {

	request, _ := http.NewRequest("GET", "/company/", nil)
	request = mux.SetURLVars(request, map[string]string{"names": "a294"})

	w := httptest.NewRecorder()
	rest.GetCompany(w, request)
	if w.Code != 400 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestHandleGetIncorrectUrl(t *testing.T) {

	request, _ := http.NewRequest("GET", "/company/name", nil)
	w := httptest.NewRecorder()
	rest.GetCompany(w, request)
	if w.Code != 400 {
		t.Errorf("Response code is %v", w.Code)
	}
}

var Token string

func TestHandleGetAuthToken(t *testing.T) {

	request, _ := http.NewRequest("GET", "/getJWT", nil)
	request.Header.Set("Access", rest.Key)
	//request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	rest.GetJWT(w, request)
	if w.Code != 200 {
		t.Fatalf("Response code is %v", w.Code)
	}

	//err := json.NewDecoder(w.Result().Body).Decode(&m)

	Token = string(w.Body.Bytes())
	if len(Token) == 0 {
		t.Fatal("empty token")
	}

}

func TestHandleCreate(t *testing.T) {
	TestHandleGetAuthToken(t)
	payload := []byte(`{
		"name"  : "ap4",
		"amount" : 113,
		"registered": true  ,
		"type" : "NonProfit",
		"Description": "test"}`)

	request, _ := http.NewRequest("POST", "/company/", bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("token", Token)
	w := httptest.NewRecorder()
	rest.CreateCompany(w, request)
	if w.Code != 200 {
		t.Fatal("Create failed")
	}
	var m models.Company
	json.NewDecoder(w.Result().Body).Decode(&m)
	if m.Name == "" {
		t.Errorf("Failed to get company details")
	}
}
