package repository

import (
	"app/entity"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type RepositoryNode struct {
	db *gorm.DB
}

func NewRepositoryNode(db *gorm.DB) *RepositoryNode {
	return &RepositoryNode{
		db: db,
	}
}

func convertNodeToNodePayload(node *entity.EntityNode) *entity.EntityNodePayload {
	var position json.RawMessage
	var positionAbsolute json.RawMessage
	var style json.RawMessage
	err := json.Unmarshal([]byte(string(node.Position)), &position)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal([]byte(string(node.PositionAbsolute)), &positionAbsolute)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal([]byte(string(*node.Style)), &style)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &entity.EntityNodePayload{
		ID:               node.ID,
		Position:         position,
		Data:             node.Data.Local().Format("2006-01-02"),
		Type:             node.Type,
		SourcePosition:   node.SourcePosition,
		TargetPosition:   node.TargetPosition,
		Hidden:           node.Hidden,
		Selected:         node.Selected,
		Dragging:         node.Dragging,
		Draggable:        node.Draggable,
		Selectable:       node.Selectable,
		Connectable:      node.Connectable,
		Resizing:         node.Resizing,
		Deletable:        node.Deletable,
		DragHandle:       node.DragHandle,
		Width:            node.Width,
		Height:           node.Height,
		ParentNodeID:     node.ParentNodeID,
		ZIndex:           node.ZIndex,
		Extent:           node.Extent,
		ExpandParent:     node.ExpandParent,
		PositionAbsolute: positionAbsolute,
		AriaLabel:        node.AriaLabel,
		Focusable:        node.Focusable,
		Style:            style,
		ClassName:        node.ClassName,
	}
}

func (n *RepositoryNode) GetAll() ([]entity.EntityNodePayload, error) {
	var nodes []entity.EntityNode
	err := n.db.Find(&nodes).Error
	var payloads []entity.EntityNodePayload
	for _, node := range nodes {
		payloads = append(payloads, *convertNodeToNodePayload(&node))
	}
	return payloads, err
}

func (n *RepositoryNode) GetByID(id string) (*entity.EntityNodePayload, error) {
	var node *entity.EntityNode
	err := n.db.Where("id = ?", id).First(&node).Error
	payload := convertNodeToNodePayload(node)
	return payload, err
}

func (n *RepositoryNode) Create(node *entity.EntityNode) (*entity.EntityNodePayload, error) {
	node.New()
	err := n.db.Create(&node).Error
	if err != nil {
		return nil, err
	}
	return n.GetByID(node.ID)
}

func (n *RepositoryNode) Update(node *entity.EntityNode) (*entity.EntityNodePayload, error) {
	err := n.db.Save(&node).Error
	if err != nil {
		return nil, err
	}
	return n.GetByID(node.ID)
}

func (n *RepositoryNode) Delete(id string) error {
	node, err := n.GetByID(id)
	if err != nil {
		return err
	}
	err = n.db.Delete(node).Error
	return err
}
