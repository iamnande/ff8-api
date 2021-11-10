package api

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"

	"github.com/iamnande/ff8-magic-api/internal/datastore"
)

func TestAPI_Calculate(t *testing.T) {

	// test: setup
	ctx := context.Background()
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	dsc, err := datastore.NewDatastore()
	if err != nil {
		t.Fatalf("failed to initialize datastore: %s", err)
		return
	}
	testAPI := NewAPI(logger, dsc)

	// test: setup test cases
	testCases := []struct {
		input              events.APIGatewayProxyRequest
		expectedStatusCode int
	}{
		{
			input: events.APIGatewayProxyRequest{
				Body: `{"name":"Life","type":"Magic","count":100}`,
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			input: events.APIGatewayProxyRequest{
				Body: `{""::::23@#$@#^$^$"`,
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			input: events.APIGatewayProxyRequest{
				Body: `{"name":"Potatoes","type":"Mashed","count":1000000.00}`,
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	// test: iterate test cases
	for i, tc := range testCases {

		// test: execute
		res, _ := testAPI.Calculate(ctx, tc.input)

		// test: validate
		if res.StatusCode != tc.expectedStatusCode {
			t.Fatalf("[%d] expected successful response: %d", i, res.StatusCode)
		}

	}

}
