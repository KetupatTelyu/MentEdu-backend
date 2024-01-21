package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"mentedu-backend/utils"
	"os"
	"time"
)

type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Name        string     `gorm:"type:varchar(128);not null"`
	Email       string     `gorm:"type:varchar(128);not null"`
	Password    string     `gorm:"type:varchar(255);not null"`
	PhoneNumber string     `gorm:"type:varchar(50);not null"`
	DOB         *time.Time `gorm:"type:date"`
	Photo       string     `gorm:"type:text"`
	Status      string     `gorm:"type:varchar(50)"`
	ForgotToken string     `gorm:"type:text"`
	CreatedBy   string     `gorm:"type:varchar(128);not null"`
	UpdatedBy   string     `gorm:"type:varchar(128);not null"`
	DeletedBy   string     `gorm:"type:varchar(128)"`
	CreatedAt   time.Time  `gorm:"type:timestamptz;not null"`
	UpdatedAt   time.Time  `gorm:"type:timestamptz;not null"`
	DeletedAt   *time.Time `gorm:"type:timestamptz"`
}

type Role struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"type:varchar(128);not null"`
	CreatedBy string     `gorm:"type:varchar(128)"`
	UpdatedBy string     `gorm:"type:varchar(128)"`
	DeletedBy string     `gorm:"type:varchar(128)"`
	CreatedAt time.Time  `gorm:"type:timestamptz"`
	UpdatedAt time.Time  `gorm:"type:timestamptz"`
	DeletedAt *time.Time `gorm:"type:timestamptz"`
}

type UserRole struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null"`
	RoleID    uint       `gorm:"not null"`
	CreatedBy string     `gorm:"type:varchar(128)"`
	UpdatedBy string     `gorm:"type:varchar(128)"`
	DeletedBy string     `gorm:"type:varchar(128)"`
	CreatedAt time.Time  `gorm:"type:timestamptz"`
	UpdatedAt time.Time  `gorm:"type:timestamptz"`
	DeletedAt *time.Time `gorm:"type:timestamptz"`
}

type Permission struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"type:varchar(128);not null"`
	Label     string     `gorm:"type:varchar(128)"`
	CreatedBy string     `gorm:"type:varchar(128)"`
	UpdatedBy string     `gorm:"type:varchar(128)"`
	DeletedBy string     `gorm:"type:varchar(128)"`
	CreatedAt time.Time  `gorm:"type:timestamptz"`
	UpdatedAt time.Time  `gorm:"type:timestamptz"`
	DeletedAt *time.Time `gorm:"type:timestamptz"`
}

type RolePermission struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey"`
	RoleID       uint       `gorm:"not null"`
	PermissionID uint       `gorm:"not null"`
	CreatedBy    string     `gorm:"type:varchar(128)"`
	UpdatedBy    string     `gorm:"type:varchar(128)"`
	DeletedBy    string     `gorm:"type:varchar(128)"`
	CreatedAt    time.Time  `gorm:"type:timestamptz"`
	UpdatedAt    time.Time  `gorm:"type:timestamptz"`
	DeletedAt    *time.Time `gorm:"type:timestamptz"`
}

