package services

import(
	b64 "encoding/base64"
	"tiny-url/models"
	"tiny-url/dao"
	"time"
	"strconv"
)

func GenerateTinyUrl(req *models.TinyUrlRequest) error {

	hashUrl := b64.StdEncoding.EncodeToString([]byte(strconv.Itoa(req.UserId)+"#"+req.Url))

	dt := time.Now().Unix()
	u := models.UrlModel{HashedUrl : hashUrl, OriginalUrl: req.Url, CreationDate : dt, ExpiryDate : dt, UserId : req.UserId}

	dao.Insert(&u)
	return nil
}