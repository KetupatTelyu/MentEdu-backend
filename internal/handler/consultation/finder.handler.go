package consultation

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mentedu-backend/internal/service/consultation"
	userService "mentedu-backend/internal/service/user"
	"mentedu-backend/resource"
	"mentedu-backend/responses"
	"net/http"
)

type ConsultationFinderHandler struct {
	consultationFinderUsecase consultation.ConsultationFinderUsecase
	userFinderUsecase         userService.UserFinderUseCase
}

func NewConsultationFinderHandler(consultationFinderUsecase consultation.ConsultationFinderUsecase, userFinderUsecase userService.UserFinderUseCase) *ConsultationFinderHandler {
	return &ConsultationFinderHandler{
		consultationFinderUsecase: consultationFinderUsecase,
		userFinderUsecase:         userFinderUsecase,
	}
}

func (cc *ConsultationFinderHandler) FindConsultationByUserID(c *gin.Context) {
	var consultationRequest resource.QueryRequest

	if err := c.ShouldBind(&consultationRequest); err != nil {
		c.JSON(400, err)
		return
	}

	consultations, total, err := cc.consultationFinderUsecase.FindAllConsultation(c.Request.Context(), consultationRequest.Query, consultationRequest.Sort, consultationRequest.Order, consultationRequest.Limit, consultationRequest.Offset)
	if err != nil {
		c.JSON(500, err)
		return
	}

	var consultationResponses []responses.ConsultationOrderResponse

	for _, consultation := range consultations {
		if err != nil {
			c.JSON(500, err)
			return
		}

		consultationResponses = append(consultationResponses, responses.ConsultationOrderResponse{
			ID:           consultation.ID.String(),
			UserID:       consultation.UserID.String(),
			Purpose:      consultation.Purpose,
			DateTime:     consultation.DateTime.String(),
			Status:       consultation.Status,
			ConsultantID: consultation.ConsultantID.String(),
		})
	}

	response := responses.NewPaginatedResponse(consultationResponses, http.StatusOK, "Success retrieving consultation data", total)

	c.JSON(200, response)
}

// consultatoin detail

func (cc *ConsultationFinderHandler) FindConsultationDetailByID(c *gin.Context) {
	var consultationRequest resource.FindByIDRequest

	if err := c.ShouldBind(&consultationRequest); err != nil {
		c.JSON(400, err)
		return
	}

	id, err := uuid.Parse(consultationRequest.ID)

	if err != nil {
		c.JSON(400, err)
		return
	}

	consultationDetail, err := cc.consultationFinderUsecase.FindConsultationDetail(c.Request.Context(), id)

	if err != nil {
		c.JSON(500, err)
		return
	}

	consultationOrder, err := cc.consultationFinderUsecase.FindConsultation(c.Request.Context(), consultationDetail.ConsultationID)

	if err != nil {
		c.JSON(500, err)
		return
	}

	consultee, err := cc.userFinderUsecase.FindUser(c.Request.Context(), consultationOrder.UserID)

	if err != nil {
		c.JSON(500, err)
		return
	}

	consultant, err := cc.userFinderUsecase.FindUser(c.Request.Context(), consultationOrder.ConsultantID)

	if err != nil {
		c.JSON(500, err)
		return
	}

	consultationDetailResponse := responses.ConsultationDetailResponse{
		ID:             consultationDetail.ID.String(),
		ConsultationID: consultationDetail.ConsultationID.String(),
		MeetingURL:     consultationDetail.MeetingURL,
		ConsultationOrder: responses.ConsultationOrderResponse{
			ID:           consultationOrder.ID.String(),
			UserID:       consultationOrder.UserID.String(),
			Purpose:      consultationOrder.Purpose,
			DateTime:     consultationOrder.DateTime.String(),
			Status:       consultationOrder.Status,
			ConsultantID: consultationOrder.ConsultantID.String(),
		},
		Consultant: responses.ProfileResponse{
			ID:          consultant.ID,
			Name:        consultant.Name,
			Email:       consultant.Email,
			PhoneNumber: consultant.PhoneNumber,
			Photo:       consultant.Photo,
			DOB: sql.NullTime{
				Time: consultant.DOB.Time,
			},
			Role:      "consultant",
			CreatedAt: consultant.CreatedAt,
		},
		Consultee: responses.ProfileResponse{
			ID:          consultee.ID,
			Name:        consultee.Name,
			Email:       consultee.Email,
			PhoneNumber: consultee.PhoneNumber,
			Photo:       consultee.Photo,
			DOB: sql.NullTime{
				Time: consultee.DOB.Time,
			},
			Role:      "user",
			CreatedAt: consultee.CreatedAt,
		},
	}

	response := responses.NewResponse(consultationDetailResponse, http.StatusOK, "Success retrieving consultation detail data")

	c.JSON(200, response)
}
