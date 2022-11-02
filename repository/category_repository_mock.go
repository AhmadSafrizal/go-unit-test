package repository

import (
	"github.com/stretchr/testify/mock"
	"golang-unit-test/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arg := repository.Mock.Called(id)
	if arg.Get(0) == nil {
		return nil
	} else {
		category := arg.Get(0).(entity.Category)
		return &category
	}
}