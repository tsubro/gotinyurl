package handlers

import (
	"encoding/json"
	"net/http"
	"tiny-url/models"
	"tiny-url/services"
	"tiny-url/utils"
	"github.com/gorilla/mux"
)

func EncodeUrl(w http.ResponseWriter, r *http.Request) {

	d := json.NewDecoder(r.Body)
	req := models.Request{}
	err := d.Decode(&req)

	if err != nil {
		sendErrorResponse(&w, 400, "Request Not In Proper Format")
		return
	}

	err = utils.ValidateInput(&req) 
	if err != nil {
		sendErrorResponse(&w, 400, err.Error())
		return
	}

	res, err := services.EncodeUrl(&req)
	

	js, _ := json.Marshal(res)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(201)
	w.Write(js)
}

func ServeRequest(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	services.ServeRequest(params["hash"], &w, r)
}

func Info(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	res,err := services.Info(params["hash"])

	if err !=  nil {
		sendErrorResponse(&w, 400, err.Error())
		return
	}

	js, _ := json.Marshal(res)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
}

func sendErrorResponse(w *http.ResponseWriter, httpStatus int, message string) {

	e := `{"message" : "`+message+`"}`
	(*w).Header().Add("content-type", "application/json")
	(*w).WriteHeader(httpStatus)
	(*w).Write([]byte(e))
}

