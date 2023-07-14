package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/takokun778/template-module/internal/cache"
	"github.com/takokun778/template-module/internal/database"
	"github.com/takokun778/template-module/internal/handler"
	"github.com/takokun778/template-module/pkg/log"
	"github.com/takokun778/template-module/pkg/openapi"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := chi.NewRouter()

	cs := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		Debug:            false,
		MaxAge:           300,
	})

	router.Use(cs)

	logger := httplog.NewLogger("access-log", httplog.Options{
		JSON: true,
	})

	router.Use(httplog.RequestLogger(logger))

	cc, err := cache.New(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	db, err := database.New(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	hdl := openapi.HandlerWithOptions(handler.New(cc, db), openapi.ChiServerOptions{
		BaseURL:    "/api",
		BaseRouter: router,
	})

	srv := NewServer(port, hdl)

	srv.Run()
}

const (
	shutdownTime      = 10
	readHeaderTimeout = 30 * time.Second
)

type Server struct {
	*http.Server
}

func NewServer(
	port string,
	handler http.Handler,
) *Server {
	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return &Server{
		Server: server,
	}
}

func (srv *Server) Run() {
	log.Log().Sugar().Infof("server running on %s", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Log().Sugar().Errorf("server closed with error: %s", err.Error())

			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Log().Sugar().Infof("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTime*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Log().Sugar().Infof("failed to gracefully shutdown:", err)
	}

	log.Log().Info("server shutdown")
}
