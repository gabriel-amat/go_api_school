package models

type ResponseModel struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
