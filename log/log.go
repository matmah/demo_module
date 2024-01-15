package log

import (
	"log/slog"
	"os"
)

func InitSlog() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("This is my second version")
}

