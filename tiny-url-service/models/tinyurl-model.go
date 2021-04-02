package models


type TinyUrlRequest struct {
	Url string `json:"url"`
}

type TinyUrlResponse struct {
	TinyUrl string `json:"tiny_url"`
}

type TinyUrlError struct {
	Message string `json:"message"`
}