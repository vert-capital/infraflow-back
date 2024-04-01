package handlers

import (
	"app/entity"
	"app/infrastructure/repository"
	usecase_edge "app/usecase/edge"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EdgeHandlers struct {
	UcEdge usecase_edge.IUsecaseEdge
}

func NewEdgeHandler(ucEdge usecase_edge.IUsecaseEdge) *EdgeHandlers {
	return &EdgeHandlers{UcEdge: ucEdge}
}

func (h *EdgeHandlers) GetAllHandler(c *gin.Context) {
	var node_id *string = nil
	node_id_str, exists := c.GetQuery("node_id")
	if exists {
		node_id = &node_id_str
	}
	edge, err := h.UcEdge.GetAll(node_id)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, edge)

}

func (h *EdgeHandlers) GetByIDHandler(c *gin.Context) {
	id := c.Param("id")
	edge, err := h.UcEdge.GetByID(id)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, edge)
}

func (h *EdgeHandlers) CreateHandler(c *gin.Context) {
	var payload entity.EntityEdgePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		handleError(c, err)
		return
	}

	styleJson, _ := payload.Style.MarshalJSON()
	style := string(styleJson)
	dataJSON, _ := payload.Data.MarshalJSON()

	entityEdge := entity.EntityEdge{
		ID:               	payload.ID,
		NodeID:	   	   		payload.NodeID,
		Animated: 	   		payload.Animated,
		Data:		   		string(dataJSON),
		Style:		   		&style,
		Selected:	   		payload.Selected,
		Source: 	   	   	payload.Source,
		Target:     	   	payload.Target,
		SourceHandleID: 	payload.SourceHandleID,
		TargetHandleID: 	payload.TargetHandleID,
		InteractionWidth:  	payload.InteractionWidth,
		SourceX: 	   		payload.SourceX,
		SourceY: 	   		payload.SourceY,
		TargetX: 	   		payload.TargetX,
		TargetY: 	   		payload.TargetY,
		SourcePosition: 	payload.SourcePosition,
		TargetPosition: 	payload.TargetPosition,
		Label:		   		payload.Label,
		LabelStyle:	   		payload.LabelStyle,
		LabelShowBg:	   	payload.LabelShowBg,
		LabelBgStyle:   	payload.LabelBgStyle,
		LabelBgPadding: 	payload.LabelBgPadding,
		LabelBgBorderRadius:payload.LabelBgBorderRadius,
		MarkerStart:	   	payload.MarkerStart,
		MarkerEnd:	   		payload.MarkerEnd,
		PathOptions:		payload.PathOptions,
	}
	edge, err := h.UcEdge.Create(&entityEdge)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, edge)
}

func (h *EdgeHandlers) UpdateHandler(c *gin.Context) {
	var entityEdge entity.EntityEdge
	id := c.GetString("id")
	entityEdge.ID = id
	if err := c.ShouldBindJSON(&entityEdge); err != nil {
		handleError(c, err)
		return
	}
	edge, err := h.UcEdge.Update(&entityEdge)
	if err != nil {
		jsonResponse(c, http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	jsonResponse(c, http.StatusOK, edge)
}

func (h *EdgeHandlers) DeleteHandler(c *gin.Context) {
	id := c.GetString("id")
	err := h.UcEdge.Delete(id)
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

func MountEdgeHandlers(gin *gin.Engine, conn *gorm.DB) {
	edgeHandlers := NewEdgeHandler(
		usecase_edge.NewUseCaseEdge(
			repository.NewRepositoryEdge(conn),
		),
	)
	group := gin.Group("/api/edge")
	// SetAuthMiddleware(conn, group)
	group.GET("/", edgeHandlers.GetAllHandler)
	group.POST("/", edgeHandlers.CreateHandler)
	group.GET("/:id", edgeHandlers.GetByIDHandler)
	group.PUT("/:id", edgeHandlers.UpdateHandler)
	group.DELETE("/:id", edgeHandlers.DeleteHandler)
}