package api

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog"

	"github.com/iamnande/ff8-magic-api/internal/calculator"
	"github.com/iamnande/ff8-magic-api/internal/datastore"
)

// API interface describes the required implementation all API providers must
// satisfy.
type API interface {

	// Calculate defines the required interaction syntax for calculating the
	// number of Triple Triad cards required for N desired Magic.
	Calculate(
		ctx context.Context,
		request events.APIGatewayProxyRequest,
	) (*events.APIGatewayProxyResponse, error)

	// Log defines the required interaction syntax for retrieving an instance
	// of the application logger.
	Log() *zerolog.Logger
}

// api is the FF8 API implementation.
type api struct {
	ds   datastore.Datastore
	calc calculator.Calculator
	log  zerolog.Logger
}

// compile time validation that the current implementation satisfies the
// defined interface.
var _ API = (*api)(nil)

// NewAPI initializes a fresh instance of the FF8 API.
func NewAPI(log zerolog.Logger, ds datastore.Datastore) API {
	return &api{
		ds:   ds,
		log:  log,
		calc: calculator.NewCalculator(ds),
	}
}

func (a *api) Log() *zerolog.Logger {
	logger := a.log.With().Logger()
	return &logger
}
