module github.com/progoci/progo-log

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/gddo v0.0.0-20200611223618-a4829ef13274
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/progoci/progo-build v0.0.0-20200704013409-a9c84b4ab9e8
	github.com/progoci/progo-core v0.0.0-20200703210147-b5e9f8fc24ff
	github.com/sirupsen/logrus v1.6.0
	go.mongodb.org/mongo-driver v1.3.4
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)

// For development
// replace github.com/progoci/progo-core => ./core
