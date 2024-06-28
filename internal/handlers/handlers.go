package handlers

import (
	"log/slog"

	task_handler "github.com/Aibekabdi/time-tracker/internal/handlers/task"
	user_handler "github.com/Aibekabdi/time-tracker/internal/handlers/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	log        *slog.Logger
	apiAddress string
	// db
	// env level
}

func NewHandler(log *slog.Logger, apiAddress string) *Handler {
	return &Handler{
		log:        log,
		apiAddress: apiAddress,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// User
	userHandler := user_handler.New(h.log, h.apiAddress)
	userGroup := router.Group("/user")
	{
		userGroup.GET("/get", userHandler.GetALL)
		userGroup.GET("/get/:user_id", userHandler.Get)
		userGroup.POST("/create", userHandler.Create)
		userGroup.PUT("/update/:user_id", userHandler.Update)
		userGroup.DELETE("/delete/:user_id", userHandler.Delete)
	}

	// Task
	taskHandler := task_handler.New(h.log)
	taskGroup := router.Group("/task")
	{
		taskGroup.GET("/get/:user_id", taskHandler.Get)
		taskGroup.POST("/start", taskHandler.Create)
		taskGroup.PATCH("/end/:task_id", taskHandler.End)
	}
	return router
}
