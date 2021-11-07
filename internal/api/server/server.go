package server

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/iamnande/ff8-api/internal/config"
)

// Server is the HTTP server instance.
type Server struct {
	*echo.Echo
}

// NewServer initializes a new instance of a Server.
func NewServer(cfg *config.Config) *Server {

	// server: initialize the underlying routing framework, echo in this case.
	router := echo.New()

	// server: pre-configure common router items
	// TODO: configure the default error handler
	router.HidePort = true
	router.HideBanner = true

	// server: configure the HTTP server instance itself
	router.Server = &http.Server{
		Addr:              cfg.Port,
		ReadTimeout:       time.Second * 5,
		IdleTimeout:       time.Second * 30,
		WriteTimeout:      time.Second * 15,
		MaxHeaderBytes:    1048576, // 1MB
		ReadHeaderTimeout: time.Second * 2,
	}

	// server: enable common router middleware items
	// TODO: request-id, logging/tracing, CORS?, binding, validation
	router.Pre(middleware.MethodOverride())
	router.Use(
		middleware.Gzip(),
		middleware.Recover(),
		middleware.Decompress(),
	)

	// server: add heartbeat endpoint
	router.GET("/livez", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// server: return fully constructed server
	return &Server{router}

}
