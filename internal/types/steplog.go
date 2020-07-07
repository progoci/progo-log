package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StepLog is the log from a single step in a service.
type StepLog struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	BuildID     string             `bson:"buildID"`
	ServiceName string             `bson:"serviceName"`
	StepName    string             `bson:"stepName"`
	StepNumber  int32              `bson:"stepNumber"`
	Command     string             `bson:"command"`
	Body        string             `bson:"body"`
}
