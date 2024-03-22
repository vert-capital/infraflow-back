package entity

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Position string
type CoordinateExtent string

type EntityNode struct {
	ID               string             `json:"id"`
	ApplicationID    string             `json:"application_id"`
	Application      *EntityApplication `json:"application"`
	Position         string             `json:"position"`
	Data             string             `json:"data"`
	Type             string             `json:"type"`
	SourcePosition   Position           `json:"source_position"`
	TargetPosition   Position           `json:"target_position"`
	Hidden           bool               `json:"hidden"`
	Selected         bool               `json:"selected"`
	Dragging         bool               `json:"dragging"`
	Draggable        bool               `json:"draggable"`
	Selectable       bool               `json:"selectable"`
	Connectable      bool               `json:"connectable"`
	Resizing         bool               `json:"resizing"`
	Deletable        bool               `json:"deletable"`
	DragHandle       string             `json:"drag_handle"`
	Width            *float64           `json:"width"`
	Height           *float64           `json:"height"`
	ParentNodeID     *string            `json:"parent_node_id"`
	ParentNode       *EntityNode        `json:"parent_node"`
	ZIndex           int                `json:"z_index"`
	Extent           *CoordinateExtent  `json:"extent"`
	ExpandParent     bool               `json:"expand_parent"`
	PositionAbsolute string             `json:"position_absolute"`
	AriaLabel        string             `json:"aria_label"`
	Focusable        bool               `json:"focusable"`
	Style            *string            `json:"style"`
	ClassName        string             `json:"class_name"`
	Label            string             `json:"label"`
}

type EntityNodePayload struct {
	ID               string              `json:"id"`
	ApplicationID    string              ` json:"application_id"`
	Application      *EntityApplication  ` json:"application"`
	Position         json.RawMessage     `json:"position"`
	Data             json.RawMessage     `json:"data"`
	Type             string              `json:"type"`
	SourcePosition   Position            `json:"source_position"`
	TargetPosition   Position            `json:"target_position"`
	Hidden           bool                `json:"hidden"`
	Selected         bool                `json:"selected"`
	Dragging         bool                `json:"dragging"`
	Draggable        bool                `json:"draggable"`
	Selectable       bool                `json:"selectable"`
	Connectable      bool                `json:"connectable"`
	Resizing         bool                `json:"resizing"`
	Deletable        bool                `json:"deletable"`
	DragHandle       string              `json:"drag_handle"`
	Width            *float64            `json:"width"`
	Height           *float64            `json:"height"`
	ParentNodeID     *string             `json:"parent_node_id"`
	ZIndex           int                 `json:"z_index"`
	Extent           *CoordinateExtent   `json:"extent"`
	ExpandParent     bool                `json:"expand_parent"`
	PositionAbsolute json.RawMessage     `json:"position_absolute"`
	AriaLabel        string              `json:"aria_label"`
	Focusable        bool                `json:"focusable"`
	Style            json.RawMessage     `json:"style"`
	ClassName        string              `json:"class_name"`
	Childrens        []EntityNodePayload `json:"childrens"`
	Label            string              `json:"label"`
}

func (n *EntityNode) New() {
	uuid, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	n.ID = uuid.String()
}
