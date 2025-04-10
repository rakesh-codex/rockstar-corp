package router

import (
	"github.com/gin-gonic/gin"
	"rockstar-corp/internal/config"
	"rockstar-corp/internal/domain/user"
	userRepo "rockstar-corp/internal/infrastructure/repository"
	userHttp "rockstar-corp/internal/interfaces/http"
	userUC "rockstar-corp/internal/usecase/user"

	"rockstar-corp/internal/infrastructure/db"
)

func New(cfg *config.Config, dbConn *db.DB) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// User module wiring
	userRepository := userRepo.NewUserRepository(dbConn.Conn)
	userUseCase := userUC.New(userRepository)
	userHttp.NewUserHandler(r, userUseCase)

	return r
}
