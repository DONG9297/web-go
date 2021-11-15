package model

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Success bool              `json:"success"`
	Code    int               `json:"status"`
	Message string            `json:"msg"`
	Data    map[string]string `json:"data"`
}

func Response(w http.ResponseWriter, success bool, code int, message string, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	rst := &Result{
		Success: success,
		Code:    code,
		Message: message,
		Data:    data,
	}
	response, _ := json.Marshal(rst)
	w.Write(response)
}
