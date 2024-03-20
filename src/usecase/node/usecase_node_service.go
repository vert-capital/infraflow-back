package usecase_node

import "app/entity"

type UseCaseNode struct {
	repo IRepositoryNode
}

func NewUseCaseNode(repo IRepositoryNode) *UseCaseNode {
	return &UseCaseNode{repo: repo}
}

func (n *UseCaseNode) GetAll() ([]entity.EntityNodePayload, error) {
	return n.repo.GetAll()
}

func (n *UseCaseNode) GetByID(id string) (*entity.EntityNodePayload, error) {
	return n.repo.GetByID(id)
}

func (n *UseCaseNode) Create(node *entity.EntityNode) (*entity.EntityNodePayload, error) {
	return n.repo.Create(node)
}

func (n *UseCaseNode) Update(node *entity.EntityNode) (*entity.EntityNodePayload, error) {
	return n.repo.Update(node)
}

func (n *UseCaseNode) Delete(id string) error {
	return n.repo.Delete(id)
}
