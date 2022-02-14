package initialization

import (
	"go-kit-example/app/api/transport"
	"go-kit-example/app/model/entity"
	"go-kit-example/app/registry"
	"go-kit-example/helper/database"
	"net/http"

	"github.com/go-kit/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func DbInit() (*gorm.DB, error) {
	// Init DB Connection
	db, err := database.NewConnectionDB(viper.GetString("database.driver"), viper.GetString("database.dbname"),
		viper.GetString("database.host"), viper.GetString("database.username"), viper.GetString("database.password"),
		viper.GetInt("database.port"))
	if err != nil {
		return nil, err
	}

	//Define auto migration here
	_ = db.AutoMigrate(&entity.Product{})

	return db, nil
}

func InitRouting(db *gorm.DB, logger log.Logger) *http.ServeMux {
	// Service registry
	prodSvc := registry.RegisterProductService(db, logger)

	// Transport initialization
	swagHttp := transport.SwaggerHttpHandler(log.With(logger, "SwaggerTransportLayer", "HTTP")) //don't delete or change this !!
	prodHttp := transport.ProductHttpHandler(prodSvc, log.With(logger, "ProductTransportLayer", "HTTP"))

	// Routing path
	mux := http.NewServeMux()
	mux.Handle("/", swagHttp) //don't delete or change this!!
	mux.Handle("/product/", prodHttp)

	return mux
}
