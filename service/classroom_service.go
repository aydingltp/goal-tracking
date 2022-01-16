package service

import (
	"goal-tracking/models"
	"gorm.io/gorm"
)

type ClassroomService struct {
	db *gorm.DB
}

func NewClassroomService(db *gorm.DB) *ClassroomService {
	return &ClassroomService{db: db}
}

func (s *ClassroomService) Create(m *models.Classroom) error {
	err := s.db.Create(m).Error
	return err
}

func (s *ClassroomService) GetAll() ([]models.Classroom, error) {
	var classrooms []models.Classroom
	err := s.db.Find(&classrooms).Error
	return classrooms, err
}
func (s *ClassroomService) GetById(id int64) (*models.Classroom, error) {
	classroom := new(models.Classroom)
	err := s.db.Where(`id = ?`, id).First(&classroom).Error
	return classroom, err
}

func (s *ClassroomService) Delete(id int64) error {
	var classroom models.Classroom
	s.db.First(&classroom, id)

	err := s.db.Delete(&classroom).Error
	return err
}
