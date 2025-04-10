package router

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"rockstar-corp/internal/config"
	"rockstar-corp/internal/infrastructure/db"
)

func New(cfg *config.Config, dbConn *db.DB) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// TODO: add more routes here (user, product, order...)

	return r
}
