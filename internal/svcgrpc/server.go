package svcgrpc

import (
	"github.com/sandrolain/permissions/internal/dbsvc"
	g "github.com/sandrolain/permissions/pkg/grpc"
)

func NewGrpcServer(s dbsvc.DBServiceInterface) g.PermissionsServiceServer {
	return &grpcServer{
		db: s,
	}
}

type grpcServer struct {
	db dbsvc.DBServiceInterface
	g.UnimplementedPermissionsServiceServer
}
