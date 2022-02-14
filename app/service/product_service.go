package service

import (
	"go-kit-example/app/model/base"
	"go-kit-example/app/model/entity"
	"go-kit-example/app/model/request"
	"go-kit-example/app/repository"
	"go-kit-example/helper/message"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type ProductService interface {
	CreateProduct(input request.SaveProductRequest) (*entity.Product, message.Message)
	GetProduct(uid string) (*entity.Product, message.Message)
	GetList(input request.ProductListRequest) ([]entity.Product, *base.Pagination, message.Message)
	UpdateProduct(uid string, input request.SaveProductRequest) message.Message
	DeleteProduct(uid string) message.Message
}

type productServiceImpl struct {
	logger      log.Logger
	baseRepo    repository.BaseRepository
	productRepo repository.ProductRepository
}

func NewProductService(
	lg log.Logger,
	br repository.BaseRepository,
	pr repository.ProductRepository,
) ProductService {
	return &productServiceImpl{lg, br, pr}
}

// swagger:route POST /product/ product SaveProductRequest
// Create product

// security:
// - apiKey: []
// responses:
//  401: SuccessResponse
//  201: SuccessResponse
func (s *productServiceImpl) CreateProduct(input request.SaveProductRequest) (*entity.Product, message.Message) {
	logger := log.With(s.logger, "ProductService", "CreateProduct")
	s.baseRepo.BeginTx()
	//Set request to entity
	product := entity.Product{
		Name:   input.Name,
		Sku:    input.Sku,
		Uom:    input.Uom,
		Weight: input.Weight,
	}

	result, err := s.productRepo.Create(&product)
	if err != nil {
		_ = level.Error(logger).Log(err)
		s.baseRepo.RollbackTx()
		return nil, message.FailedMsg
	}
	s.baseRepo.CommitTx()

	return result, message.SuccessMsg
}

// swagger:route GET /product/ product get_product
// Get product
//
// security:
// - apiKey: []
// responses:
//  401: SuccessResponse
//  200: SuccessResponse
func (s *productServiceImpl) GetProduct(uid string) (*entity.Product, message.Message) {
	logger := log.With(s.logger, "ProductService", "GetProduct")

	result, err := s.productRepo.FindByUid(&uid)
	if err != nil {
		_ = level.Error(logger).Log(err)
		return nil, message.FailedMsg
	}

	if result == nil {
		return nil, message.FailedMsg
	}

	return result, message.SuccessMsg
}

// swagger:route GET /product/list product productList
// Get products
//
// security:
// - apiKey: []
// responses:
//  401: SuccessResponse
//  200: SuccessResponse
func (s *productServiceImpl) GetList(input request.ProductListRequest) ([]entity.Product, *base.Pagination, message.Message) {
	logger := log.With(s.logger, "ProductService", "GetList")

	//Set default value
	if input.Limit <= 0 {
		input.Limit = 10
	}
	if input.Page <= 0 {
		input.Page = 1
	}

	filter := map[string]interface{}{
		"name": input.Name,
		"sku":  input.Sku,
		"uom":  input.UOM,
	}

	result, pagination, err := s.productRepo.FindByParams(input.Limit, input.Page, input.Sort, filter)
	if err != nil {
		_ = level.Error(logger).Log(err)
		return nil, nil, message.FailedMsg
	}

	if result == nil {
		_ = level.Warn(logger).Log(message.ErrNoData)
		return nil, nil, message.FailedMsg
	}

	return result, pagination, message.SuccessMsg
}

// swagger:route PUT /product/{id} product SaveProductRequest
// Update product
//
// security:
// - apiKey: []
// responses:
//  401: SuccessResponse
//  200: SuccessResponse
func (s *productServiceImpl) UpdateProduct(uid string, input request.SaveProductRequest) message.Message {
	logger := log.With(s.logger, "ProductService", "UpdateProduct")

	_, err := s.productRepo.FindByUid(&uid)
	if err != nil {
		_ = level.Error(logger).Log(err)
		return message.FailedMsg
	}

	data := map[string]interface{}{
		"name":   input.Name,
		"sku":    input.Sku,
		"uom":    input.Uom,
		"weight": input.Weight,
	}

	err = s.productRepo.Update(uid, data)
	if err != nil {
		_ = level.Error(logger).Log(err)
		return message.FailedMsg
	}

	return message.SuccessMsg
}

// swagger:route DELETE /product/{id} product delete_product
// Delete product
//
// security:
// - apiKey: []
// responses:
//  401: SuccessResponse
//  200: SuccessResponse
func (s *productServiceImpl) DeleteProduct(uid string) message.Message {
	logger := log.With(s.logger, "ProductService", "DeleteProduct")

	_, err := s.productRepo.FindByUid(&uid)
	if err != nil {
		_ = level.Error(logger).Log(err)
		return message.FailedMsg
	}

	err = s.productRepo.Delete(uid)
	if err != nil {
		_ = level.Error(logger).Log(err)
		return message.FailedMsg
	}

	return message.SuccessMsg
}
