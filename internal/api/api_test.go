package api

import (
	"os"
	"testing"

	"github.com/rs/zerolog"

	"github.com/iamnande/ff8-magic-api/internal/datastore"
)

var (
	testAPI API
)

func TestNewAPI(t *testing.T) {

	// test: setup
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	dsc, err := datastore.NewDatastore()
	if err != nil {
		t.Fatalf("failed to initialize datastore: %s", err)
		return
	}

	// test: execute with no failures
	testAPI = NewAPI(logger, dsc)

}

func TestAPI_Log(t *testing.T) {

	// test: execution
	actual := testAPI.Log()

	// test: validation
	if actual.GetLevel() != zerolog.TraceLevel {
		t.Fatalf("expected default zerolog")
	}

}
