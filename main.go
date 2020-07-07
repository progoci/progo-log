package main

import (
	"fmt"
	"net"
	"path/filepath"

	"github.com/progoci/progo-core/config"
	"github.com/progoci/progo-log/internal/app"
	"github.com/progoci/progo-log/pkg/store"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/progoci/progo-log/progolog"
)

func main() {
	// Logger.
	logger := logrus.New()

	// Config.
	envPath, _ := filepath.Abs("./.env")
	config, err := config.New(envPath)
	if err != nil {
		logger.Fatalf("could not get configuration file: %v", err)
	}

	// Store.
	store, err := getStorer(config)
	if err != nil {
		logger.Fatalf("could not create storer: %v", err)
	}

	_ = &app.App{
		Config: config,
		Store:  store,
		Log:    logger,
	}

	port := fmt.Sprintf(":%s", config.Get("HOST_PORT"))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServer(grpcServer, NewProgologServer(store, logger))
	grpcServer.Serve(lis)
}

// getStorer returns an initialized object of store.Store.
func getStorer(config *config.Config) (*store.Store, error) {
	opts := &store.Opts{
		Host:     config.Get("DB_HOST"),
		Port:     config.Get("DB_PORT"),
		Database: config.Get("DB_NAME"),
		URI:      config.Get("DB_URI"),
	}

	return store.New(opts)
}
