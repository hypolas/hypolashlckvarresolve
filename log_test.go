package hypolashlckvarresolve

import (
	"testing"
)

var (
	stringList = []string{
		"http://#CMDSTART# ls -la #CMDEND#:8082/ping",
		"test #CMDSTART# hostname -i #CMDEND#/blabla",
	}
)

// Test if log is wheel write.
func TestResolve(t *testing.T) {
	for _, str := range stringList {
		resolveVariable(str)
	}
}
