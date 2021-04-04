package models


type TinyUrlRequest struct {
	UserId int 	`json:"user_id"`
	Url string 		`json:"url"`
}

type TinyUrlResponse struct {
	TinyUrl string `json:"tiny_url"`
}

type TinyUrlError struct {
	Message string `json:"message"`
}