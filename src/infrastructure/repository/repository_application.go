package repository

import (
	"app/entity"

	"gorm.io/gorm"
)

type RepositoryApplication struct {
	db *gorm.DB
}

func NewRepositoryApplication(db *gorm.DB) *RepositoryApplication {
	return &RepositoryApplication{
		db: db,
	}
}

func (a *RepositoryApplication) GetAll() ([]entity.EntityApplication, error) {
	var applications []entity.EntityApplication
	err := a.db.Find(&applications).Error
	return applications, err
}

func (a *RepositoryApplication) GetByID(id string) (*entity.EntityApplication, error) {
	var application entity.EntityApplication
	err := a.db.Where("id = ?", id).Find(&application).Error
	return &application, err
}

func (a *RepositoryApplication) Create(entity entity.EntityApplication) (*entity.EntityApplication, error) {
	entity.New()
	err := a.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}
	return a.GetByID(entity.ID)
}

func (a *RepositoryApplication) Update(entity entity.EntityApplication) (*entity.EntityApplication, error) {
	err := a.db.Save(entity).Error
	return &entity, err
}

func (a *RepositoryApplication) Delete(id string) error {
	entity, err := a.GetByID(id)
	if err != nil {
		return err
	}
	err = a.db.Delete(entity).Error
	return err
}
