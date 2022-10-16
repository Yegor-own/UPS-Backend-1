package models

type EncodeData struct {
	Message string `json:"message"`
	Rot     int    `json:"rot"`
}

type EncodeResult struct {
	Message string `json:"message"`
}
