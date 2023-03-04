package hypolashlckvarresolve

import (
	"github.com/hypolas/hypolaslogger"
)

func makeLogger(logPath string) hypolaslogger.HypolasLogger {
	l := hypolaslogger.NewLogger(logPath)
	return l
}

var (
	logPath = ""
	logf    = makeLogger(logPath)
)
