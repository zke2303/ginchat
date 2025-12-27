package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanfeng/ginchat/internal/model"
	"github.com/nanfeng/ginchat/internal/model/request"
	"github.com/nanfeng/ginchat/internal/pkg/xerr"
	"github.com/nanfeng/ginchat/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) Register(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.POST("", h.CreateUser)
	}
}

// @BasePath /api/v1
// PingExample godoc
// @Summary 创建一个用户
// @Schemes
// @Description 创建一个新用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @Success 200 {string} map[string]any
// @Router /user [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest

	// 1.从请求中获取数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: xerr.CodeInvalidParams,
			Msg:  err.Error(),
		})
	}

	// 2.调用 servive 层
	id, err := h.svc.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: xerr.CodeInternal,
			Msg:  err.Error(),
		})
	}

	// 3.返回成功信息
	c.JSON(http.StatusOK, model.Success(map[string]any{
		"id": id,
	}))
}
