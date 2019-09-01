package logs

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"progo/build/pkg/database"
	"progo/core/log"
)

// Logs describes a document for the logs of running the steps of a build.
type Logs struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	BuildID    string             `bson:"build_id"`
	Tasks      []taskLogs         `bson:"tasks"`
	lastTask   int                `bson:"last_task"`
	lastCmd    int                `bson:"last_cmd"`
	collection database.Collection
	ctx        context.Context
}

// NewLogs create a base logs document
func NewLogs(ctx context.Context, db database.Database, buildID string) *Logs {

	return &Logs{
		BuildID:    buildID,
		Tasks:      []taskLogs{},
		collection: db.Collection("logs"),
		ctx:        ctx,
	}

}

// FindLogs gets a document describing logs for a task.
func FindLogs(ctx context.Context, conn database.Database, buildID string,
	taskUUID string) *Logs {

	filter := map[string]string{
		"build_id": buildID,
	}

	logs := Logs{}

	err := conn.Collection("logs").FindOne(ctx, filter, &options.FindOneOptions{}).Decode(&logs)

	if err != nil {
		log.Print("debug", "Could not decode result for logs", err)

		return nil
	}

	return &logs
}

// Save stores the logs document into the database.
func (l Logs) Save() error {

	_, err := l.collection.InsertOne(l.ctx, l, &options.InsertOneOptions{})

	if err != nil {
		log.Print("error", "Error inserting logs", err)

		return err
	}

	return nil
}

/*
[
   {
     $project: {
       lastTask: { $arrayElemAt: [ { $slice: [ "$tasks", -1 ] }, 0 ] },
     }
   },
   {
     $project: {
       lastCmd: { $arrayElemAt: [ { $slice: [ "$lastTask.commands", -1 ] }, 0 ] },
     }
   },
   {
     $set: {
       "lastCmd.logs" : { $concat: ["$lastCmd.logs", "hello"] }
     }
   }
]
*/
// AppendCommandOuput concatenates the buffer.
func (l Logs) AppendCommandOutput(buildID string, taskUUID string, execID string, buf []byte) error {

	filter := map[string]string{
		"build_id": l.BuildID,
	}

	/*pipeline := bson.A{
		bson.D{{
			"$set", bson.D{{
				"tasks.$[t].$[s].logs", bson.D{{
					"$concat", bson.A{"$lastCmd.logs", "hey"},
				}},
			}},
		}},
	}*/

	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"tasks.$[t].commands.$[s].logs": "$tasks.$[t].commands.$[s].logs" + "hello",
		},
	}
	options := &options.UpdateOptions{
		ArrayFilters: &options.ArrayFilters{
			Filters: []interface{}{
				map[string]string{
					"t.uuid": taskUUID,
				},
				map[string]string{
					"s.id": execID,
				},
			},
		},
	}

	result, err := l.collection.UpdateOne(l.ctx, filter, update, options)

	if err != nil {
		log.Print("error", "Error appending command output", err)
		panic(err)
		//return err
	}
	if result.MatchedCount == 0 {
		err = errors.New("Filter did not match any documents")

		log.Print("error", "Error appending command output", err)

		return err
	}

	return nil

}

// AddTask appends a task to the logs.
func (l Logs) AddTask(taskUUID string) error {

	filter := map[string]string{
		"build_id": l.BuildID,
	}

	update := map[string]interface{}{
		"$push": map[string]*taskLogs{
			"tasks": newTaskLogs(taskUUID),
		},
		"$inc": map[string]interface{}{
			"last_task": 1,
		},
	}

	result, err := l.collection.UpdateOne(l.ctx, filter, update)

	if err != nil {
		log.Print("error", "Error inserting new task logs", err)

		return err
	}
	if result.MatchedCount == 0 {
		err = errors.New("Filter did not match any documents")

		log.Print("error", "Error inserting new task logs", err)

		return err
	}

	return nil
}

// AddCommand appends a command logs at the last task.
func (l Logs) AddCommand(taskUUID string, execID string, cmd string) error {
	filter := map[string]string{
		"build_id": l.BuildID,
	}

	// We need to append the new command at then end of an array which is inside
	// another array (the task array).
	// See https://docs.mongodb.com/manual/reference/operator/update/positional-filtered/#position-nested-arrays-filtered
	update := map[string]interface{}{
		"$push": map[string]*commandLogs{
			"tasks.$[t].commands": newCommandLogs(execID, cmd),
		},
		"$inc": map[string]interface{}{
			"last_cmd": 1,
		},
	}
	options := &options.UpdateOptions{
		ArrayFilters: &options.ArrayFilters{
			Filters: []interface{}{
				map[string]string{
					"t.uuid": taskUUID,
				},
			},
		},
	}

	result, err := l.collection.UpdateOne(l.ctx, filter, update, options)

	if err != nil {
		log.Print("error", "Error inserting new task logs", err)

		return err
	}
	if result.MatchedCount == 0 {
		err = errors.New("Filter did not match any documents")

		log.Print("error", "Error inserting new task logs", err)

		return err
	}

	return nil
}
