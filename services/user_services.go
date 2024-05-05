package services

import (
	"github.com/phongtran11/go-project/database"
	"github.com/phongtran11/go-project/models"
	"github.com/phongtran11/go-project/pkg/constants"
	"github.com/phongtran11/go-project/pkg/dto/request"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const countFilterExp string = "count(*) > 0"

type TUserServices struct {
	database *gorm.DB
}

func UserServices() *TUserServices {
	return &TUserServices{
		database: database.GetDB(),
	}
}

func (userServices *TUserServices) Create(request *request.TRegisterRequest) error {
	user := models.User{FirstName: request.FirstName, LastName: request.LastName, Email: request.Email, Password: request.Password}

	// Check if email exists
	exists, err := userServices.existsByEmail(user.Email)

	if err != nil {
		return err
	}
	if exists {
		return &constants.TErrorMap{Message: constants.EmailExists}
	}

	// encrypt password
	encryptPassword := []byte(request.Password)
	hp, err := bcrypt.GenerateFromPassword(encryptPassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hp)

	// insert user transaction
	ctx := userServices.database.Begin()

	db := ctx.Create(&user)

	if db.Error != nil {
		ctx.Rollback()
		return db.Error
	}

	ctx.Commit()

	return nil
}

func (userServices *TUserServices) FindAll() ([]models.User, error) {
	var users []models.User

	db := userServices.database.Find(&users)

	if db.Error != nil {
		return nil, db.Error
	}

	return users, nil
}

func (userServices *TUserServices) FindByEmail(email string) (*models.User, error) {
	var user models.User

	db := userServices.database.Where("email = ?", email).First(&user)

	if db.Error != nil {
		return nil, db.Error
	}

	return &user, nil
}

func (s *TUserServices) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select(countFilterExp).
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}
