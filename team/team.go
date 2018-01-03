package team

import (
	"dix975.com/score/validation"
	"github.com/dix975/logger"
	"github.com/dix975/www"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
)

var teamSchema gojsonschema.Schema


type Team struct {
	Name string `json:"name"`
}

func PostTeam(response http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()

	var team *Team
	err := validation.FromPost(request, "team.json", &team)

	if err != nil {
		www.Render(response, request, 400, err)
		return
	}

	logger.Info.Println(team)

	www.RenderOK(response, request, team)
}
