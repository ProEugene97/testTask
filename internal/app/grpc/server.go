package grpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"testTask/internal/pkg/api"
	"testTask/internal/pkg/database"
	"testTask/internal/pkg/proto"
	"time"
)

type GRPCServer struct {
	addr    string
	db      database.IDatabase
	logger  *zap.Logger
}

func NewGRPCServer(addr string, db database.IDatabase, logger *zap.Logger) *GRPCServer {
	return &GRPCServer{
		addr,
		db,
		logger,
	}
}


func (gs *GRPCServer) Start() error {
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(gs.logger),
		)),
	)
	proto.RegisterSubServiceServer(server, api.NewSubService(gs.db, gs.logger))

	lis, err := net.Listen("tcp", gs.addr)
	if err != nil {
		return err
	}

	gs.logger.Info("Start grpc server",
		zap.String("addr", gs.addr),
		zap.Time("start", time.Now()),
	)

	return server.Serve(lis)
}