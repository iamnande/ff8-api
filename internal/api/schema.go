package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/aws/aws-lambda-go/events"

	"github.com/iamnande/ff8-magic-api/internal/datastore"
)

// input is the calculation input operation.
type input struct {
	Name  string               `json:"name"`
	Type  datastore.RecordType `json:"type"`
	Count float64              `json:"count"`
}

func (i *input) validate() (bool, map[string]string) {

	// validate: error bag
	ok := true
	bag := make(map[string]string)

	//  validate: name
	if i.Name != "" {
		if utf8.RuneCountInString(i.Name) < 4 {
			ok = false
			bag["name"] = "name must be at least 4 characters"
		}
	} else {
		ok = false
		bag["name"] = "name is required"
	}

	//  validate: count
	if i.Count < 1.0 {
		ok = false
		bag["count"] = "count must be at least 1.0 or higher"
	}

	//  validate: return items
	return ok, bag

}

// bind will translate the HTTP request schema into an input object.
func bind(body string) (*input, error) {

	// bind: deserialize request into input
	req := new(input)
	if err := json.Unmarshal([]byte(body), req); err != nil {
		return nil, err
	}

	// bind: validate we have required fields
	if ok, bag := req.validate(); !ok {
		var msg string
		for key, err := range bag {
			msg += fmt.Sprintf("%s: %s,", key, err)
		}
		msg = strings.TrimRight(msg, ",")
		return nil, fmt.Errorf(msg)
	}

	// bind: return clean request input
	return req, nil

}

// response will convert native types into an API Gateway consumable response
// entity.
func response(code int, obj interface{}) (*events.APIGatewayProxyResponse, error) {

	// response: serialize response object
	body, err := json.Marshal(obj)

	// response: generate internal error if marshaling error found
	if err != nil {
		code = http.StatusInternalServerError
		body, _ = json.Marshal(NewAPIError(
			fmt.Errorf("failed to generate response object"),
		))
	}

	// response: return converted response
	return &events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil

}
