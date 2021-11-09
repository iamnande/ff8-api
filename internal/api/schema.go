package api

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/aws/aws-lambda-go/events"

	"github.com/iamnande/ff8-api/internal/datastore"
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
	bag := make(map[string]string, 0)

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
		strings.TrimRight(msg, ",")
		return nil, fmt.Errorf(msg)
	}

	// bind: return clean request input
	return req, nil

}

// response will convert native types into an API Gateway consumable response
// entity.
func response(code int, obj interface{}) (*events.APIGatewayV2HTTPResponse, error) {

	// response: serialize response object
	body, _ := json.Marshal(obj)

	// response: return converted response
	return &events.APIGatewayV2HTTPResponse{
		StatusCode: code,
		Body:       string(body),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil

}
