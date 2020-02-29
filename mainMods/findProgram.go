package mainMods

import (
	"github.com/emicklei/go-restful"
	"log"
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

	db := Database()

	result,err := db.Query("select * from program where id = ?",usr.ID)
			if err != nil {
				log.Fatal(err)
			}
			for result.Next(){
				err:=result.Scan(&usr.ID,&usr.TITLE,&usr.DESC)
				if err!=nil{
					panic(err.Error())
				}
				log.Println("program ran successfully",usr.TITLE)
			}
}