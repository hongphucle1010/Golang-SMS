package repository

import (
	"sms/app/model"
	"sms/app/utils"
)

type IStudentRepository interface {
	GetStudents() ([]model.Student, error)
}

type MongoDBStudentRepository struct {}

type JsonStudentRepository struct {}

func NewJsonStudentRepository() IStudentRepository {
	return &JsonStudentRepository{}
}

func (r *JsonStudentRepository) GetStudents() ([]model.Student, error) {
	students, err := utils.ReadJsonFile[model.Student]("data/minidata.json")
	if err != nil {
		return []model.Student{}, err
	}

	return students, nil
}

