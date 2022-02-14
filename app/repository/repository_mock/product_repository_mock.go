package repository_mock

import (
	"go-kit-example/app/model/base"
	"go-kit-example/app/model/entity"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindByParams(limit int, page int, sort string, filter map[string]interface{}) ([]entity.Product, *base.Pagination, error) {
	arguments := repository.Mock.Called(limit, page, sort, filter)

	return arguments.Get(0).([]entity.Product), arguments.Get(1).(*base.Pagination), nil
}

func (repository *ProductRepositoryMock) FindByUid(uid *string) (*entity.Product, error) {
	arguments := repository.Mock.Called(uid)
	if arguments.Get(0) == nil {
		return nil, nil
	} else {
		product := arguments.Get(0).(entity.Product)
		return &product, nil
	}
}

func (repository *ProductRepositoryMock) Create(product *entity.Product) (*entity.Product, error) {
	return product, nil
}

func (repository *ProductRepositoryMock) Delete(uid string) error {
	return nil
}

func (repository *ProductRepositoryMock) Update(uid string, input map[string]interface{}) error {
	return nil
}

func (repository *ProductRepositoryMock) Paginate(value interface{}, pagination *base.Pagination, db *gorm.DB, currRecord int64) func(db *gorm.DB) *gorm.DB {
	return nil
}
