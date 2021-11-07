package api

import (
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/iamnande/ff8-api/internal/datastore"
)

type CalculateRequest struct {
	Name  string               `json:"name"`
	Type  datastore.RecordType `json:"type"`
	Count float64              `json:"count"`
}

// CalculateResponse is something nice.
// TODO: struct comments
type CalculateResponse struct {
	Name     string `json:"name"`
	Card     string `json:"card"`
	Quantity int    `json:"quantity"`
}

// HandleCalculate handles the calculation of number of cards required for the
// desired magic.
// TODO: move all business logic into private method below
func (api *FF8API) HandleCalculate(c echo.Context) error {

	// calculate: deserialize input
	// TODO: log something, like anything
	input := new(CalculateRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request format",
		})
	}

	// calculate: fetch desired item from datastore
	record, err := api.ds.DescribeRecord(input.Name, input.Type)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": fmt.Sprintf("%s of type %s not found",
				input.Name, input.Type),
		})
	}

	// calculate: perform ratio calculation
	quantity := input.Count / record.CardMagicRatio
	output := &CalculateResponse{
		Name:     record.Name,
		Card:     record.CardEquivalent,
		Quantity: int(math.Round(quantity*10) / 10),
	}

	// calculate: fetch items to be calculated
	return c.JSONPretty(http.StatusOK, output, "  ")

}
