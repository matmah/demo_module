package log

import (
	"log/slog"
	"os"
)

func InitSlog() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Debug("This is my first version")
}

