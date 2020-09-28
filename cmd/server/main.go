package main

import (
	"github.com/ejiro-edwin/finance-app/internal/api"
	"github.com/ejiro-edwin/finance-app/internal/api/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main(){
	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithField("version", config.Version).Debug("starting server.")

	router, err := api.NewRouter()
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}

	const addr = "0.0.0.0:8080"
	server := http.Server{
		Handler: router,
		Addr: addr,
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed{
		logrus.WithError(err).Error("Server failed")
	}
}