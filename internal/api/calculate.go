package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/iamnande/ff8-api/internal/datastore"
)

// CalculateInput is the calculation input operation.
type CalculateInput struct {
	Name  string               `json:"name"`
	Type  datastore.RecordType `json:"type"`
	Count float64              `json:"count"`
}

// CalculateOutput is the calculation input operation.
type CalculateOutput struct {
	Name     string `json:"name"`
	Card     string `json:"card"`
	Quantity int    `json:"quantity"`
}

// Calculate is the API endpoint for calculating the conversion of cards to
// Magic or Limit Breaks.
// TODO: test the handler
func (api *FF8API) Calculate(c echo.Context) error {

	// calculate: deserialize input
	input := new(CalculateInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request format",
		})
	}

	// calculate: perform calculation
	output, err := api.calc.CardMagicRatio(c.Request().Context(), input.Name, input.Type, input.Count)
	if err != nil {
		// TODO: move the errors to a file (known types/schema)
		// TODO: move the error handling to a central function/method
		if errors.Is(err, datastore.ErrItemNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// calculate: fetch items to be calculated
	return c.JSON(http.StatusOK, &CalculateOutput{
		Name:     output.Name,
		Card:     output.Card,
		Quantity: output.Quantity,
	})

}
