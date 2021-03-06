package main

import (
	"github.com/ejiro-edwin/finance-app/internal/api"
	"github.com/ejiro-edwin/finance-app/internal/api/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Create Server object and start Listener
func main(){
	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithField("version", config.Version).Debug("starting server.")

	//Creating new router
	router, err := api.NewRouter()
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}

	const addr = "0.0.0.0:8080"
	server := http.Server{
		Handler: router,
		Addr: addr,
	}

	// Starting Server
	if err := server.ListenAndServe(); err != http.ErrServerClosed{
		logrus.WithError(err).Error("Server failed")
	}
}