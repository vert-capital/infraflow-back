package usecase_node

import "app/entity"

type IRepositoryNode interface {
	GetAll(application_id *string) ([]entity.EntityNodePayload, error)
	GetByID(id string) (*entity.EntityNodePayload, error)
	Create(node *entity.EntityNode) (*entity.EntityNodePayload, error)
	Update(node *entity.EntityNode) (*entity.EntityNodePayload, error)
	Delete(id string) error
}

type IUsecaseNode interface {
	GetAll(application_id *string) ([]entity.EntityNodePayload, error)
	GetByID(id string) (*entity.EntityNodePayload, error)
	Create(node *entity.EntityNode) (*entity.EntityNodePayload, error)
	Update(node *entity.EntityNode) (*entity.EntityNodePayload, error)
	Delete(id string) error
}
