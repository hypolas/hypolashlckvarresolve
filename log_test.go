package hypolashlckvarresolve

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var (
	stringList = []string{
		"http://#CMDSTART# hostname #CMDEND#:8082/ping",
		"test #CMDSTART# hostname -i #CMDEND#/blabla",
	}
)

// TestResolve test different strings format
func TestResolve(t *testing.T) {
	os.Setenv("HYPOLAS_LOGS_FILE", "test/logs.log")
	logf = makeLogger("")
	for _, str := range stringList {
		ResolveVariable(str)
	}

	readFile, err := os.Open(logf.LogFile.Name())

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
