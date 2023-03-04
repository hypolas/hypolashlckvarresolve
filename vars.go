package hypolashlckvarresolve

import (
	"github.com/hypolas/hypolaslogger"
)

func makeLogger() hypolaslogger.HypolasLogger {
	l := hypolaslogger.NewLogger("test/log.txt")
	return l
}

var (
	logf = makeLogger()
)
