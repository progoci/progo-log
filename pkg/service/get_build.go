package service

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// StartGetBuild starts the loom WebSocket for streaming logs.
func (l *loomService) StartGetBuild(ctx context.Context) http.HandlerFunc {

	upgrader := newGetBuildUpgrader()
	reader := newGetBuildReader()
	//logsDB := database.StartConnection(config.Get("DB_NAME"))

	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}

		log.Println("Client connected...")

		reader(ws)
	}
}

// newGetBuildUpgrader returns a new websocket upgrader.
func newGetBuildUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

// newGetBuildReader returns a new websocket reader.
func newGetBuildReader() WebSocketReader {
	return func(conn *websocket.Conn) {

		defer func() {
			log.Println("Closing connection")
			conn.Close()
		}()

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)

				break
			}

			log.Println(string(p))

			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)

				break
			}
		}
	}
}
