package repository

import (
	"fmt"
	"sms/app/model"
	"sms/app/utils"

	"github.com/spf13/viper"
)

type IStudentRepository interface {
	GetStudents() ([]model.Student, error)
	AddStudent(req model.CreateStudentRequest) (model.Student, error)
	UpdateStudent(student model.Student) (model.Student, error)
	DeleteStudent(id int) error
	GetStudentByID(id int) (model.Student, error)
}

type MongoDBStudentRepository struct{}

type JsonStudentRepository struct{}

func NewJsonStudentRepository() IStudentRepository {
	return &JsonStudentRepository{}
}

func (r *JsonStudentRepository) GetStudents() ([]model.Student, error) {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return []model.Student{}, err
	}

	return students, nil
}

func (r *JsonStudentRepository) AddStudent(req model.CreateStudentRequest) (model.Student, error) {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return model.Student{}, err
	}
	student := model.Student{
		ID:    utils.GenerateNextID(students, func(s model.Student) int { return s.ID }),
		Name:  req.Name,
		Email: req.Email,
		Dob:   req.Dob,
		Gpa:   req.Gpa,
	}
	students = append(students, student)
	return student, utils.WriteJsonFile(viper.GetString("database.path"), students)
}

func (r *JsonStudentRepository) UpdateStudent(student model.Student) (model.Student, error) {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return model.Student{}, err
	}
	updated := false
	for i, s := range students {
		if s.ID == student.ID {
			students[i] = student
			updated = true
			break
		}
	}
	if !updated {
		return model.Student{}, fmt.Errorf("student with ID %d not found", student.ID)
	}
	return student, utils.WriteJsonFile(viper.GetString("database.path"), students)
}

func (r *JsonStudentRepository) DeleteStudent(id int) error {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return err
	}
	newStudents := make([]model.Student, 0, len(students))
	deleted := false
	for _, s := range students {
		if s.ID == id {
			deleted = true
			continue
		}
		newStudents = append(newStudents, s)
	}
	if !deleted {
		return fmt.Errorf("student with ID %d not found", id)
	}
	return utils.WriteJsonFile(viper.GetString("database.path"), newStudents)
}

func (r *JsonStudentRepository) GetStudentByID(id int) (model.Student, error) {
	students, err := utils.ReadJsonFile[model.Student](viper.GetString("database.path"))
	if err != nil {
		return model.Student{}, err
	}
	for _, s := range students {
		if s.ID == id {
			return s, nil
		}
	}
	return model.Student{}, fmt.Errorf("student with ID %d not found", id)
}
