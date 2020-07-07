package store

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/progoci/progo-log/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store is the log storage manager.
type Store struct {
	database Database
}

// Database describes a connection to a MongoDB database.
type Database interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

// Opts is the information to connnect to the database.
// If URI is set, it overwrites all other connection configuration.
type Opts struct {
	Host     string
	Port     string
	Database string
	URI      string
}

// New initializes a Store and a creates connection to a MongoDB database.
func New(opts *Opts) (*Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := opts.URI
	if uri == "" {
		uri = "mongodb://" + opts.Host + ":" + opts.Port
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	return &Store{
		database: client.Database(opts.Database),
	}, nil
}

// Save stores a new log.
func (s *Store) Save(body []byte, buildID, serviceName, stepName, command string, stepNumber int32) error {
	lg := types.StepLog{
		BuildID:     buildID,
		ServiceName: serviceName,
		StepName:    stepName,
		StepNumber:  stepNumber,
		Command:     command,
		Body:        string(body),
	}

	_, err := s.database.Collection("logs").InsertOne(context.Background(), lg)
	if err != nil {
		return errors.Wrap(err, "failed to store logs in database")
	}

	return nil
}
