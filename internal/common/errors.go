package common

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func ToErrorResponse(e int) []byte {

	var errResp ErrorResponse
	errMessage := http.StatusText(e)

	if errMessage == "" {
		errResp.Message = http.StatusText(http.StatusInternalServerError)
	} else {
		errResp.Message = errMessage
	}

	responseJSON, _ := json.Marshal(errResp)

	return responseJSON
}
