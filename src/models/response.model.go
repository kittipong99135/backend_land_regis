package models

type JsonResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type StringResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}
