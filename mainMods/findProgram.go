package mainMods

import (
	"github.com/emicklei/go-restful"
	"net/http"
)

func (u ProgramResource) FindProgram(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("program-id")
	usr := u.Programs[id]
	if len(usr.ID) == 0 {
		response.WriteErrorString(http.StatusNotFound, "program could not be found.")
	} else {
		response.WriteEntity(usr)
	}
	Database(usr,"findProgram")
}