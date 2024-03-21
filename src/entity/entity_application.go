package entity

import "github.com/google/uuid"

type EntityApplication struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (e *EntityApplication) New() {
	uuid, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	e.ID = uuid.String()
}
