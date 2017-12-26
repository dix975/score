package team

import (
	"net/http"
	"encoding/json"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"fmt"
	"os"
	"github.com/dix975/score/validation"
	"github.com/dix975/www"
	"github.com/dix975/logger"
)

var teamSchema gojsonschema.Schema

func init() {

	dir, err := os.Getwd()

	teamSchemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file:///%v/schemas/team.json", dir))
	schema, err := gojsonschema.NewSchema(teamSchemaLoader)
	if err != nil {
		panic(err)
	}
	teamSchema = *schema
}

type Team struct {
	Name string `json:"name"`
}

func from(request *http.Request) (team Team, err error) {

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	validationResult, err := teamSchema.Validate(
		gojsonschema.NewStringLoader(string(body)))

	if err != nil {
		panic(err.Error())
	}

	if !validationResult.Valid() {
		err = validation.NewValidationError(validationResult.Errors())
		return
	}

	err = json.Unmarshal(body, &team)
	if err != nil {
		return
	}

	return
}

func PostTeam(response http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()

	team, err := from(request)
	if err != nil {
		www.Render(response, request, 400, err)
		return
	}

	logger.Info.Println(team)

	www.RenderOK(response, request, "")
}
