package builder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mentedu-backend/api/user"
	"mentedu-backend/internal/app/config"
	"mentedu-backend/internal/repository"
	auth2 "mentedu-backend/internal/service/auth"
)

func BuildAuthHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	ar := repository.NewAuthRepository(db)
	ur := repository.NewUserRepository(db)

	as := auth2.NewAuthService(cfg, ur, ar)

	user.UserAuthHTTPHandler(cfg, router, as)
}
