package mainMods

import(
	"github.com/emicklei/go-restful"
	"net/http"
	"log"
)
func (u *ProgramResource) CreateProgram(request *restful.Request, response *restful.Response) {
	prog := Program{ID: request.PathParameter("program-id")}
	err := request.ReadEntity(&prog)
	if err == nil {
		u.Programs[prog.ID] = prog
		response.WriteHeaderAndEntity(http.StatusCreated, prog)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
	db := Database()
	row, err := db.Query("INSERT INTO program(id,title,des) values (?,?,?);",prog.ID,prog.TITLE,prog.DESC)
	if err != nil {
				log.Fatal(err)
				}
	log.Println(row,"inserted")
}