package repository

import (
	"backend/internal/domain/models"
	"fmt"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	Name  string
	Email string
}

func (UserDB) TableName() string {
	return "users"
}

type UserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	FindUser(id uint) (*models.User, error)
}

type gormUserRepository struct {
	db *gorm.DB
}

// MapToDB: User -> UserDB への変換
func UserRepoMapToDB(u *models.User) (*UserDB, error) {
	dbUser := &UserDB{}
	err := copier.Copy(dbUser, u)
	return dbUser, err
}

// MapToDomain: UserDB -> User への変換
func UserRepoMapToDomain(dbU *UserDB) (*models.User, error) {
	user := &models.User{}
	err := copier.Copy(user, dbU)
	return user, err
}

// MapToDBArray: []User -> []UserDB への変換
func UserRepoMapToDBArray(us []models.User) ([]UserDB, error) {
	var dbUsers []UserDB
	err := copier.Copy(&dbUsers, us)
	return dbUsers, err
}

// MapToDomainArray: []UserDB -> []User への変換
func UserRepoMapToDomainArray(dbUs []UserDB) ([]models.User, error) {
	var users []models.User
	err := copier.Copy(&users, dbUs)
	return users, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Create(user *models.User) error {
	dbuser, err := UserRepoMapToDB(user)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}
	return r.db.Create(dbuser).Error
}

func (r *gormUserRepository) GetAll() ([]models.User, error) {
	var dbusers []UserDB
	err := r.db.Find(&dbusers).Error
	if err != nil {
		return nil, err
	}

	usersPtr, err := UserRepoMapToDomainArray(dbusers)
	if err != nil {
		return nil, err
	}

	return usersPtr, nil
}

func (r *gormUserRepository) FindUser(id uint) (*models.User, error) {
	var dbusers UserDB
	err := r.db.Where("id = ?", id).First(&dbusers).Error
	if err != nil {
		return nil, err
	}

	userPtr, err := UserRepoMapToDomain(&dbusers)
	if err != nil {
		return nil, err
	}

	return userPtr, nil
}
