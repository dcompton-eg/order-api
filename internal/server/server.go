package server

import (
	"github.com/dcompton-eg/order-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
}

func NewServer(logger *logrus.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s Server) RegisterHandlers(engine *gin.Engine) {
	appGroup := engine.Group("/",
		handlers.NewBaseHandler(),
	)

	oh := handlers.NewOrderHandler(s.logger)
	appGroup.POST("/orders", oh.Create)
}
