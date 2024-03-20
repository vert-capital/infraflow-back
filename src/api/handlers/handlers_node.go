package handlers

import (
	"app/entity"
	"app/infrastructure/repository"
	usecase_node "app/usecase/node"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NodeHandlers struct {
	UcNode usecase_node.IUsecaseNode
}

func NewNodeHandler(ucNode usecase_node.IUsecaseNode) *NodeHandlers {
	return &NodeHandlers{UcNode: ucNode}
}

func (h *NodeHandlers) GetAllHandler(c *gin.Context) {
	node, err := h.UcNode.GetAll()
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, node)

}

func (h *NodeHandlers) GetByIDHandler(c *gin.Context) {
	id := c.Param("id")
	node, err := h.UcNode.GetByID(id)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, node)
}

func (h *NodeHandlers) CreateHandler(c *gin.Context) {
	var payload entity.EntityNodePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		handleError(c, err)
		return
	}
	data, err := time.Parse("2006-01-02", payload.Data)
	if err != nil {
		handleError(c, err)
		return
	}
	positionJSON, _ := payload.Position.MarshalJSON()
	positionAbsoluteJson, _ := payload.PositionAbsolute.MarshalJSON()
	styleJson, _ := payload.Style.MarshalJSON()
	style := string(styleJson)

	entityNode := entity.EntityNode{
		ID:               payload.ID,
		Position:         string(positionJSON),
		Data:             data,
		Type:             payload.Type,
		SourcePosition:   payload.SourcePosition,
		TargetPosition:   payload.TargetPosition,
		Hidden:           payload.Hidden,
		Selected:         payload.Selected,
		Dragging:         payload.Dragging,
		Draggable:        payload.Draggable,
		Selectable:       payload.Selectable,
		Connectable:      payload.Connectable,
		Resizing:         payload.Resizing,
		Deletable:        payload.Deletable,
		DragHandle:       payload.DragHandle,
		Width:            payload.Width,
		Height:           payload.Height,
		ParentNodeID:     payload.ParentNodeID,
		ZIndex:           payload.ZIndex,
		Extent:           payload.Extent,
		ExpandParent:     payload.ExpandParent,
		PositionAbsolute: string(positionAbsoluteJson),
		AriaLabel:        payload.AriaLabel,
		Focusable:        payload.Focusable,
		Style:            &style,
		ClassName:        payload.ClassName,
	}
	node, err := h.UcNode.Create(&entityNode)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, node)
}

func (h *NodeHandlers) UpdateHandler(c *gin.Context) {
	var entityNode entity.EntityNode
	id := c.GetString("id")
	entityNode.ID = id
	if err := c.ShouldBindJSON(&entityNode); err != nil {
		handleError(c, err)
		return
	}
	node, err := h.UcNode.Create(&entityNode)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, node)
}

func (h *NodeHandlers) DeleteHandler(c *gin.Context) {
	id := c.GetString("id")
	err := h.UcNode.Delete(id)
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

func MountNodeHandlers(gin *gin.Engine, conn *gorm.DB) {
	nodeHandlers := NewNodeHandler(
		usecase_node.NewUseCaseNode(
			repository.NewRepositoryNode(conn),
		),
	)
	group := gin.Group("/api/node")
	// SetAuthMiddleware(conn, group)
	group.GET("/", nodeHandlers.GetAllHandler)
	group.POST("/", nodeHandlers.CreateHandler)
	group.GET("/:id", nodeHandlers.GetByIDHandler)
	group.PUT("/:id", nodeHandlers.UpdateHandler)
	group.DELETE("/:id", nodeHandlers.DeleteHandler)
}
