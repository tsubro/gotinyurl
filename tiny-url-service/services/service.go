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

func EncodeUrl(req *models.Request) (*models.Response, error) {

	hash := b64.StdEncoding.EncodeToString([]byte(strconv.Itoa(req.UserId) + "#" + req.Url))
	hash = hash[len(hash)-7:]
	tinyUrl := base_uri + hash

	crTime := time.Now().Unix()
	exTime := time.Now().Add(24).Unix()
	u := models.UrlModel{Hash: hash, OriginalUrl: req.Url, CreationTime: crTime, ExpiryTime: exTime, UserId: req.UserId}

	err := dao.Insert(&u)
	if err != nil {
		return nil, err
	}
	res := models.Response{TinyUrl: tinyUrl, Url : req.Url, ExpiryTime: exTime}

	return &res, nil
}

func ServeRequest(hash string, w *http.ResponseWriter, r *http.Request) error {

	res, err := dao.Get(hash)
	if err != nil {
		return  err
	}

	http.Redirect(*w, r, res.OriginalUrl, http.StatusMovedPermanently)
	return nil
}

func Info(hash string) (*models.Response, error) {
	res, err := dao.Get(hash)

	if err != nil {
		return  nil, err
	}

	tinyUrl := base_uri + hash
	return &models.Response{TinyUrl: tinyUrl, Url: res.OriginalUrl, ExpiryTime : res.ExpiryTime}, nil
}
