package repository

import (
	"challenge-08/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (r *ProductRepositoryMock) FindByID(id uint) *models.Product {
	arguments := r.Mock.Called(id)

	fmt.Println(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(models.Product)
	return &product
}
