package models

import (
	"github.com/sandrolain/gomsvc/pkg/dblib"
	"github.com/sandrolain/gomsvc/pkg/grpclib"
)

type Config struct {
	Grpc     grpclib.EnvServerConfig
	Postgres dblib.EnvConfig
}
