package error

import (
	"fmt"
	"net/http"
)

type Http struct {
	Description string `json:"description,omitempty"`
	StatusCode  int    `json:"statusCode"`
}

func (e Http) Error() string {
	return fmt.Sprintf("description: %s", e.Description)
}

func NewHttpError(description string, statusCode int) Http {
	return Http{
		Description: description,
		StatusCode:  statusCode,
	}
}

// InvalidRequestBody returns an HTTP response for invalid request bodies.
//	@Summary		Invalid Request Body
//	@Description	Indicates that the request body is invalid.
//	@ID				invalid-request-body
//	@Produce		json
//	@Success		400	{object}	Http	"Invalid request body"
func InvalidRequestBody() Http {
	return NewHttpError("Invalid request body", http.StatusBadRequest)
}

// InvalidId returns an HTTP response for invalid IDs.
//	@Summary		Invalid ID
//	@Description	Indicates that the provided ID is invalid.
//	@ID				invalid-id
//	@Produce		json
//	@Success		400	{object}	Http	"Invalid ID"
func InvalidId() Http {
	return NewHttpError("Invalid ID", http.StatusBadRequest)
}

// NotFound returns an HTTP response for resource not found errors.
//	@Summary		Not Found
//	@Description	Indicates that the requested resource was not found.
//	@ID				not-found
//	@Produce		json
//	@Success		404	{object}	Http	"Not found"
func NotFound() Http {
	return NewHttpError("Not found", http.StatusNotFound)
}
