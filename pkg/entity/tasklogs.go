package logs

// TaskLogs describes the logs for a specific step in a build.
type taskLogs struct {
	UUID     string        `bson:"uuid"`
	ExitCode int           `bson:"exit_code"`
	Cmds     []commandLogs `bson:"commands"`
}

// newTaskLogs creates a base task logs document.
func newTaskLogs(taskUUID string) *taskLogs {

	return &taskLogs{
		UUID:     taskUUID,
		ExitCode: -1,
		Cmds:     []commandLogs{},
	}

}
