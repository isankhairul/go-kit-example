package registry

import (
	rp "go-kit-example/app/repository"
	"go-kit-example/app/service"

	"github.com/go-kit/log"
	"gorm.io/gorm"
)

func RegisterProductService(db *gorm.DB, logger log.Logger) service.ProductService {
	return service.NewProductService(
		logger,
		rp.NewBaseRepository(db),
		rp.NewProductRepository(rp.NewBaseRepository(db)),
	)
}
