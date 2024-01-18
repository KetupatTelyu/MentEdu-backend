package consultation

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mentedu-backend/internal/app/middleware"
	"mentedu-backend/internal/model"
	"mentedu-backend/internal/service/consultation"
	"mentedu-backend/internal/service/user"
	"mentedu-backend/resource"
	"time"
)

type ConsultationCreatorHandler struct {
	consultationCreatorUsecase consultation.ConsultationCreatorUsecase
	userFinderUseCase          user.UserFinderUseCase
	consultationUpdaterUsecase consultation.ConsultationUpdaterUsecase
}

func NewConsultationCreatorHandler(consultationCreatorUsecase consultation.ConsultationCreatorUsecase, userFinderUseCase user.UserFinderUseCase, consultationUpdater consultation.ConsultationUpdaterUsecase) *ConsultationCreatorHandler {
	return &ConsultationCreatorHandler{
		consultationCreatorUsecase: consultationCreatorUsecase,
		userFinderUseCase:          userFinderUseCase,
		consultationUpdaterUsecase: consultationUpdater,
	}
}

func (cc *ConsultationCreatorHandler) CreateConsultation(c *gin.Context) {
	var consultationRequest resource.ConsultationOrderRequest

	if err := c.ShouldBind(&consultationRequest); err != nil {
		c.JSON(400, err)
		return
	}

	userID, err := uuid.Parse(consultationRequest.UserID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	consultantID, err := uuid.Parse(consultationRequest.ConsultantID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	parseTime, err := time.Parse("2006-01-02 15:04:05", consultationRequest.DateTime)

	if err != nil {
		c.JSON(400, err)
		return
	}

	consultationOrder := model.NewConsultationOrder(uuid.New(), userID, consultationRequest.Purpose, parseTime, model.ConsultationOrderStatusPending, consultantID, userID.String())

	if err := cc.consultationCreatorUsecase.CreateConsultation(c.Request.Context(), consultationOrder); err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, consultationOrder)
}

func (cc *ConsultationCreatorHandler) CreateConsultationDetail(c *gin.Context) {
	var consultationDetailRequest resource.ConsultationDetailRequest

	if err := c.ShouldBind(&consultationDetailRequest); err != nil {
		c.JSON(400, err)
		return
	}

	consultationID, err := uuid.Parse(consultationDetailRequest.ConsultationOrderID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	consultationDetail := model.NewConsultationDetail(uuid.New(), consultationID, consultationDetailRequest.MeetingURL, middleware.UserID.String())

	if err := cc.consultationCreatorUsecase.CreateConsultationDetail(c.Request.Context(), consultationDetail); err != nil {
		c.JSON(500, err)
		return
	}

	err = cc.consultationUpdaterUsecase.UpdateStatus(c, consultationID, model.ConsultationOrderStatusAccepted)

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, consultationDetail)
}
