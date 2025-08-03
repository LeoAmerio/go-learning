package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int                 `json:"status"`
	Data          interface{}         `json:"data"`
	Message       string              `json:"message"`
	contentType   string              `json:"-"`
	responseWrite http.ResponseWriter `json:"-"`
}

func CreateDefaultResponse(rw http.ResponseWriter) *Response {
	return &Response{
		Status:        http.StatusOK,
		Data:          nil,
		Message:       "",
		contentType:   "application/json",
		responseWrite: rw,
	}
}

func (resp *Response) Send() {
	resp.responseWrite.Header().Set("Content-Type", resp.contentType)
	resp.responseWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(resp)
	fmt.Fprintln(resp.responseWrite, string(output))
	// json.NewEncoder(resp.responseWrite).Encode(resp)
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (resp *Response) NotFound(message string) {
	resp.Status = http.StatusNotFound
	resp.Message = message
}

func SendNotFound(rw http.ResponseWriter, message string) {
	response := CreateDefaultResponse(rw)
	response.NotFound(message)
	response.Send()
}

func (resp *Response) UnprocessableEntity(message string) {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = message
}

func SendUnprocessableEntity(rw http.ResponseWriter, message string) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity(message)
	response.Send()
}
