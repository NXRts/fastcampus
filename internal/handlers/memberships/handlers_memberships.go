package memberships

import (
	"context"

	"github.com/NXRts/fsatcampus/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SingUp(ctx context.Context, req memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoutes() {
	route := h.Group("/memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sing-up", h.SingUp)
}
