package hypolashlckvarresolve

import (
	"github.com/hypolas/hypolaslogger"
)

func makeLogger(logPath string) hypolaslogger.HypolasLogger {
	l := hypolaslogger.NewLogger(logPath)
	return l
}

var (
	testPath = "test/log.txt"
	logf     = makeLogger(testPath)
)
