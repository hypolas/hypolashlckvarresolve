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

// TestResolve test different strings format
func TestResolve(t *testing.T) {
	for _, str := range stringList {
		ResolveVariable(str)
	}
}
