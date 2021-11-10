package api

import (
	"fmt"
	"testing"
)

func TestNewAPIError(t *testing.T) {

	// test: initialize a new error to validate against
	expected := fmt.Errorf("blew up gloriously")

	// test: initialize a new API error
	actual := NewAPIError(expected)

	// test: validate our error message was stored properly
	if actual.Error() != expected.Error() {
		t.Fatal("expected error message to match")
	}

}
