package service

import (
	"sms/app/model"
	"sms/app/repository"
)

type IStudentService interface {
	GetAllStudents() ([]model.Student, error)
}

type StudentService struct {
	repository repository.IStudentRepository
}

func NewStudentService(repository repository.IStudentRepository) IStudentService {
	return &StudentService{repository}
}

func (s *StudentService) GetAllStudents() ([]model.Student, error) {
	return s.repository.GetStudents()
}