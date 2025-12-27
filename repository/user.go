package repository

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

type UserRepository interface {
	GetUserByID(id uint) (*User, error)
	CreateUser(user *User) error
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	db.AutoMigrate(&User{})
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) GetUserByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}
