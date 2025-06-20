package service

import "sms/app/model"

type IStudentService interface {
	GetAllStudents() ([]model.Student, error)
}

type StudentService struct{}

func (s *StudentService) GetAllStudents() ([]model.Student, error) {
	return []model.Student{
		{
			ID: 1,
			Name: "John Doe",
			Email: "john.doe@example.com",
		},
		{
			ID: 2,
			Name: "Jane Doe",
			Email: "jane.doe@example.com",
		},
	}, nil
}