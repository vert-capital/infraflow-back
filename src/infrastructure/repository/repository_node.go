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

func convertNodeToNodePayload(node *entity.EntityNode, childrens []entity.EntityNodePayload) *entity.EntityNodePayload {
	var position json.RawMessage
	var positionAbsolute json.RawMessage
	var style json.RawMessage
	var data json.RawMessage

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
	err = json.Unmarshal([]byte(string(node.Data)), &data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &entity.EntityNodePayload{
		ID:               node.ID,
		ApplicationID:    node.ApplicationID,
		Application:      node.Application,
		Position:         position,
		Data:			  data,
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
		Childrens:        childrens,
		Label:			  node.Label,
	}
}

func (n *RepositoryNode) GetAll(application_id *string) (payloads []entity.EntityNodePayload, err error) {
	var nodes []entity.EntityNode

	if application_id != nil {
		err = n.db.Where("application_id = ? and parent_node_id is null", application_id).Find(&nodes).Error
	} else {
		err = n.db.Where("parent_node_id is null").Find(&nodes).Error
	}
	if err != nil {
		return payloads, err
	}
	for _, node := range nodes {
		childrens, err := n.GetNodeChildrens(node.ID)
		if err != nil {
			return payloads, err
		}
		payloads = append(payloads, *convertNodeToNodePayload(&node, childrens))
	}

	if len(payloads) == 0 {
		payloads = []entity.EntityNodePayload{}
	}
	return payloads, err
}

func (n *RepositoryNode) GetNodeChildrens(parent_id string) (payloads []entity.EntityNodePayload, err error) {
	var nodes []entity.EntityNode
	err = n.db.Where("parent_node_id = ?", parent_id).Find(&nodes).Error
	if err != nil {
		return payloads, err
	}
	for _, node := range nodes {
		childrens, err := n.GetNodeChildrens(node.ID)
		if err != nil {
			return payloads, err
		}
		payloads = append(payloads, *convertNodeToNodePayload(&node, childrens))
	}
	return payloads, err

}

func (n *RepositoryNode) GetByID(id string) (*entity.EntityNodePayload, error) {
	var node *entity.EntityNode
	err := n.db.Where("id = ?", id).First(&node).Error
	if err != nil {
		return nil, err
	}
	childrens, err := n.GetNodeChildrens(id)
	if err != nil {
		return nil, err
	}
	payload := convertNodeToNodePayload(node, childrens)
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
