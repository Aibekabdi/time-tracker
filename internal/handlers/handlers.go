package handlers

import (
	"log/slog"

	task_handler "github.com/Aibekabdi/time-tracker/internal/handlers/task"
	task_repo "github.com/Aibekabdi/time-tracker/internal/repository/postgres/task"
	task_service "github.com/Aibekabdi/time-tracker/internal/service/task"

	user_handler "github.com/Aibekabdi/time-tracker/internal/handlers/user"
	user_repo "github.com/Aibekabdi/time-tracker/internal/repository/postgres/user"
	user_service "github.com/Aibekabdi/time-tracker/internal/service/user"
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

func (h *Handler) InitRoutes(pool *pgxpool.Pool) *gin.Engine {
	router := gin.New()

	// User
	userRepo := user_repo.NewRepository(pool)
	userService := user_service.NewService(userRepo, h.apiAddress)
	userHandler := user_handler.New(h.log, h.apiAddress, userService)
	userGroup := router.Group("/user")
	{
		userGroup.GET("/get", userHandler.Get)
		userGroup.POST("/create", userHandler.Create)
		userGroup.PUT("/update/:user_id", userHandler.Update)
		userGroup.DELETE("/delete/:user_id", userHandler.Delete)
	}

	// Task
	taskRepo := task_repo.NewRepository(pool)
	taskService := task_service.NewService(taskRepo)
	taskHandler := task_handler.New(h.log, taskService)
	taskGroup := router.Group("/task")
	{
		taskGroup.GET("/get/:user_id", taskHandler.GetALL)
		taskGroup.GET("/get/:task_id", taskHandler.Get)
		taskGroup.POST("/start", taskHandler.Create)
		taskGroup.PATCH("/end/:task_id", taskHandler.End)
	}
	return router
}
