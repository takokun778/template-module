package log

import "sync"

var debug bool      //nolint:gochecknoglobals
var lock sync.Mutex //nolint:gochecknoglobals

func SetDebug() {
	lock.Lock()
	defer lock.Unlock()

	debug = true
}
