package main

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/taeho-io/auth/server"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		addr := ":80"
		log := logrus.WithFields(logrus.Fields{
			"addr":        addr,
			"server_type": "grpc",
		})

		cfg, err := server.NewConfig(server.NewSettings())
		if err != nil {
			log.Error(err)
			return
		}

		log.Info("Starting Auth gRPC server")
		if err := server.ServeGRPC(addr, cfg); err != nil {
			log.Error(err)
			return
		}
	}()

	go func() {
		defer wg.Done()

		addr := ":81"
		log := logrus.WithFields(logrus.Fields{
			"addr":        addr,
			"server_type": "grpc",
		})

		cfg, err := server.NewConfig(server.NewSettings())
		if err != nil {
			log.Error(err)
			return
		}

		log.Info("Starting Auth gRPC server")
		if err := server.ServeGRPC(addr, cfg); err != nil {
			log.Error(err)
			return
		}
	}()

	go func() {
		defer wg.Done()

		addr := ":82"
		log := logrus.WithFields(logrus.Fields{
			"addr":        addr,
			"server_type": "http",
		})

		cfg, err := server.NewConfig(server.NewSettings())
		if err != nil {
			log.Error(err)
			return
		}

		log.Info("Starting Auth HTTP server")
		if err := server.ServeHTTP(addr, cfg); err != nil {
			log.Error(err)
			return
		}
	}()

	wg.Wait()
	os.Exit(1)

}
