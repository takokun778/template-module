package log

import (
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
)

type key struct{}

func SetLogCtx(ctx context.Context) context.Context {
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}

	if debug {
		opts.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	return context.WithValue(ctx, key{}, logger.With("tid", uuid.NewString()))
}

func GetLogCtx(ctx context.Context) *slog.Logger {
	v := ctx.Value(key{})

	log, ok := v.(*slog.Logger)

	if !ok {
		return Log()
	}

	return log
}
