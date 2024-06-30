package handlers

import (
	"log/slog"

	task_handler "github.com/Aibekabdi/time-tracker/internal/handlers/task"
	user_handler "github.com/Aibekabdi/time-tracker/internal/handlers/user"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	log        *slog.Logger
	apiAddress string
}

func NewHandler(log *slog.Logger, apiAddress string) *Handler {
	return &Handler{
		log:        log,
		apiAddress: apiAddress,
	}
}

func (h *Handler) InitRoutes(db *pgxpool.Pool) *gin.Engine {
	router := gin.New()

	// User
	userHandler := user_handler.New(h.log, h.apiAddress)
	userGroup := router.Group("/user")
	{
		userGroup.GET("/get", userHandler.Get)
		userGroup.POST("/create", userHandler.Create)
		userGroup.PUT("/update/:user_id", userHandler.Update)
		userGroup.DELETE("/delete/:user_id", userHandler.Delete)
	}

	// Task
	taskHandler := task_handler.New(h.log)
	taskGroup := router.Group("/task")
	{
		taskGroup.GET("/get/:user_id", taskHandler.GetALL)
		taskGroup.GET("/get/:task_id", taskHandler.Get)
		taskGroup.POST("/start", taskHandler.Create)
		taskGroup.PATCH("/end/:task_id", taskHandler.End)
	}
	return router
}
