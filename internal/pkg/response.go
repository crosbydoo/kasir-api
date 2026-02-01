package pkg

import (
	"encoding/json"
	"net/http"
)

type ResponsePayload struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	response := ResponsePayload{
		Code:    code,
		Status:  true,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ResponseError(w http.ResponseWriter, code int, message string, data interface{}) {
	response := ResponsePayload{
		Code:    code,
		Status:  false,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