type Category struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"type:varchar(128);not null"`
	CreatedBy string     `gorm:"type:varchar(128)"`
	UpdatedBy string     `gorm:"type:varchar(128)"`
	DeletedBy string     `gorm:"type:varchar(128)"`
	CreatedAt time.Time  `gorm:"type:timestamptz"`
	UpdatedAt time.Time  `gorm:"type:timestamptz"`
	DeletedAt *time.Time `gorm:"type:timestamptz"`
}

type ArticleCategory struct {
	ArticleID  int `json:"article_id"`
	CategoryID int `json:"category_id"`
}

func main() {
	godotenv.Load(".env")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate the tables
	err = db.AutoMigrate(&User{}, &Role{}, &UserRole{}, &Permission{}, &RolePermission{})
	if err != nil {
		log.Fatal(err)
	}

	id, err := uuid.Parse("50574f94-18b2-4f69-9933-5d1dc99bff3a")

	if err != nil {
		log.Fatal(err)
	}

	// Define the admin user and role
	adminUser := User{
		ID:          id,
		Name:        "Admin",
		Email:       "admin@example.com",
		Password:    "adminpassword",
		PhoneNumber: "1234567890",
		CreatedBy:   "seeder",
		UpdatedBy:   "seeder",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	adminUser.Password, _ = utils.HashPassword(adminUser.Password)

	adminRole := Role{
		Name:      "admin",
		CreatedBy: "seeder",
		UpdatedBy: "seeder",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Create admin user and role
	err = db.Create(&adminUser).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Create(&adminRole).Error
	if err != nil {
		log.Fatal(err)
	}

	// Assign admin role to admin user
	userRole := UserRole{
		ID:        uuid.New(),
		UserID:    adminUser.ID,
		RoleID:    adminRole.ID,
		CreatedBy: "seeder",
		UpdatedBy: "seeder",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = db.Create(&userRole).Error
	if err != nil {
		log.Fatal(err)
	}

	normalRole := Role{
		Name:      "user",
		CreatedBy: "seeder",
		UpdatedBy: "seeder",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = db.Create(&normalRole).Error
	if err != nil {
		log.Fatal(err)
	}

	consultant := User{
		ID:          uuid.New(),
		Name:        "Consultant",
		Email:       "consultant@mail.com",
		Password:    "consultantpassword",
		PhoneNumber: "1234567890",
		CreatedBy:   "seeder",
		UpdatedBy:   "seeder",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	consultant.Password, _ = utils.HashPassword(consultant.Password)

	err = db.Create(&consultant).Error

	if err != nil {
		log.Fatal(err)
	}

	consultantRole := Role{
		Name:      "consultant",
		CreatedBy: "seeder",
		UpdatedBy: "seeder",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = db.Create(&consultantRole).Error

	if err != nil {
		log.Fatal(err)
	}

	consultantUserRole := UserRole{
		ID:        uuid.New(),
		UserID:    consultant.ID,
		RoleID:    consultantRole.ID,
		CreatedBy: "seeder",
		UpdatedBy: "seeder",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = db.Create(&consultantUserRole).Error

	if err != nil {
		log.Fatal(err)
	}

	// Define permissions
	readPermission := Permission{
		Name:      "read",
		Label:     "Read",
		CreatedBy: "seeder",
		UpdatedBy: "seeder",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	writePermission := Permission{
		Name:      "write",
		Label:     "Write",
		CreatedBy: "seeder",
		UpdatedBy: "seeder",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	deletePermission := Permission{
		Name:      "delete",
		Label:     "Delete",
		CreatedBy: "seeder",
		UpdatedBy: "seeder",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	updatePermission := Permission{
		Name:      "update",
		Label:     "Update",
		CreatedBy: "seeder",
		UpdatedBy: "seeder",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Create permissions
	err = db.Create(&readPermission).Error

	if err != nil {
		log.Fatal(err)
	}

	err = db.Create(&writePermission).Error

	if err != nil {
		log.Fatal(err)
	}

	err = db.Create(&deletePermission).Error

	if err != nil {
		log.Fatal(err)
	}

	err = db.Create(&updatePermission).Error

	if err != nil {
		log.Fatal(err)
	}

	// Assign permissions to roles
	assignPermissions(db, adminRole.ID, readPermission.ID)
	assignPermissions(db, adminRole.ID, writePermission.ID)
	assignPermissions(db, adminRole.ID, deletePermission.ID)
	assignPermissions(db, adminRole.ID, updatePermission.ID)
	assignPermissions(db, normalRole.ID, readPermission.ID)
	assignPermissions(db, normalRole.ID, writePermission.ID)
	assignPermissions(db, normalRole.ID, deletePermission.ID)
	assignPermissions(db, normalRole.ID, updatePermission.ID)
	assignPermissions(db, consultantRole.ID, readPermission.ID)

	fmt.Println("Seeder executed successfully")
}

// Helper function to assign permissions to roles
func assignPermissions(db *gorm.DB, roleID uint, permissionID uint) {
	rolePermission := RolePermission{
		ID:           uuid.New(),
		RoleID:       roleID,
		PermissionID: permissionID,
		CreatedBy:    "seeder",
		UpdatedBy:    "seeder",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := db.Create(&rolePermission).Error
	if err != nil {
		log.Fatal(err)
	}
}
