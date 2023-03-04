package hypolashlckvarresolve

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var (
	stringList = []string{
		"http://#CMDSTART# ls -la #CMDEND#:8082/ping",
		"test #CMDSTART# hostname -i #CMDEND#/blabla",
	}
)

const (
	testPath = "test/log.txt"
)

// TestResolve test different strings format
func TestResolve(t *testing.T) {
	logf = makeLogger(testPath)
	for _, str := range stringList {
		ResolveVariable(str)
	}

	readFile, err := os.Open(logPath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
