package usecase_application

import "app/entity"

type UseCaseApplication struct {
	repo IRepositoryApplication
}

func NewUseCaseApplication(repo IRepositoryApplication) *UseCaseApplication {
	return &UseCaseApplication{
		repo: repo,
	}
}

func (a *UseCaseApplication) GetAll() ([]entity.EntityApplication, error) {
	return a.repo.GetAll()
}

func (a *UseCaseApplication) GetByID(id string) (*entity.EntityApplication, error) {
	return a.repo.GetByID(id)
}

func (a *UseCaseApplication) Create(entity entity.EntityApplication) (*entity.EntityApplication, error) {
	return a.repo.Create(entity)
}

func (a *UseCaseApplication) Update(entity entity.EntityApplication) (*entity.EntityApplication, error) {
	return a.repo.Update(entity)
}

func (a *UseCaseApplication) Delete(id string) error {
	return a.repo.Delete(id)
}
