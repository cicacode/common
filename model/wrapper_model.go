package model

type Lang struct {
	Id string `json:"id"`
	En string `json:"en"`
}

type Response struct {
	Code    int  `json:"code"`
	Message Lang `json:"message"`
}

type Token struct {
	Token string `json:"token"`
}
