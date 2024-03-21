package usecase_application

import "app/entity"

type IRepositoryApplication interface {
	GetAll() ([]entity.EntityApplication, error)
	GetByID(id string) (*entity.EntityApplication, error)
	Create(entity entity.EntityApplication) (*entity.EntityApplication, error)
	Update(entity entity.EntityApplication) (*entity.EntityApplication, error)
	Delete(id string) error
}

type IUsecaseApplication interface {
	GetAll() ([]entity.EntityApplication, error)
	GetByID(id string) (*entity.EntityApplication, error)
	Create(entity entity.EntityApplication) (*entity.EntityApplication, error)
	Update(entity entity.EntityApplication) (*entity.EntityApplication, error)
	Delete(id string) error
}
