package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/iamnande/ff8-api/internal/datastore"

	"github.com/iamnande/ff8-api/internal/api"
	"github.com/iamnande/ff8-api/internal/config"
)

// TODO: test everything mo betta
func main() {

	// api: initialize environment configuration
	cfg := config.MustLoad()

	// api: initialize logging instance
	zerolog.SetGlobalLevel(cfg.LogLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// api: initialize root context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// api: initialize datastore implementation client
	ds, err := datastore.NewStaticDS()
	if err != nil {
		logger.Fatal().Msgf("failed to initialize datastore: %s", err)
	}

	// api: initialize the API service
	ff8API := api.NewFF8API(cfg, logger, ds)

	// api: initialize listener
	startListener(ctx, ff8API)

}

// startListener is a function to initialize the HTTP listener.
func startListener(ctx context.Context, api *api.FF8API) {

	// api: handle shutdown gracefully
	go func() {
		<-ctx.Done()
		api.Log().Info().Msg("shutdown signal received")
		_ = api.Shutdown()
	}()

	// api: start RESTful listener
	api.Log().Info().Msg("starting listener")
	if err := api.Serve(); !errors.Is(err, http.ErrServerClosed) {
		api.Log().Fatal().Msgf("failed to start rest listener on %q: %s", api.Address(), err)
	}

}
