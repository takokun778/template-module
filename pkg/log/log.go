package log

import (
	"fmt"
	"log/slog"
	"os"
)

func Log() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func ErrorAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func MsgAttr(msg string, args ...interface{}) string {
	return fmt.Sprintf(msg, args...)
}
