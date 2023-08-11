package log

import (
	"fmt"
	"log/slog"
	"os"
)

func Log() *slog.Logger {
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}

	if debug {
		opts.Level = slog.LevelDebug
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}

func MsgAttr(msg string, args ...interface{}) string {
	return fmt.Sprintf(msg, args...)
}

func ErrorAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}
