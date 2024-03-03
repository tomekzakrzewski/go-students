package error

import (
	"fmt"
	"net/http"
)

type Http struct {
	Description string            `json:"description,omitempty"`
	StatusCode  int               `json:"statusCode"`
	Values      map[string]string `json:"values,omitempty"`
}

func (e Http) Error() string {
	return fmt.Sprintf("description: %s", e.Description)
}

func NewHttpError(description string, statusCode int, values map[string]string) Http {
	return Http{
		Description: description,
		StatusCode:  statusCode,
		Values:      values,
	}
}

func InvalidRequestBody() Http {
	return NewHttpError("Invalid request body", http.StatusBadRequest, nil)
}

func InvalidId() Http {
	return NewHttpError("Invalid ID", http.StatusBadRequest, nil)
}

func NotFound() Http {
	return NewHttpError("Not found", http.StatusNotFound, nil)
}

func InvalidValidation(errors map[string]string) Http {
	return NewHttpError("", http.StatusBadRequest, errors)
}
