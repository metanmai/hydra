package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var log = logrus.New()

		log.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})

		log.SetLevel(logrus.InfoLevel)
		log.WithFields(logrus.Fields{
			"path":   ctx.Request.URL.Path,
			"method": ctx.Request.Method,
		}).Info("Request received")

		ctx.Set("logger", log)

		ctx.Next()
	}
}