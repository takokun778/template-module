package handler

import (
	"net/http"

	"github.com/takokun778/template-module/internal/cache"
	"github.com/takokun778/template-module/internal/database"
	"github.com/takokun778/template-module/pkg/openapi"
)

var _ openapi.ServerInterface = (*Handler)(nil)

type Handler struct {
	cache    *cache.Cache
	database *database.Database
}

func New(
	cache *cache.Cache,
	database *database.Database,
) *Handler {
	return &Handler{
		cache:    cache,
		database: database,
	}
}

func (hdl *Handler) V1HealthAPI(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Write([]byte("OK"))
}

func (hdl *Handler) V1HealthCache(
	w http.ResponseWriter,
	r *http.Request,
) {
	cmd := hdl.cache.Client.B().Ping().Build()

	if err := hdl.cache.Client.Do(r.Context(), cmd).Error(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Write([]byte("OK"))
}

func (hdl *Handler) V1HealthDB(
	w http.ResponseWriter,
	r *http.Request,
) {
	if err := hdl.database.Client.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Write([]byte("OK"))
}

func (hdl *Handler) V1Hello(
	w http.ResponseWriter,
	r *http.Request,
) {
}
