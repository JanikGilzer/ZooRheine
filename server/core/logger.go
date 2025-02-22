package core

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func Logger_init() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(Logger)
}
