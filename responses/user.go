package responses

import (
	"database/sql"
	"github.com/google/uuid"
	"mentedu-backend/internal/model"
	"time"
)

type ProfileResponse struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	PhoneNumber string       `json:"phone_number"`
	Photo       string       `json:"photo"`
	DOB         sql.NullTime `json:"dob"`
	Role        string       `json:"role"`
	CreatedAt   time.Time    `json:"created_at"`
}

func FromUserModel(m *model.User) *ProfileResponse {
	return &ProfileResponse{
		ID:          m.ID,
		Name:        m.Name,
		Email:       m.Email,
		Password:    m.Password,
		PhoneNumber: m.PhoneNumber,
		Photo:       m.Photo,
		DOB:         m.DOB,
		CreatedAt:   m.CreatedAt,
	}
}

type LoginResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}

type RegisterResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}

type RoleResponse struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Permissions []*PermissionResponse `json:"permissions"`
}

func FromRoleModel(m *model.Role) *RoleResponse {
	var permissions []*PermissionResponse

	for _, permission := range m.RolePermissions {
		permissions = append(permissions, &PermissionResponse{
			ID:    permission.PermissionID,
			Name:  permission.Permission.Name,
			Label: permission.Permission.Label,
		})
	}

	return &RoleResponse{
		ID:          m.ID,
		Name:        m.Name,
		Permissions: permissions,
	}
}

type PermissionResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

func FromPermissionModel(m *model.Permission) *PermissionResponse {
	return &PermissionResponse{
		ID:    m.ID,
		Name:  m.Name,
		Label: m.Label,
	}
}
