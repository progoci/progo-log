/**
 * gRPC server for logs.
 */

package main

import (
	"io"
	"log"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	"github.com/progoci/progo-log/pkg/store"
	pb "github.com/progoci/progo-log/progolog"
)

// ProgologServer is the gRPC server for logs.
type ProgologServer struct {
	Storer *store.Store
	Logger *logrus.Logger
}

// NewProgologServer initializes a progolog server.
func NewProgologServer(storer *store.Store, logger *logrus.Logger) *ProgologServer {
	return &ProgologServer{storer, logger}
}

// Store saves logs from a single step.
//
// A stream of logs is generated for each step, which is just an attachment to
// a docker exec instance.
func (l *ProgologServer) Store(stream pb.Logger_StoreServer) error {
	var logs []byte
	var buildID, serviceName, stepName, command string
	var stepNumber int32

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			err := l.Storer.Save(logs, buildID, serviceName, stepName, command, stepNumber)
			if err != nil {
				l.Logger.Errorf("failed to save logs: %v", err)
				return errors.Wrap(err, "failed to save logs")
			}

			return stream.SendAndClose(&pb.Response{})
		}
		if err != nil {
			return errors.Wrap(err, "failed to receive log")
		}

		// Fist packet of the logs.
		if stepNumber == 0 {
			log.Println("here")
			buildID = msg.BuildID
			serviceName = msg.ServiceName
			stepName = msg.StepName
			stepNumber = msg.StepNumber
			command = msg.Command
		}
		logs = append(logs, msg.Body...)
	}
}
