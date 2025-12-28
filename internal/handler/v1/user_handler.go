package v1

import (
	"errors"
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
		users.GET(":id", h.GetById)
		users.DELETE(":id", h.Delete)
		users.PUT("", h.Update)
		users.POST("/login", h.Login)
	}
}

// CreateUser 创建用户
// @Summary 创建一个新用户
// @Description 创建一个新用户账号
// @Tags user module
// @Accept json
// @Produce json
// @Param request body request.CreateUserRequest true "创建用户请求参数"
// @Success 200 {object} model.Response "成功响应"
// @Failure 400 {object} model.Response "请求参数错误"
// @Failure 500 {object} model.Response "服务器内部错误"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest

	// 1.从请求中获取数据
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: xerr.CodeInvalidParams,
			Msg:  err.Error(),
		})
		return
	}

	// 2.调用 servive 层
	id, err := h.svc.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: xerr.CodeInternal,
			Msg:  err.Error(),
		})
		return
	}

	// 3.返回成功信息
	c.JSON(http.StatusOK, model.Success(map[string]any{
		"id": id,
	}))
}

// GetUserByIdRequest 用于根据路径参数id查询用户信息
type GetUserByIdRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// GetUserById
// @Schemes
// @Summary 获取用户信息
// @Tags user module
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 json model.User
// @Failure 400 json model.Response "请求参数错误"
// @Failure 404 json model.Response	"用户不存在"
// @Router /users/{id} [get]
func (h *UserHandler) GetById(c *gin.Context) {
	var req GetUserByIdRequest
	// 1.从请求中获取参数，并进行校验
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: xerr.CodeInvalidParams,
			Msg:  err.Error(),
		})
		return
	}

	// 2.调用 service 层
	user, err := h.svc.GetById(req.ID)
	if err != nil {

		// 判断错误类型
		var ce *xerr.CodeError
		if errors.As(err, &ce) {
			c.JSON(http.StatusOK, model.Response{
				Code: xerr.CodeNotFound,
				Msg:  err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, model.Response{
			Code: xerr.CodeInternal,
			Msg:  err.Error(),
		})
		return
	}

	// 3.返回查询结果
	c.JSON(http.StatusOK, model.Success(user))
}

// GetUserByIdRequest 用于根据路径参数id删除用户
type DeleteUserByIdRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

// Delete delete user by id
// @Tags user module
// @Summary delete user by id
// @Param id path string ture "user id"
// @Accept json
// @Produce json
// @Success 200 json model.User
// @Failure 400 json model.Response
// @Failure 404 json model.Response
// @Failure 500 json model.Response
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	// 1.从请求中获取参数，并进行校验
	var req DeleteUserByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: xerr.CodeInvalidParams,
			Msg:  err.Error(),
		})
	}

	// 2.调用 service 层
	if err := h.svc.Delete(req.ID); err != nil {
		var ec *xerr.CodeError
		if errors.As(err, &ec) {
			c.JSON(http.StatusOK, model.Response{
				Code: xerr.CodeNotFound,
				Msg:  err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, model.Response{
			Code: xerr.CodeInternal,
			Msg:  err.Error(),
		})
		return
	}

	// 返回成功信息
	c.JSON(http.StatusOK, model.Success(nil))
}

// Update 更新用户信息
// @Summary  更新用户信息
// @Tags user module
// @Param request body request.UpdateUserRequest ture "更新用户请求参数"
// @Accept json
// @Produce json
// @Success 200 json model.Response
// @Failure 400 json model.Response
// @Failure 404 json model.Response
// @Router /users [put]
func (h *UserHandler) Update(c *gin.Context) {
	var req request.UpdateUserRequest
	// 1.获取请求参数，并进行校验
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: xerr.CodeInvalidParams,
			Msg:  err.Error(),
		})
		return
	}

	// 2.调用 service 层
	if err := h.svc.Update(&req); err != nil {
		var ec *xerr.CodeError
		if errors.As(err, &ec) {
			c.JSON(http.StatusOK, model.Response{
				Code: xerr.CodeNotFound,
				Msg:  err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, model.Response{
			Code: xerr.CodeInternal,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Success(nil))
}

// Login
// @Summary login
// @Tags user modulwe
// @Param login body request.LoginRequest true "用户登入参数"
// @Accept json
// @Produce json
// @Success 200 json model.Response
// @Router /users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	// 1.从请求中获取信息，并校验
	var req *request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: xerr.CodeInvalidParams,
			Msg:  err.Error(),
		})
		return
	}

	// 2.调用 service 层
	if err := h.svc.Login(req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: xerr.CodeNotFound,
			Msg:  err.Error(),
		})
		return
	}

	// 3.返回成功
	c.JSON(http.StatusOK, model.Success(nil))
}
