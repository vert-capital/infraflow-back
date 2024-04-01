package usecase_edge

import "app/entity"

type UseCaseEdge struct {
	repo IRepositoryEdge
}

func NewUseCaseEdge(repo IRepositoryEdge) *UseCaseEdge {
	return &UseCaseEdge{repo: repo}
}

func (e *UseCaseEdge) GetAll(node_id *string) ([]entity.EntityEdgePayload, error) {
	return e.repo.GetAll(node_id)
}

func (e *UseCaseEdge) GetByID(id string) (*entity.EntityEdgePayload, error) {
	return e.repo.GetByID(id)
}

func (e *UseCaseEdge) Create(edge *entity.EntityEdge) (*entity.EntityEdgePayload, error) {
	return e.repo.Create(edge)
}

func (e *UseCaseEdge) Update(edge *entity.EntityEdge) (*entity.EntityEdgePayload, error) {
	return e.repo.Update(edge)
}

func (e *UseCaseEdge) Delete(id string) error {
	return e.repo.Delete(id)
}
