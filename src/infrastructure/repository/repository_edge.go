package repository

import (
	"app/entity"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type RepositoryEdge struct {
	db *gorm.DB
}

func NewRepositoryEdge(db *gorm.DB) *RepositoryEdge {
	return &RepositoryEdge{
		db: db,
	}
}

func convertEdgeToEdgePayload(edge *entity.EntityEdge) *entity.EntityEdgePayload {
	var style json.RawMessage
	var data json.RawMessage
	
	err := json.Unmarshal([]byte(string(*edge.Style)), &style)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal([]byte(string(edge.Data)), &data)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &entity.EntityEdgePayload{
		ID:               		edge.ID,
		NodeID:    		  		edge.NodeID,
		Node:      		  		edge.Node,
		Animated:         		edge.Animated,
		Data:             		data,
		Style:   		  		style,
		Selected:         		edge.Selected,
		Source:           		edge.Source,
		Target:        	  		edge.Target,
		SourceHandleID:   		edge.SourceHandleID,
		TargetHandleID:   		edge.TargetHandleID,
		InteractionWidth: 		edge.InteractionWidth,
		SourceX:          		edge.SourceX,
		SourceY:       	  		edge.SourceY,
		TargetX:          		edge.TargetX,
		TargetY:          		edge.TargetY,
		SourcePosition:   		edge.SourcePosition,
		TargetPosition:   		edge.TargetPosition,
		Label:     		  		edge.Label,
		LabelStyle:       		edge.LabelStyle,
		LabelShowBg:      		edge.LabelShowBg,
		LabelBgStyle:     		edge.LabelBgStyle,
		LabelBgPadding:   		edge.LabelBgPadding,
		LabelBgBorderRadius:    edge.LabelBgBorderRadius,
		MarkerStart:    		edge.MarkerStart,
		MarkerEnd:    			edge.MarkerEnd,
		PathOptions:    		edge.PathOptions,
	}
}

func (n *RepositoryEdge) GetAll(node_id *string) (payloads []entity.EntityEdgePayload, err error) {
	var edges []entity.EntityEdge
	
	if node_id != nil {
		err = n.db.Where("node_id = ? is null", node_id).Find(&edges).Error
	} 

	if err != nil {
		return payloads, err
	}

	for _, edge := range edges {
		payloads = append(payloads, *convertEdgeToEdgePayload(&edge))
	}

	if len(payloads) == 0 {
		payloads = []entity.EntityEdgePayload{}
	}

	return payloads, err
}

func (n *RepositoryEdge) GetByID(id string) (*entity.EntityEdgePayload, error) {
	var edge *entity.EntityEdge
	err := n.db.Where("id = ?", id).First(&edge).Error
	if err != nil {
		return nil, err
	}
	payload := convertEdgeToEdgePayload(edge)
	return payload, err
}

func (n *RepositoryEdge) Create(edge *entity.EntityEdge) (*entity.EntityEdgePayload, error) {
	edge.New()
	err := n.db.Create(&edge).Error
	if err != nil {
		return nil, err
	}
	return n.GetByID(edge.ID)
}

func (n *RepositoryEdge) Update(edge *entity.EntityEdge) (*entity.EntityEdgePayload, error) {
	err := n.db.Save(&edge).Error
	if err != nil {
		return nil, err
	}
	return n.GetByID(edge.ID)
}

func (n *RepositoryEdge) Delete(id string) error {
	edge, err := n.GetByID(id)
	if err != nil {
		return err
	}
	err = n.db.Delete(edge).Error
	return err
}