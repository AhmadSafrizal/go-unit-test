package service

import (
	"golang-unit-test/repository"
	"golang-unit-test/entity"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("CATEGORY NOT FOUND")
	} else {
		return category, nil
	}
}
