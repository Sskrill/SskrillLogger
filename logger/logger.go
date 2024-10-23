package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type customLogger struct {
	logger *slog.Logger
}

func NewCSLoggerSTDOUT(opts *slog.HandlerOptions) customLogger {
	handler := slog.NewTextHandler(os.Stdout, opts)
	return customLogger{logger: slog.New(handler)}
}
func (cL customLogger) PrintInfo(info string, infoArgs ...interface{}) {
	cL.logger.Info(fmt.Sprintf(info, infoArgs))
}
func (cL customLogger) PrintError(error string, errorArgs ...interface{}) {
	cL.logger.Error(fmt.Sprintf(error, errorArgs))
}
func (cL customLogger) PrintDebug(debug string, debugArgs ...interface{}) {
	cL.logger.Debug(fmt.Sprintf(debug, debugArgs))
}
