package builder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	user2 "mentedu-backend/api/user"
	"mentedu-backend/internal/app/config"
	"mentedu-backend/internal/repository"
	userService "mentedu-backend/internal/service/user"
	"mentedu-backend/sdk/gcp"
)

func BuildUserHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	rr := repository.NewRoleRepository(db)
	urr := repository.NewUserRoleRepository(db)
	pr := repository.NewPermissionRepository(db)
	rpr := repository.NewRolePermissionRepository(db)

	uc := userService.NewUserCreator(cfg, ur, rr, urr, pr, rpr)
	uf := userService.NewUserFinder(ur, rr, pr, urr, rpr)
	ud := userService.NewUserDeleter(ur, urr, rr, rpr, pr)
	up := userService.NewUserUpdater(ur, rr, pr, urr, rpr)

	cloudStorage, err := gcp.NewGoogleCloudStorage(cfg.GCPConfig.ProjectID, cfg.GCPConfig.Bucket, cfg.GCPConfig.ServiceAccountPath)

	if err != nil {
		panic(err)
	}

	user2.UserCreatorHTTPHandler(cfg, router, uc, uf, cloudStorage)
	user2.UserFinderHTTPHandler(cfg, router, uc, uf, cloudStorage)
	user2.UserDeleterHTTPHandler(cfg, router, ud, uf, cloudStorage)
	user2.UserUpdaterHTTPHandler(cfg, router, up, uf, cloudStorage)
	user2.UserRegisterHTTPHandler(cfg, router, uc, uf, cloudStorage)
	user2.UserProfileHTTPHandler(cfg, router, uf, cloudStorage)
}
