package settings

import (
	"Blockchain/database"
	"sync"
)

var (
	ServiceIsRunning bool
	ProgramIsRunning bool
	WritingSync      sync.Mutex
	Port             int
)

var Db, _ = database.Connection()
