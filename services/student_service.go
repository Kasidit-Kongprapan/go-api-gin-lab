package services

import (
	"database/sql"
	"errors"

	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}
var (
    ErrStudentNotFound = errors.New("student not found")
    ErrInvalidInput    = errors.New("invalid input")
)

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
    student, err := s.Repo.GetByID(id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrStudentNotFound
        }
        return nil, err
    }
    return student, nil
}

func (s *StudentService) CreateStudent(student models.Student) error {
	if err := validateStudent(student, true); err != nil {
		return err
	}
	return s.Repo.Create(student)
}
func (s *StudentService) UpdateStudent(id string, student models.Student) error {
	if err := validateStudent(student, false); err != nil {
		return err
	}

	err := s.Repo.Update(id, student)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrStudentNotFound
		}
		return err
	}

	return nil
}
func (s *StudentService) DeleteStudent(id string) error {
    err := s.Repo.Delete(id)
    if err != nil {
        if err == sql.ErrNoRows {
            return ErrStudentNotFound
        }
        return err
    }
    return nil
}
func validateStudent(s models.Student, isCreate bool) error {
	if isCreate && s.Id == "" {
		return errors.New("id is required")
	}
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.GPA < 0.0 || s.GPA > 4.0 {
		return errors.New("gpa must be between 0.00 and 4.00")
	}
	return nil
}