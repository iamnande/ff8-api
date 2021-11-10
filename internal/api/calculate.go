package api

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Calculate is the API handler for calculating the conversion of cards to
// Magic or Limit Breaks.
func (a *api) Calculate(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {

	// calculate: fetch request object
	req, err := bind(request.Body)
	if err != nil {
		return response(http.StatusBadRequest, NewAPIError(err))
	}

	// calculate: perform calculation
	res, err := a.calc.CardMagicRatio(ctx, req.Name, req.Type, req.Count)
	if err != nil {
		// TODO: check err type here to provide better status code context
		return response(http.StatusInternalServerError, NewAPIError(err))
	}

	// calculate: return calculation
	return response(http.StatusOK, res)

}
