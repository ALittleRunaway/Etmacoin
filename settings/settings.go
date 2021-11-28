package settings

import (
	"database/sql"
	"sync"
)

var (
	ServiceIsRunning bool
	ProgramIsRunning bool
	WritingSync      sync.Mutex
	Port             int
	Db               *sql.DB
)
