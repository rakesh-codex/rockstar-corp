package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rockstar-corp/internal/usecase/user"
)

type UserHandler struct {
	uc user.UseCase
}

func NewUserHandler(r *gin.Engine, uc user.UseCase) {
	h := &UserHandler{uc: uc}

	group := r.Group("/users")
	{
		group.POST("/", h.Register)
		group.GET("/:id", h.GetProfile)
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.uc.Register(c.Request.Context(), &u); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	u, err := h.uc.GetProfile(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, u)
}
