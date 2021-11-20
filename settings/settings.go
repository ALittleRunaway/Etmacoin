package settings

import (
	"sync"
)

var (
	ServiceIsRunning bool
	ProgramIsRunning bool
	WritingSync      sync.Mutex
	Port             int
)
