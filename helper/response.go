package models

import "strings"

// response structure
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// for empty object structure
type Empty struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

// for error response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splettedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splettedError,
		Data:    data,
	}
	return res
}
