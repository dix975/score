package score

import (
	"net/http"
	"github.com/dix975/www"
)

func HandleRoot(request http.ResponseWriter, response *http.Request) {

	www.RenderOK(request, response, "Have fun with score!")

}
