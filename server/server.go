package server

import (
	"log"
	"net"

	"github.com/soheilhy/cmux"
	"golang.org/x/net/context"
)

func Serve() error {
	cfg, err := NewConfig(NewSettings())
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tcpMux := cmux.New(lis)
	grpcL := tcpMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings(
		"content-type", "application/grpc"))
	httpL := tcpMux.Match(cmux.HTTP1Fast())
	go func() {
		grpcServer, err := NewGRPCServer(cfg)
		if err != nil {
			log.Fatal("Unable to initialize gRPC server")
		}
		if err := grpcServer.Serve(grpcL); err != nil {
			log.Fatal("Unable to start gRPC server")
		}
	}()
	go func() {
		httpServer, err := NewHttpServer(cfg)
		if err != nil {
			log.Fatal("Unable to initialize HTTP server")
		}
		if err := httpServer.Serve(httpL); err != nil {
			log.Fatal("Unable to start HTTP server")
		}
	}()

	return tcpMux.Serve()
}
