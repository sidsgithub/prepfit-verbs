package mainMods

import(
	"github.com/emicklei/go-restful"
	"net/http"
)

func (u *ProgramResource) CreateProgram(request *restful.Request, response *restful.Response) {
	usr := Program{ID: request.PathParameter("program-id")}
	err := request.ReadEntity(&usr)
	if err == nil {
		u.Programs[usr.ID] = usr
		response.WriteHeaderAndEntity(http.StatusCreated, usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
	Database(usr,"createProgram")
}