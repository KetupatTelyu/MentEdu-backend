package builder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	article2 "mentedu-backend/api/article"
	"mentedu-backend/internal/app/config"
	"mentedu-backend/internal/repository"
	"mentedu-backend/internal/service/article"
	userService "mentedu-backend/internal/service/user"
	"mentedu-backend/sdk/local"
)

func BuildArticleHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	ar := repository.NewArticleRepository(db)
	ur := repository.NewUserRepository(db)
	rp := repository.NewRoleRepository(db)
	pr := repository.NewPermissionRepository(db)
	userRole := repository.NewUserRoleRepository(db)
	rolePermission := repository.NewRolePermissionRepository(db)

	cloudStorage := local.NewLocalStorage(cfg.LocalStorage.BasePath)

	as := article.NewArticleCreator(cfg, ar, cloudStorage)
	af := article.NewArticleFinder(ar)
	au := article.NewArticleUpdater(ar)
	ad := article.NewArticleDeleter(ar)

	uf := userService.NewUserFinder(ur, rp, pr, userRole, rolePermission)

	article2.ArticleCreatorHTTPHandler(cfg, router, as, uf, cloudStorage)
	article2.ArticleFinderHTTPHandler(cfg, router, af, uf, cloudStorage)
	article2.ArticleUpdaterHTTPHandler(cfg, router, au, uf, cloudStorage)
	article2.ArticleDeleterHTTPHandler(cfg, router, ad, uf, cloudStorage)
}
