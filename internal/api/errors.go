package api

import (
	"github.com/google/uuid"
)

// BaseAPIError is the base error constructor model for API error responses.
type BaseAPIError struct {

	// ID is the ID of the error record, intentionally kept separate from the
	// RequestID.
	ID string `json:"error_id"`

	// Message is the client consumable error description.
	Message string `json:"message"`
}

// NewAPIError initializes a new API Error object with a unique ID.
func NewAPIError(err error) *BaseAPIError {
	return &BaseAPIError{
		Message: err.Error(),
		ID:      uuid.NewString(),
	}
}

// Error satisfies the generic error interface. Since we're constructing HTTP
// response errors, this is generally useless outside satisfying the interface.
func (err *BaseAPIError) Error() string {
	return err.Message
}
