package endpoint

import (
	"context"
	"fmt"
	"go-kit-example/app/model/base"
	"go-kit-example/app/model/request"
	"go-kit-example/app/service"

	"github.com/go-kit/kit/endpoint"
)

type ProductEndpoint struct {
	Save   endpoint.Endpoint
	Show   endpoint.Endpoint
	List   endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
}

func MakeProductEndpoints(s service.ProductService) ProductEndpoint {
	return ProductEndpoint{
		Save:   makeSaveProduct(s),
		Show:   makeShowProduct(s),
		List:   makeGetProducts(s),
		Update: makeUpdateProduct(s),
		Delete: makeDeleteProduct(s),
	}
}

func makeSaveProduct(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(request.SaveProductRequest)
		result, msg := s.CreateProduct(req)
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil, nil), nil
		}

		return base.SetHttpResponse(msg.Code, msg.Message, result, nil), nil
	}
}

func makeShowProduct(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		result, msg := s.GetProduct(fmt.Sprint(rqst))
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil, nil), nil
		}

		return base.SetHttpResponse(msg.Code, msg.Message, result, nil), nil
	}
}

func makeGetProducts(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(request.ProductListRequest)
		result, pagination, msg := s.GetList(req)
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil, nil), nil
		}

		return base.SetHttpResponse(msg.Code, msg.Message, result, pagination), nil
	}
}

func makeUpdateProduct(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		req := rqst.(request.SaveProductRequest)
		msg := s.UpdateProduct(req.Uid, req)

		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil, nil), nil
		}

		return base.SetHttpResponse(msg.Code, msg.Message, nil, nil), nil
	}
}

func makeDeleteProduct(s service.ProductService) endpoint.Endpoint {
	return func(ctx context.Context, rqst interface{}) (resp interface{}, err error) {
		msg := s.DeleteProduct(fmt.Sprint(rqst))
		if msg.Code == 4000 {
			return base.SetHttpResponse(msg.Code, msg.Message, nil, nil), nil
		}

		return base.SetHttpResponse(msg.Code, msg.Message, nil, nil), nil
	}
}
