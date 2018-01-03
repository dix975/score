package configuration

import (
	"dix975.com/score/test"
	"fmt"
	"os"
	"testing"
)


func TestConfig(t *testing.T) {

	err := os.Chdir("..")
	if err != nil {
		panic(err)
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("%v/%v", dir, "configuration/configuration-test.json")

	os.Setenv("CONFIG_FILE", path)

	folder := Config().SchemaFolder
	test.AssertEqual(t, folder, "test/fixtures/schemas", "")

	os.Setenv("CONFIG_FILE", "")
	Reset()

	folder = Config().SchemaFolder
	test.AssertEqual(t, folder, "schemas", "")

}
