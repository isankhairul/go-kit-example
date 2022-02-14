//  Example Go kit:
//   version: 1.0
//   title: example Go Kit Api
//  Schemes: http
//  Host: localhost:5601
//  BasePath: /
//  Produces:
//    - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"go-kit-example/app/api/initialization"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func main() {
	//Load configuration
	viper.SetConfigType("yaml")
	var profile string = "dev"
	if os.Getenv("KD_ENV") != "" {
		profile = "prd"
	}

	var configFileName []string
	configFileName = append(configFileName, "config-")
	configFileName = append(configFileName, profile)
	viper.SetConfigName(strings.Join(configFileName, ""))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err)
	}

	// Logging init
	logfile, err := os.OpenFile(viper.GetString("server.output-file-path"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer logfile.Close()
	var logger log.Logger
	{
		// if viper.GetString("server.log-output") == "file" {
		// 	w := log.NewSyncWriter(logfile)
		// 	logger = log.NewLogfmtLogger(w)
		// } else {
		// 	logger = log.NewLogfmtLogger(os.Stderr)
		// }

		// force log to file
		w := log.NewSyncWriter(logfile)
		logger = log.NewLogfmtLogger(w)

		logger = level.NewFilter(logger, level.AllowAll())
		logger = level.NewInjector(logger, level.InfoValue())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	// Init DB Connection
	db, err := initialization.DbInit()
	if err != nil {
		_ = logger.Log("Err Db connection :", err.Error())
		panic(err.Error())
	}
	_ = logger.Log("message", "Connection Db Success")

	// Routing initialization
	mux := initialization.InitRouting(db, logger)
	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)

	go func() {
		_ = logger.Log("transport", "HTTP", "addr", viper.GetInt("server.port"))
		errs <- http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("server.port")), nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	_ = logger.Log("exit", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
