package common

import (
	"encoding/json"
	"errors"
)

type ErrorCode int

const (
	BadRequest          ErrorCode = 400
	Unauthorized        ErrorCode = 403
	InternalServerError ErrorCode = 500
)

func (ec ErrorCode) IsValid() bool {
	switch ec {
	case BadRequest, Unauthorized, InternalServerError:
		return true
	}
	return false
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func ToErrorResponse(e ErrorCode) ([]byte, error) {

	errorResp := ErrorResponse{
		Message: "Internal Server Error",
	}

	if e == BadRequest {
		errorResp = ErrorResponse{
			Message: "Bad Request",
		}
	}
	if e == Unauthorized {
		errorResp = ErrorResponse{
			Message: "Unauthorized",
		}
	}

	isValid := e.IsValid()
	responseJSON, _ := json.Marshal(errorResp)

	if !isValid {
		return responseJSON, errors.New("Error code invalid")
	}

	return responseJSON, nil
}
