package handlers

import(
	"net/http"
	"encoding/json"
	"tiny-url/models"
	"tiny-url/utils"
)

func GenerateTinyUrl(w http.ResponseWriter, r *http.Request) {

	d := json.NewDecoder(r.Body)
	req := models.TinyUrlRequest{}

	err := d.Decode(&req)

	if err != nil {
		e,_ := json.Marshal(models.TinyUrlError{Message : "Request Not In Proper Format"})
		w.WriteHeader(400)
		w.Write(e)
		return
	}

	err = utils.ValidateInput(&req) 
	if err != nil {
		e,_ := json.Marshal(models.TinyUrlError{Message : err.Error()})
		w.WriteHeader(400)
		w.Write(e)
		return
	}

	res := models.TinyUrlResponse{}
	res.TinyUrl = req.Url

	js, _ := json.Marshal(res)

	w.WriteHeader(201)
	w.Write(js)
}