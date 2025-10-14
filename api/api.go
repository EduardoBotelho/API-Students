package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/EduardoBotelho/API-STUDENTS/db"

)

// Estrutura da API refatorada para incluir o handler
type API struct {
	Echo           *echo.Echo
	DB *db.StudentsHandler
}

// NewServer cria uma nova instância da API
func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

// ConfigureRoutes configura todas as rotas da API
func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

// Start inicia o servidor
func (api *API) Start() error {
	return api.Echo.Start(":8080")
}
