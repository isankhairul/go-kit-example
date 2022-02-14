package consul

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/log"
	"github.com/hashicorp/consul/api"
	_ "github.com/lib/pq"
	_ "github.com/spf13/viper/remote"
)

func ConsulRegisterService(serviceName string, port int, logger log.Logger) *consul.Registrar {
	// Register cd Specify the information of an instance.
	host, _ := os.Hostname()
	asr := api.AgentServiceRegistration{
		// Every service instance must have an unique ID.
		ID:   fmt.Sprintf("%v", host),
		Name: serviceName,
		// These two values are the location of an instance.
		Address: host,
		Port:    port,
	}
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "consul:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		_ = logger.Log("err", err)
		os.Exit(1)
	}
	sdClient := consul.NewClient(consulClient)
	return consul.NewRegistrar(sdClient, &asr, logger)
}
