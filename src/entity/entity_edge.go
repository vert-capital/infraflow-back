package entity

import (
	"encoding/json"

	"github.com/google/uuid"
)

type EntityEdge struct {
	ID			   		string             	`json:"id"`
	NodeID	   	   		string 				`json:"node_id"`
	Node	   	   		*EntityNode 		`json:"node"`
	Animated 	   		bool               	`json:"animated"`
	Data		   		string             	`json:"data"`
	Style		   		*string            	`json:"style"`
	Selected	   		bool               	`json:"selected"`
	Source 	   	   		string             	`json:"source"`
	Target     	   		string             	`json:"target"`
	SourceHandleID 		*string            	`json:"source_handle_id"`
	TargetHandleID 		*string            	`json:"target_handle_id"`
	InteractionWidth  	int            		`json:"interaction_width"`
	SourceX 	   		float64            	`json:"source_x"`
	SourceY 	   		float64            	`json:"source_y"`
	TargetX 	   		float64            	`json:"target_x"`
	TargetY 	   		float64            	`json:"target_y"`
	SourcePosition 		Position           	`json:"source_position"`
	TargetPosition 		Position           	`json:"target_position"`
	Label		   		string             	`json:"label"`
	LabelStyle	   		*string            	`json:"label_style"`
	LabelShowBg	   		bool               	`json:"label_show_bg"`
	LabelBgStyle   		*string            	`json:"label_bg_style"`
	LabelBgPadding 		int            		`json:"label_bg_padding"`
	LabelBgBorderRadius int            		`json:"label_bg_border_radius"`
	MarkerStart	   		string             	`json:"marker_start"`
	MarkerEnd	   		string             	`json:"marker_end"`
	PathOptions	   		*string            	`json:"path_options"`
}

type EntityEdgePayload struct {
	ID			   		string             	`json:"id"`
	NodeID	   	   		string 				`json:"node_id"`
	Node	   	   		*EntityNode 		`json:"node"`
	Animated 	   		bool               	`json:"animated"`
	Data		   		json.RawMessage     `json:"data"`
	Style		   		json.RawMessage     `json:"style"`
	Selected	   		bool               	`json:"selected"`
	Source 	   	   		string             	`json:"source"`
	Target     	   		string             	`json:"target"`
	SourceHandleID 		*string            	`json:"source_handle_id"`
	TargetHandleID 		*string            	`json:"target_handle_id"`
	InteractionWidth  	int            		`json:"interaction_width"`
	SourceX 	   		float64            	`json:"source_x"`
	SourceY 	   		float64            	`json:"source_y"`
	TargetX 	   		float64            	`json:"target_x"`
	TargetY 	   		float64            	`json:"target_y"`
	SourcePosition 		Position           	`json:"source_position"`
	TargetPosition 		Position           	`json:"target_position"`
	Label		   		string             	`json:"label"`
	LabelStyle	   		*string            	`json:"label_style"`
	LabelShowBg	   		bool               	`json:"label_show_bg"`
	LabelBgStyle   		*string            	`json:"label_bg_style"`
	LabelBgPadding 		int            		`json:"label_bg_padding"`
	LabelBgBorderRadius int            		`json:"label_bg_border_radius"`
	MarkerStart	   		string             	`json:"marker_start"`
	MarkerEnd	   		string             	`json:"marker_end"`
	PathOptions	   		*string            	`json:"path_options"`
}

func (e *EntityEdge) New() {
	uuid, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	e.ID = uuid.String()
}

