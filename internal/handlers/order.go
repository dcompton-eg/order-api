package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Order struct {
	ID       string `json:"id"`
	Location string `json:"location"`
	Item     string `json:"item"`
}

func NewOrderHandler(logger *logrus.Logger) *OrderHandler {
	return &OrderHandler{
		logger: logger,
	}
}

type OrderHandler struct {
	logger *logrus.Logger
}

func (oh *OrderHandler) Create(c *gin.Context) {
	var o Order
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o.ID = "TODO"

	// TODO store order.

	c.JSON(http.StatusCreated, gin.H{"order": o})
}
