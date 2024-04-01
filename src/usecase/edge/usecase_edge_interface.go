package usecase_edge

import "app/entity"

type IRepositoryEdge interface {
	GetAll(node_id *string) ([]entity.EntityEdgePayload, error)
	GetByID(id string) (*entity.EntityEdgePayload, error)
	Create(edge *entity.EntityEdge) (*entity.EntityEdgePayload, error)
	Update(edge *entity.EntityEdge) (*entity.EntityEdgePayload, error)
	Delete(id string) error
}

type IUsecaseEdge interface {
	GetAll(node_id *string) ([]entity.EntityEdgePayload, error)
	GetByID(id string) (*entity.EntityEdgePayload, error)
	Create(edge *entity.EntityEdge) (*entity.EntityEdgePayload, error)
	Update(edge *entity.EntityEdge) (*entity.EntityEdgePayload, error)
	Delete(id string) error
}
