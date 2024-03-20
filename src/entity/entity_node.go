package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Position string
type CoordinateExtent string

type EntityNode struct {
	ID               string            `json:"id"`
	Position         string            `json:"position"`
	Data             time.Time         `json:"data"`
	Type             string            `json:"type,omitempty"`
	SourcePosition   Position          `json:"source_position,omitempty"`
	TargetPosition   Position          `json:"target_position,omitempty"`
	Hidden           bool              `json:"hidden,omitempty"`
	Selected         bool              `json:"selected,omitempty"`
	Dragging         bool              `json:"dragging,omitempty"`
	Draggable        bool              `json:"draggable,omitempty"`
	Selectable       bool              `json:"selectable,omitempty"`
	Connectable      bool              `json:"connectable,omitempty"`
	Resizing         bool              `json:"resizing,omitempty"`
	Deletable        bool              `json:"deletable,omitempty"`
	DragHandle       string            `json:"drag_handle,omitempty"`
	Width            *float64          `json:"width,omitempty"`
	Height           *float64          `json:"height,omitempty"`
	ParentNodeID     *string           `json:"parent_node_id,omitempty"`
	ParentNode       *EntityNode       `json:"parent_node,omitempty"`
	ZIndex           int               `json:"z_index,omitempty"`
	Extent           *CoordinateExtent `json:"extent,omitempty"`
	ExpandParent     bool              `json:"expand_parent,omitempty"`
	PositionAbsolute string            `json:"position_absolute,omitempty"`
	AriaLabel        string            `json:"aria_label,omitempty"`
	Focusable        bool              `json:"focusable,omitempty"`
	Style            *string           `json:"style,omitempty"`
	ClassName        string            `json:"class_name,omitempty"`
}

type EntityNodePayload struct {
	ID               string            `json:"id"`
	Position         json.RawMessage   `json:"position"`
	Data             string            `json:"data"`
	Type             string            `json:"type,omitempty"`
	SourcePosition   Position          `json:"source_position,omitempty"`
	TargetPosition   Position          `json:"target_position,omitempty"`
	Hidden           bool              `json:"hidden,omitempty"`
	Selected         bool              `json:"selected,omitempty"`
	Dragging         bool              `json:"dragging,omitempty"`
	Draggable        bool              `json:"draggable,omitempty"`
	Selectable       bool              `json:"selectable,omitempty"`
	Connectable      bool              `json:"connectable,omitempty"`
	Resizing         bool              `json:"resizing,omitempty"`
	Deletable        bool              `json:"deletable,omitempty"`
	DragHandle       string            `json:"drag_handle,omitempty"`
	Width            *float64          `json:"width,omitempty"`
	Height           *float64          `json:"height,omitempty"`
	ParentNodeID     *string           `json:"parent_node_id,omitempty"`
	ZIndex           int               `json:"z_index,omitempty"`
	Extent           *CoordinateExtent `json:"extent,omitempty"`
	ExpandParent     bool              `json:"expand_parent,omitempty"`
	PositionAbsolute json.RawMessage   `json:"position_absolute,omitempty"`
	AriaLabel        string            `json:"aria_label,omitempty"`
	Focusable        bool              `json:"focusable,omitempty"`
	Style            json.RawMessage   `json:"style,omitempty"`
	ClassName        string            `json:"class_name,omitempty"`
}

func (n *EntityNode) New() {
	uuid, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	n.ID = uuid.String()
}
