package main

import (
	"github.com/sandrolain/gomsvc/pkg/dblib"
	"github.com/sandrolain/gomsvc/pkg/grpclib"
	"github.com/sandrolain/gomsvc/pkg/svc"
	"github.com/sandrolain/permissions/internal/dbsvc"
	im "github.com/sandrolain/permissions/internal/models"
	"github.com/sandrolain/permissions/internal/svcgrpc"
	g "github.com/sandrolain/permissions/pkg/grpc"
	"github.com/sandrolain/permissions/pkg/models"
)

func main() {
	svc.Service(svc.ServiceOptions{
		Name:    "permissions",
		Version: "1.0.0",
	}, func(cfg im.Config) {

		svc.Logger().Info("Open Database Connection")
		db := svc.PanicWithError(dblib.GormOpenPostgres(dblib.FromEnvConfig(cfg.Postgres)))

		svc.OnExit(func() {
			svc.Logger().Info("Close Database Connection")
			conn, err := db.DB()
			if err != nil {
				conn.Close()
			}
		})

		svc.PanicIfError(db.AutoMigrate(&models.Permission{}))

		opts := grpclib.ServerOptionsFromEnvConfig(cfg.Grpc)
		opts.ServiceDesc = &g.PermissionsService_ServiceDesc
		opts.Handler = svcgrpc.NewGrpcServer(dbsvc.NewDBService(db))

		s := svc.PanicWithError(grpclib.NewGrpcServer(opts))

		svc.PanicIfError(s.Start())
	})
}
