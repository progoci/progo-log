package service

import (
	"flag"
	"net/http"

	"progo/core/log"
	"progo/loom/pkg/endpoint"
	servicehttp "progo/loom/pkg/http"
	"progo/loom/pkg/kafka"
	"progo/loom/pkg/service"
)

// Run starts the loom service.
func Run(port string) {
	var (
		httpAddr = flag.String("http.addr", port, "HTTP listen address")
	)
	flag.Parse()

	svc := service.NewLoomService()

	var h http.Handler
	{
		endpoints := endpoint.MakeEndpoints(svc)
		h = servicehttp.NewService(endpoints)
	}

	server := &http.Server{
		Addr:    *httpAddr,
		Handler: h,
	}

	runLogsListener()

	log.Print("debug", "", "Running loom service on port "+port)
	log.Print("debug", "", server.ListenAndServe())
}

// Creates a new Kafka Consumer for logs and reads messages in a goroutine.
func runLogsListener() {
	consumer := kafka.NewLogsConsumer()

	go kafka.Read(consumer)
}
