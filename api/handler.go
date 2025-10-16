package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/EduardoBotelho/API-STUDENTS/db"

)

// getStudents busca todos os estudantes
func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "No students found")
	}
	return c.JSON(http.StatusOK, students)
}

// createStudent cria um novo estudante
func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	return c.JSON(http.StatusCreated, student)
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
	id := c.Param("id")
	updateStud := fmt.Sprintf("Update %s student", id)
	return c.String(http.StatusOK, updateStud)
}

// deleteStudent deleta um estudante por ID
func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delete %s student", id)
	return c.String(http.StatusOK, deleteStud)
}
