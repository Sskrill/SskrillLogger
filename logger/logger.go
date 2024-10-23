package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type CustomLogger struct {
	logger *slog.Logger
}

func NewCSLoggerSTDOUT(opts *slog.HandlerOptions) CustomLogger {
	handler := slog.NewTextHandler(os.Stdout, opts)
	return CustomLogger{logger: slog.New(handler)}
}
func (cL CustomLogger) PrintInfo(info string, infoArgs ...interface{}) {
	cL.logger.Info(fmt.Sprintf(info, infoArgs))
}
func (cL CustomLogger) PrintError(error string, errorArgs ...interface{}) {
	cL.logger.Error(fmt.Sprintf(error, errorArgs))
}
func (cL CustomLogger) PrintDebug(debug string, debugArgs ...interface{}) {
	cL.logger.Debug(fmt.Sprintf(debug, debugArgs))
}
