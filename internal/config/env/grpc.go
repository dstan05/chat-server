package env

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net"
)

type GrpcConfig struct {
	host string
	port string
}

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

func NewGRPCConfig(v *viper.Viper) (*GrpcConfig, error) {
	c := &GrpcConfig{}
	c.host = v.GetString(grpcHostEnvName)

	if len(c.host) == 0 {
		return c, errors.New("grpc host not found")
	}

	c.port = v.GetString(grpcPortEnvName)
	if len(c.port) == 0 {
		return c, errors.New("grpc port not found")
	}

	return c, nil
}

func (cfg *GrpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
