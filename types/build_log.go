package types

// BuildLogs represent the message sent to the Loom service to store logs for
// a build.
type BuildLogs struct {
	BuildID  string
	TaskUUID string
	CmdID    string
	Logs     string
}
