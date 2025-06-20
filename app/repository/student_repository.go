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

	// students := []model.Student{
	// 	{
	// 		ID:    221499,
	// 		Name:  "Susan Wilson",
	// 		Email: "susan.wilson@example.com",
	// 		Dob:   "2001-11-20T00:00:00Z",
	// 		Gpa:   3.99,
	// 	},
	// 	{
	// 		ID:    221897,
	// 		Name:  "Marilyn Spencer",
	// 		Email: "marilyn.spencer@example.com",
	// 		Dob:   "2001-03-09T00:00:00Z",
	// 		Gpa:   3.97,
	// 	},
	// 	{
	// 		ID:    221465,
	// 		Name:  "Cheryl Henry",
	// 		Email: "cheryl.henry@example.com",
	// 		Dob:   "2001-03-29T00:00:00Z",
	// 		Gpa:   3.96,
	// 	},
	// }

	return students, nil
}

