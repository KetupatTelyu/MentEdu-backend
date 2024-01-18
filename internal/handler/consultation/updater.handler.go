package consultation

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mentedu-backend/internal/service/consultation"
	"mentedu-backend/resource"
	"mentedu-backend/responses"
	"net/http"
)

type ConsultationUpdaterHandler struct {
	consultationUpdaterUsecase consultation.ConsultationUpdaterUsecase
}

func NewConsultationUpdaterHandler(consultationUpdaterUsecase consultation.ConsultationUpdaterUsecase) *ConsultationUpdaterHandler {
	return &ConsultationUpdaterHandler{
		consultationUpdaterUsecase: consultationUpdaterUsecase,
	}
}

func (cc *ConsultationUpdaterHandler) UpdateConsultationStatus(c *gin.Context) {
	var consultationRequest resource.ConsultationStatusRequest

	if err := c.ShouldBindJSON(&consultationRequest); err != nil {
		c.JSON(400, err)
		return
	}

	consultationOrderID, err := uuid.Parse(consultationRequest.ConsultationOrderID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	if err := cc.consultationUpdaterUsecase.UpdateStatus(c.Request.Context(), consultationOrderID, consultationRequest.Status); err != nil {
		c.JSON(500, err)
		return
	}

	response := responses.NewResponse(nil, http.StatusOK, "Success updating consultation status")

	c.JSON(200, response)
}
