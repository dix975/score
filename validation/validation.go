package validation

import (
	"dix975.com/logger"
	"dix975.com/score/configuration"
	"encoding/json"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"net/http"
	"os"
)

type Issues struct {
	Field       string
	Type        string
	Description string
	Value       interface{}
}

type Error struct {
	Issues []Issues
}

func (e *Error) Error() string {

	var message = ""

	for _, err := range e.Issues {
		message += fmt.Sprintf("- %s\n", err.Description)
	}

	return message
}

func NewValidationError(resultErrors []gojsonschema.ResultError) *Error {

	validationIssues := []Issues{}
	for _, err := range resultErrors {
		validationIssues = append(validationIssues,
			Issues{
				Field:       err.Field(),
				Type:        err.Type(),
				Description: err.Description(),
				Value:       err.Value(),
			})
	}

	return &Error{
		Issues: validationIssues,
	}
}

func LoadSchema(name string) (schema gojsonschema.Schema, err error) {

	dir, err := os.Getwd()

	teamSchemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file:///%v/%v/%v", dir, configuration.Config().SchemaFolder, name))

	s, err := gojsonschema.NewSchema(teamSchemaLoader)

	if err != nil {
		logger.Error.Printf("Error : %v", err)
		return
	}

	schema = *s

	return
}

func Validate(jsonString string, schemaName string) (err error){

	// Todo : add cache here!
	schema, err := LoadSchema(schemaName)
	if err != nil {
		return
	}

	validationResult, err := schema.Validate(
		gojsonschema.NewStringLoader(string(jsonString)))

	if err != nil {
		return
	}

	if !validationResult.Valid() {
		err = NewValidationError(validationResult.Errors())
		return
	}

	return
}

func FromPost(request *http.Request, schemaName string, model interface{}) (err error) {

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}


	err = Validate(string(body), schemaName)

	if err != nil {
		return
	}


	err = json.Unmarshal(body, &model)

	return
}
