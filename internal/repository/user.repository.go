package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"mentedu-backend/internal/model"
	"time"
)

type UserRepositoryUseCase interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.User, int64, error)
	GetById(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetUsersByRoleID(ctx context.Context, roleID int) ([]*model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryUseCase {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.db.WithContext(ctx).Model(&model.User{}).Create(&user).Error; err != nil {
		return errors.Wrap(err, "error creating user")
	}
	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	oldTime := user.UpdatedAt
	newTime := time.Now()

	var txnError error

	txnError = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sourceModel := new(model.User)
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(sourceModel, user.ID).Error; err != nil {
			log.Println("[UserRepository-Update]", err)
			return err
		}

		updates := sourceModel.MapUpdateFrom(user)
		if len(*updates) > 0 {
			(*updates)["updated_at"] = newTime

			if err := tx.Model(&model.User{}).Where("id = ?", user.ID).UpdateColumns(updates).Error; err != nil {
				log.Println("[UserRepository-Update]", err)
				return err
			}
		}

		return nil
	})

	if txnError != nil {
		user.UpdatedAt = oldTime
		return txnError
	}

	user.UpdatedAt = newTime
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&model.User{}).Delete(&model.User{}, id).Error; err != nil {
		return errors.Wrap(err, "error deleting user")
	}
	return nil
}

func (repo *UserRepository) GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.User, int64, error) {
	var users []*model.User

	var total int64

	q := repo.db.WithContext(ctx).Preload("UserRole.Role").Model(&model.User{})

	if query != "" {
		q = q.Where("email ILIKE ?", "%"+query+"%")
	}

	if sort != "" && order != "" {
		q = q.Order(sort + " " + order)
	} else {
		q = q.Order("created_at DESC")
	}

	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}

	if err := q.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *UserRepository) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("UserRole.Role").Model(&model.User{}).First(&user, id).Error; err != nil {
		return nil, errors.Wrap(err, "error getting user by id")
	}
	return &user, nil
}

func (repo *UserRepository) GetUsersByRoleID(ctx context.Context, roleID int) ([]*model.User, error) {
	var users []*model.User

	rawSQL := `
		SELECT users.*
		FROM users
		INNER JOIN user_roles ON users.id = user_roles.user_id
		WHERE user_roles.role_id = ?;
	`

	result := repo.db.WithContext(ctx).Raw(rawSQL, roleID).Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
