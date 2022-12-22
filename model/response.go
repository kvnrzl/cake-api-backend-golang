package model

type ResponseStatusOK struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseError struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}
