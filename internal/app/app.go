package app

import (
	"github.com/progoci/progo-core/config"
	log "github.com/sirupsen/logrus"

	"github.com/progoci/progo-log/pkg/store"
)

// App contains dependencies used across the application.
type App struct {
	Config *config.Config
	Store  *store.Store
	Log    *log.Logger
}
