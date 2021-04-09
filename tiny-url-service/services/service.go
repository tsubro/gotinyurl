package services

import (
	b64 "encoding/base64"
	"net/http"
	"strconv"
	"time"
	"tiny-url/dao"
	"tiny-url/models"
)

const base_uri = "http://localhost:8080/tiny/"

func GenerateTinyUrl(req *models.TinyUrlRequest) (string, error) {

	hash := b64.StdEncoding.EncodeToString([]byte(strconv.Itoa(req.UserId) + "#" + req.Url))
	hash = hash[len(hash)-7:]
	hashUrl := base_uri + hash

	dt := time.Now().Unix()
	ex_dt := time.Now().Add(24).Unix()
	u := models.UrlModel{Hash: hash, OriginalUrl: req.Url, CreationDate: dt, ExpiryDate: ex_dt, UserId: req.UserId}

	err := dao.Insert(&u)
	return hashUrl, err
}

func ServeRequest(hash string, req *models.ServeRequest, w *http.ResponseWriter, r *http.Request) {

	res, _ := dao.Get(hash)
	http.Redirect(*w, r, res.OriginalUrl, http.StatusSeeOther)

}
