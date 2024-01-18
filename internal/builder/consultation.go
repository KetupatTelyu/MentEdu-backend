package builder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mentedu-backend/api/consultation"
	"mentedu-backend/internal/app/config"
	"mentedu-backend/internal/repository"
	consultation2 "mentedu-backend/internal/service/consultation"
	userService "mentedu-backend/internal/service/user"
)

func BuildConsultationHandler(cfg config.Config, router *gin.Engine, db *gorm.DB) {
	cr := repository.NewConsultationOrderRepository(db)
	cdr := repository.NewConsultationDetailsRepository(db)
	ur := repository.NewUserRepository(db)
	rp := repository.NewRoleRepository(db)
	pr := repository.NewPermissionRepository(db)
	userRole := repository.NewUserRoleRepository(db)
	rolePermission := repository.NewRolePermissionRepository(db)

	cc := consultation2.NewConsultationCreator(cr, cdr)
	cu := consultation2.NewConsultationUpdater(cr)
	cf := consultation2.NewConsultationFinder(cr, cdr)
	uf := userService.NewUserFinder(ur, rp, pr, userRole, rolePermission)

	consultation.ConsultationCreatorHTTPHandler(cfg, router, cc, uf, cu)
	consultation.ConsultationFinderHTTPHandler(cfg, router, cf, uf)
	consultation.ConsultationUpdaterHTTPHandler(cfg, router, cu, uf)
}
