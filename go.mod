module progo/loom

go 1.12

require (
	github.com/go-kit/kit v0.9.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/websocket v1.4.0
	github.com/segmentio/kafka-go v0.3.3
	go.mongodb.org/mongo-driver v1.1.0
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586 // indirect
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	golang.org/x/text v0.3.2 // indirect

	progo/core v0.0.0
)

replace progo/core => ../../core
