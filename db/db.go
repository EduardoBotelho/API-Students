package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentsHandler struct {
	DB *gorm.DB
}
type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initiate Sqlite: %s", err)
	}
	if err != nil {
		// Handle the error, e.g., log it and exit
		panic("failed to connect database")
	}
	// Now 'db' is your GORM database object, ready for operations

	if err := db.AutoMigrate(&Student{}); err != nil {
		log.Fatal().Err(err).Msgf("Failed", err)
	}
	return db
}

func NewStudentHandler(db *gorm.DB) *StudentsHandler {
	return &StudentsHandler{DB: db}
}

func (s *StudentsHandler) AddStudent(student Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		fmt.Println("Error creating student:", result.Error)
		log.Error().Msg("Failed to create student")
		return result.Error
	}

	fmt.Println("Created student:", student.ID)
	return nil
}
func (s *StudentsHandler) GetStudents() ([]Student, error) { //Busca todos os estudantes no banco
	students := []Student{}

	err := s.DB.Find(&students).Error

	return students, err
}

func (s *StudentsHandler) GetStudent(id int) (Student, error) { //Busca todos os estudantes no banco
	var student Student
	err := s.DB.First(&student, id)
	return student, err.Error
}
