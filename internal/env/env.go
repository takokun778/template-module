package env

import (
	"os"
	"sync"

	"github.com/takokun778/template-module/pkg/log"
)

type Env string

const (
	prd   Env = "prd"
	stg   Env = "stg"
	dev   Env = "dev"
	local Env = "local"
	empty Env = ""
)

var (
	env  Env        //nolint:gochecknoglobals
	lock sync.Mutex //nolint:gochecknoglobals
)

func Init() {
	lock.Lock()
	defer lock.Unlock()

	e := Env(os.Getenv("ENV"))

	log.Log().Info(log.MsgAttr("environment: %s", e))

	env = e
}

func Get() Env {
	return env
}

func (e Env) String() string {
	return string(e)
}

func (e Env) IsProduction() bool {
	return e == prd
}

func (e Env) IsStaging() bool {
	return e == stg
}

func (e Env) IsDevelopment() bool {
	return e == dev
}

func (e Env) IsLocal() bool {
	return e == local
}
