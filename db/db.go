package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"

	"github.com/EduardoBotelho/API-STUDENTS/schemas"
)

type StudentsHandler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initiate Sqlite: %s", err)
	}
	if err != nil {
		// Handle the error, e.g., log it and exit
		panic("failed to connect database")
	}
	// Now 'db' is your GORM database object, ready for operations

	if err := db.AutoMigrate(&schemas.Student{}); err != nil {
		log.Fatal().Err(err).Msgf("Failed:%s", err)
	}
	return db
}

func NewStudentHandler(db *gorm.DB) *StudentsHandler {
	return &StudentsHandler{DB: db}
}

func (s *StudentsHandler) AddStudent(student schemas.Student) error {

	if result := s.DB.Create(&student); result.Error != nil {
		fmt.Println("Error creating student:", result.Error)
		log.Error().Msg("Failed to create student")
		return result.Error
	}

	fmt.Println("Created student:", student.ID)
	return nil
}
func (s *StudentsHandler) GetStudents() ([]schemas.Student, error) { //Busca todos os estudantes no banco
	students := []schemas.Student{}

	err := s.DB.Find(&students).Error

	return students, err
}

func (s *StudentsHandler) GetStudent(id int) (schemas.Student, error) { //Busca todos os estudantes no banco
	var student schemas.Student
	err := s.DB.First(&student, id)
	return student, err.Error
}
func (s *StudentsHandler) UpdateStudent(updateStudent schemas.Student) error { //Busca todos os estudantes no banco
	return s.DB.Save(&updateStudent).Error
}

func (s *StudentsHandler) DeleteStudent(student schemas.Student) error { //Busca todos os estudantes no banco
	return s.DB.Delete(&student).Error
}
