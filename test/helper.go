package test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)


var configFilePath *string

func Setup() {

	if configFilePath == nil {

		err := os.Chdir("..")
		if err != nil {
			panic(err)
		}
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		path := fmt.Sprintf("%v/%v", dir, "configuration/configuration-test.json")
		configFilePath = &path
		os.Setenv("CONFIG_FILE", *configFilePath)
	}

}

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {

	if a == b {
		return
	}

	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}

	t.Error(message)
}

func AssertHasSuffix(t *testing.T, s string, suffix string) {

	if strings.HasSuffix(s, suffix) {
		return
	}

	t.Error(fmt.Sprintf("String [%v] has not suffix [%v]", s, suffix))
}

type NopCloser struct {
	io.Reader
}

func (NopCloser) Close() error { return nil }


type BadReader struct {
	Err error
}

func (r *BadReader) Read(p []byte) (n int, err error){
	return 0, r.Err
}
