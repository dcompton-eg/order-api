package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewBaseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("logger", logrus.New())
	}
}
