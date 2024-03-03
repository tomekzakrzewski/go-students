package db

import (
	"github.com/tomekzakrzewski/go-students/models"
	"gorm.io/gorm"
)

type StudentStoreInterface interface {
	Insert(student *models.Student) (*models.Student, error)
	FindAll() (*[]models.Student, error)
	FindById(id int) (*models.Student, error)
}

type StudentStore struct {
	db *gorm.DB
}

func NewStudentStore(db *gorm.DB) *StudentStore {
	return &StudentStore{
		db: db,
	}
}

func (s *StudentStore) Insert(student *models.Student) (*models.Student, error) {
	if err := s.db.Create(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentStore) FindAll() (*[]models.Student, error) {
	var students *[]models.Student
	if err := s.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (s *StudentStore) FindById(id int) (*models.Student, error) {
	var student *models.Student
	if err := s.db.Where("id = ?", id).First(&student).Error; err != nil {
		return nil, err
	}

	return student, nil
}

func (s *StudentStore) Update(id int, student models.Student) error {
	var updateStudent models.Student
	result := s.db.Model(&updateStudent).Where("id = ?", id).Updates(student)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *StudentStore) Delete(id int) error {
	var deletedStudent models.Student
	result := s.db.Where("id = ?", id).Delete(&deletedStudent)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
