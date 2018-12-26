package server

import (
	"log"
	"net"
	"sync"
)

func Serve() error {
	cfg, err := NewConfig(NewSettings())
	if err != nil {
		return err
	}

	grpcLis, err := net.Listen("tcp", ":80")
	if err != nil {
		return err
	}
	grpcServer, err := NewGRPCServer(cfg)
	if err != nil {
		return err
	}

	httpLis, err := net.Listen("tcp", ":81")
	if err != nil {
		return err
	}
	httpServer, err := NewHttpServer(cfg)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1) // 1 is correct since we may want to exit when either grpc or http server is gone.

	go func() {
		defer wg.Done()

		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatal("Unable to start gRPC server")
		}
	}()
	go func() {
		defer wg.Done()

		if err := httpServer.Serve(httpLis); err != nil {
			log.Fatal("Unable to start HTTP server")
		}
	}()

	wg.Wait()

	return nil
}
