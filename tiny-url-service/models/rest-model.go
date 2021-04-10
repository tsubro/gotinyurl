package models


type Request struct {
	UserId int 					`json:"user_id"`
	Url string 					`json:"url"`
}

type Response struct {
	TinyUrl string 				`json:"tiny_url"`
	Url string					`json:"url"`
	ExpiryTime int64			`json:"expiry_time"`
}

