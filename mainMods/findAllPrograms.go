package mainMods

import (
	"github.com/emicklei/go-restful"
)

func (u ProgramResource) FindAllPrograms(request *restful.Request, response *restful.Response) {
	list := []Program{}
	for _, each := range u.Programs {
		list = append(list, each)
	}
	response.WriteEntity(list)
	// Database(list,"findAllPrograms")
}