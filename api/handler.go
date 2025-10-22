package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/EduardoBotelho/API-STUDENTS/schemas"

)

//createStudent cria um novo estudante
func (api *API) createStudent(c echo.Context) error {
	student := schemas.Student{}
	if err := c.Bind(&student); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	return c.JSON(http.StatusCreated, student)
}

// getStudents busca todos os estudantes
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "No students found")
	}
	return c.JSON(http.StatusOK, students)
}

// getStudent busca um estudante por ID
func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Student ID")
	}

	student, err := api.DB.GetStudent(id)
	//nao encontrar o student  com esse id-Status not found(404)
	//ou pode ter algum problema para encontrar o student
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Student")
	}

	return c.JSON(http.StatusOK, student)
}

// updateStudent atualiza um estudante por ID
func (api *API) updateStudent(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}
	receivedStudent := schemas.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}

	updatingStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	student := updateStudentInfo(receivedStudent, updatingStudent)
	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save student")
	}
	return c.JSON(http.StatusOK, student)
}

// deleteStudent deleta um estudante por ID
func (api *API) deleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Id for delete")
	}
	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err := api.DB.DeleteStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to Delete student")
	}

	return c.JSON(http.StatusOK, student)
}

func updateStudentInfo(receivedStudent, student schemas.Student) schemas.Student {
	if receivedStudent.Name != "" {
		student.Name = receivedStudent.Name
	}
	if receivedStudent.Email != "" {
		student.Email = receivedStudent.Email
	}
	cpfInt, err := strconv.Atoi(receivedStudent.CPF)
	if err != nil {
		panic(err.Error())
	}
	if cpfInt > 0 {
		student.CPF = receivedStudent.CPF
	}
	if receivedStudent.Age > 0 {
		student.Age = receivedStudent.Age
	}
	if receivedStudent.Active != student.Active {
		student.Active = receivedStudent.Active
	}
	return student
}
