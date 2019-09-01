package logs

// commandLogs describes the logs for a specific command in a task.
type commandLogs struct {
	ID       string `bson:"id"`
	Command  string `bson:"command"`
	ExitCode int    `bson:"exit_code"`
	Logs     string `bson:"logs"`
}

func newCommandLogs(execID string, cmd string) *commandLogs {

	return &commandLogs{
		ID:       execID,
		Command:  cmd,
		ExitCode: -1,
	}

}
