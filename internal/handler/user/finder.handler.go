package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"mentedu-backend/common"
	"mentedu-backend/internal/app/middleware"
	userService "mentedu-backend/internal/service/user"
	"mentedu-backend/resource"
	"mentedu-backend/responses"
	"mentedu-backend/utils"
	"strconv"
)

type UserFinderHandler struct {
	userFinder   userService.UserFinderUseCase
	cloudStorage utils.CloudStorage
}

func NewUserFinderHandler(userFinder userService.UserFinderUseCase, cloudStorage utils.CloudStorage) *UserFinderHandler {
	return &UserFinderHandler{
		userFinder:   userFinder,
		cloudStorage: cloudStorage,
	}
}

func (ufh *UserFinderHandler) FindUserById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	user, err := ufh.userFinder.FindUser(c.Request.Context(), id)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	role, err := ufh.userFinder.FindRole(c.Request.Context(), user.UserRole.RoleID)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	response := responses.FromUserModel(user)

	response.Role = role.Name

	c.JSON(200, response)
}

func (ufh *UserFinderHandler) UserProfile(c *gin.Context) {
	id := middleware.UserID

	user, err := ufh.userFinder.FindUser(c.Request.Context(), id)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	role, err := ufh.userFinder.FindRole(c.Request.Context(), user.UserRole.RoleID)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	response := responses.FromUserModel(user)

	response.Role = role.Name

	c.JSON(200, response)
}

func (ufh *UserFinderHandler) FindUsers(c *gin.Context) {
	var request resource.QueryRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	users, err := ufh.userFinder.FindAllUser(c.Request.Context(), request.Query, request.Sort, request.Order, request.Limit, request.Offset)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	var response []*responses.ProfileResponse

	for _, user := range users {
		role, err := ufh.userFinder.FindRole(c.Request.Context(), user.UserRole.RoleID)

		if err != nil {
			c.JSON(400, common.ErrBadRequest)
			return
		}

		r := responses.FromUserModel(user)
		r.Role = role.Name

		response = append(response, r)
	}

	c.JSON(200, response)
}

func (ufh *UserFinderHandler) FindRoleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Println("[UserFinderHandler-FindRoleById] line 105", err)
		c.JSON(400, common.ErrBadRequest)
		return
	}

	role, err := ufh.userFinder.FindRole(c.Request.Context(), id)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, role)
}

func (ufh *UserFinderHandler) FindRoles(c *gin.Context) {
	var request resource.QueryRequest

	if err := c.ShouldBind(&request); err != nil {
		log.Println("[UserFinderHandler-FindRoles] line 97", err)
		c.JSON(400, common.ErrBadRequest)
		return
	}

	roles, err := ufh.userFinder.FindAllRole(c.Request.Context(), request.Query, request.Sort, request.Order, request.Limit, request.Offset)

	if err != nil {
		log.Println("[UserFinderHandler-FindRoles] line 105", err)
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, roles)
}

func (ufh *UserFinderHandler) FindPermissionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	permission, err := ufh.userFinder.FindPermission(c.Request.Context(), id)

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, permission)
}

func (ufh *UserFinderHandler) FindPermissions(c *gin.Context) {
	permissions, err := ufh.userFinder.FindAllPermission(c.Request.Context())

	if err != nil {
		c.JSON(400, common.ErrBadRequest)
		return
	}

	c.JSON(200, permissions)
}
