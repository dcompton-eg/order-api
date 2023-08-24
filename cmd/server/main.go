package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/dcompton-eg/order-api/internal/server"
)

func main() {
	var (
		port = "8080"
	)

	appServer := server.NewServer(logrus.New())
	r := gin.New()
	r.Use(gin.Recovery())
	appServer.RegisterHandlers(r)

	logrus.Info("starting order-api server")

	e := r.Run(":" + port)
	if e != nil {
		logrus.WithError(e).WithField("port", port).Panic("error starting server")
	}
}
