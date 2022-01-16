package service

import (
	"errors"
	"goal-tracking/models"
	"goal-tracking/models/viewModels"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Login(m *viewModels.LoginVm) (*models.User, error) {
	user := new(models.User)
	err := s.db.Where("email = ? and password = ?", m.EMail, m.Password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return user, err
}

func (s *UserService) Create(m *models.User) error {
	err := s.db.Create(m).Error
	return err
}

func (s *UserService) GetAll() ([]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	return users, err
}
func (s *UserService) GetById(id int64) (*models.User, error) {
	user := new(models.User)
	err := s.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}
func (s *UserService) Delete(id int64) error {
	var user models.User
	s.db.First(&user, id)

	err := s.db.Delete(&user).Error
	return err
}
