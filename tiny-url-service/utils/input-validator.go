package utils

import(
	"tiny-url/models"
	"net/url"
	"errors"
)

func ValidateInput(req *models.TinyUrlRequest) error {

	if req.Url == "" {
		return errors.New("Input Url Is Empty")
	}
	_, err := url.ParseRequestURI(req.Url)
	if err != nil {
		return errors.New("Input Url Not In Proper Format")
	}
}