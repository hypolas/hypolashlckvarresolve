package hypolashlckvarresolve

import (
	"github.com/hypolas/hypolaslogger"
	"os"
)

func makeLogger(logPath string) hypolaslogger.HypolasLogger {
	l := hypolaslogger.NewLogger(logPath)
	return l
}

var (
	logf = makeLogger(os.Getenv("HYPOLAS_LOGS_FOLDER"))
)
