package models


type UrlModel struct {

	HashedUrl string
	OriginalUrl string
	CreationDate int64
	ExpiryDate int64
	UserId int

}