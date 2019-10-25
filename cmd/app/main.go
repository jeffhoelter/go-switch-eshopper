package main

import (
	"fmt"
	"net/http"

	"github.com/jeffhoelter/go-switch-eshopper/app/app"
	"github.com/jeffhoelter/go-switch-eshopper/app/router"
	"github.com/jeffhoelter/go-switch-eshopper/config"
	lr "github.com/jeffhoelter/go-switch-eshopper/util/logger"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	application := app.New(logger)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Printf("Starting server %s\n", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
