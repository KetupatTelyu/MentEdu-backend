package consultation

import (
	"github.com/gin-gonic/gin"
	"mentedu-backend/internal/app/config"
	"mentedu-backend/internal/app/middleware"
	consultation2 "mentedu-backend/internal/handler/consultation"
	"mentedu-backend/internal/service/consultation"
	userService "mentedu-backend/internal/service/user"
)

func ConsultationCreatorHTTPHandler(cfg config.Config, router *gin.Engine, consultationCreatorUsecase consultation.ConsultationCreatorUsecase, userFinder userService.UserFinderUseCase, consultationUpdater consultation.ConsultationUpdaterUsecase) {
	hndlr := consultation2.NewConsultationCreatorHandler(consultationCreatorUsecase, userFinder, consultationUpdater)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Consultant(cfg))
	{
		api.POST("/consultation", hndlr.CreateConsultation)
		api.POST("/consultation/detail", hndlr.CreateConsultationDetail)
	}
}

func ConsultationFinderHTTPHandler(cfg config.Config, router *gin.Engine, consultationFinderUsecase consultation.ConsultationFinderUsecase, userFinder userService.UserFinderUseCase) {
	hndlr := consultation2.NewConsultationFinderHandler(consultationFinderUsecase, userFinder)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	{
		api.GET("/consultation", hndlr.FindConsultationByUserID)
		api.GET("/consultation/detail/:id", hndlr.FindConsultationDetailByID)
	}
}

func ConsultationUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, consultationUpdaterUsecase consultation.ConsultationUpdaterUsecase, userFinder userService.UserFinderUseCase) {
	hndlr := consultation2.NewConsultationUpdaterHandler(consultationUpdaterUsecase)

	api := router.Group("/api")

	api.Use(middleware.Auth(cfg))
	api.Use(middleware.Consultant(cfg))
	{
		api.PUT("/consultation/status/:id", hndlr.UpdateConsultationStatus)
	}
}
