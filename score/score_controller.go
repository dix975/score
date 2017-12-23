package score

import (
	"net/http"
	"dix975.com/www"
)

func HandleRoot(request http.ResponseWriter, response *http.Request) {

	www.RenderOK(request, response, "Have fun with score!")

}
