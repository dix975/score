package validation

import (
	"github.com/xeipuuv/gojsonschema"
	"fmt"
)

type ValidationIssues struct {
	Field       string
	Type        string
	Description string
	Detail      gojsonschema.ErrorDetails
	Value       interface{}
}

type ValidationError struct {
	Issues []ValidationIssues
}

func (e *ValidationError) Error() string {

	var message = ""

	for _, err := range e.Issues {
		message += fmt.Sprintf("- %s\n", err)
	}

	return message
}

func NewValidationError(resultErrors []gojsonschema.ResultError) *ValidationError {

	validationIssues := []ValidationIssues{}
	for _, err := range resultErrors {
		validationIssues = append(validationIssues,
			ValidationIssues{
				Field:       err.Field(),
				Type:        err.Type(),
				Description: err.Description(),
				Value:       err.Value(),
				Detail:      err.Details(),
			})
	}

	return &ValidationError{
		Issues: validationIssues,
	}

}


