package log

import (
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
)

type key struct{}

func SetLogCtx(ctx context.Context) context.Context {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return context.WithValue(ctx, key{}, logger.With("tie", uuid.NewString()))
}

func GetLogCtx(ctx context.Context) *slog.Logger {
	v := ctx.Value(key{})

	log, ok := v.(*slog.Logger)

	if !ok {
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	return log
}
