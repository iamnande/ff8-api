package api

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/iamnande/ff8-magic-api/internal/calculator"
	"github.com/iamnande/ff8-magic-api/internal/datastore"
)

func TestValidate(t *testing.T) {

	// test: setup test cases
	testCases := []struct {
		input          *input
		expectError    bool
		expectErrorBag map[string]string
	}{
		{
			input: &input{
				Name:  "Firaga",
				Type:  datastore.Magic,
				Count: 300,
			},
			expectError: false,
		},
		{
			input: &input{
				Name:  "Nah",
				Type:  datastore.Magic,
				Count: 300,
			},
			expectError: true,
			expectErrorBag: map[string]string{
				"name": "name must be at least 4 characters",
			},
		},
		{
			input: &input{
				Type:  datastore.Magic,
				Count: 300,
			},
			expectError: true,
			expectErrorBag: map[string]string{
				"name": "name is required",
			},
		},
		{
			input: &input{
				Name:  "Firaga",
				Type:  datastore.Magic,
				Count: 0.1,
			},
			expectError: true,
			expectErrorBag: map[string]string{
				"count": "count must be at least 1.0 or higher",
			},
		},
	}

	// test: iterate test cases
	for i, tc := range testCases {

		// test: execute
		ok, bag := tc.input.validate()

		// test: validate errors
		if tc.expectError == ok {

			// test: check error matches our expectation
			if !reflect.DeepEqual(tc.expectErrorBag, bag) {
				t.Fatalf("[%d] expected error bags to match", i)
			}

		}

	}

}

func TestBind(t *testing.T) {

	// test: setup test cases
	testCases := []struct {
		inputBody     string
		expectedInput *input
		expectError   bool
	}{
		{
			inputBody: `{"name":"Aero","type":"Magic","count":100}`,
			expectedInput: &input{
				Name:  "Aero",
				Type:  datastore.Magic,
				Count: 100,
			},
			expectError: false,
		},
		{
			inputBody:   `{""::::23@#$@#^$^$"`,
			expectError: true,
		},
		{
			inputBody:   `{"name":"Nah","type":"Magic","count":100}`,
			expectError: true,
		},
	}

	// test: iterate test case
	for i, tc := range testCases {

		// test: execute
		actual, err := bind(tc.inputBody)

		// test: validate errors
		if !tc.expectError && err != nil {
			t.Fatalf("[%d] did not expect error: %s", i, err)
		}

		// test: validate expectations
		if !reflect.DeepEqual(tc.expectedInput, actual) {
			t.Fatalf("[%d] expected proper deserialization: %+v", i, actual)
		}

	}

}

func TestResponse(t *testing.T) {

	// test: setup expectations
	testCases := []struct {
		inputCode    int
		inputBody    interface{}
		expectedBody string
	}{
		{
			inputCode: http.StatusOK,
			inputBody: &calculator.Response{
				Name:  "Water",
				Card:  "Fastitocalon",
				Count: 6,
			},
			expectedBody: `{"name":"Water","card":"Fastitocalon","count":6}`,
		},
		{
			inputCode:    http.StatusOK,
			inputBody:    make(chan int),
			expectedBody: `"message":"failed to generate response object"`,
		},
	}

	// test: iterate test cases
	for i, tc := range testCases {

		// test: execute functionality
		actual, err := response(tc.inputCode, tc.inputBody)

		// test: ensure we never fail the lambda, always return a response
		if err != nil || actual == nil {
			t.Fatalf("catastrophic lambda failure not expected: %s", err)
		}

		// test: validate expectations
		if !strings.Contains(actual.Body, tc.expectedBody) {
			t.Fatalf("[%d] expected serialization of response: %+v",
				i, actual.Body)
		}

	}

}

// 1st Section Updates
