// pkg/middleware/logger.go
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const RequestIDKey = "X-Request-ID"

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		// Generate or get request ID.
		requestID := ctx.GetHeader(RequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx.Writer.Header().Set(RequestIDKey, requestID)
		ctx.Set(RequestIDKey, requestID)

		// Process request.
		ctx.Next()

		// Log after response.
		latency := time.Since(start)
		status := ctx.Writer.Status()

		log.WithFields(log.Fields{
			"timestamp":  time.Now().Format(time.RFC3339),
			"status":     status,
			"latency":    latency,
			"request_id": requestID,
		}).Infof("%s %s", ctx.Request.Method, ctx.FullPath())
	}
}
