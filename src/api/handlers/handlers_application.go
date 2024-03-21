package handlers

import (
	"app/entity"
	"app/infrastructure/repository"
	usecase_application "app/usecase/application"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApplicationHandlers struct {
	UcApplication usecase_application.IUsecaseApplication
}

func NewApplicationHandler(ucApplication usecase_application.IUsecaseApplication) *ApplicationHandlers {
	return &ApplicationHandlers{UcApplication: ucApplication}
}

func (h *ApplicationHandlers) GetAllHandler(c *gin.Context) {
	apps, err := h.UcApplication.GetAll()
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, apps)
}

func (h *ApplicationHandlers) GetByIDHandler(c *gin.Context) {
	id := c.Param("id")
	apps, err := h.UcApplication.GetByID(id)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, apps)
}

func (h *ApplicationHandlers) CreateHandler(c *gin.Context) {
	var payload entity.EntityApplication
	if err := c.ShouldBindJSON(&payload); err != nil {
		handleError(c, err)
		return
	}
	apps, err := h.UcApplication.Create(payload)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, apps)
}

func (h *ApplicationHandlers) UpdateHandler(c *gin.Context) {
	var entityApplication entity.EntityApplication
	id := c.GetString("id")
	entityApplication.ID = id
	if err := c.ShouldBindJSON(&entityApplication); err != nil {
		handleError(c, err)
		return
	}
	node, err := h.UcApplication.Update(entityApplication)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, node)
}

func (h *ApplicationHandlers) DeleteHandler(c *gin.Context) {
	id := c.GetString("id")
	err := h.UcApplication.Delete(id)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusNoContent, gin.H{
		"status": "Deleted successfully",
	})
}

func MountApplicationHandlers(gin *gin.Engine, conn *gorm.DB) {
	applicationHandlers := NewApplicationHandler(
		usecase_application.NewUseCaseApplication(
			repository.NewRepositoryApplication(conn),
		),
	)
	group := gin.Group("/api/application")
	// SetAuthMiddleware(conn, group)
	group.GET("/", applicationHandlers.GetAllHandler)
	group.POST("/", applicationHandlers.CreateHandler)
	group.GET("/:id", applicationHandlers.GetByIDHandler)
	group.PUT("/:id", applicationHandlers.UpdateHandler)
	group.DELETE("/:id", applicationHandlers.DeleteHandler)
}
