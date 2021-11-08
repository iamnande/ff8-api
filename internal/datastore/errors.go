package datastore

import (
	"errors"
)

var (
	// ErrItemNotFound indicates the item was not found
	ErrItemNotFound = errors.New("item not found")
)
