package main

import (
	"os"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/iamnande/ff8-api/internal/api"

	"github.com/iamnande/ff8-api/internal/datastore"
)

var (
	ff8 api.API
)

func main() {

	// main: initialize logging instance
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// main: initialize datastore client
	dsc, err := datastore.NewDatastore()
	if err != nil {
		logger.Fatal().Msgf("failed to initialize datastore: %s", err)
		return
	}

	// main: initialize "API" instance
	ff8 = api.NewAPI(logger, dsc)

	// main: kick off lambda processing
	runtime.Start(ff8.Calculate)

}
