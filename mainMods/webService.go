package mainMods

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
)

func (u ProgramResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/programs").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"programs"}

	ws.Route(ws.GET("/").To(u.FindAllPrograms).
		// docs
		Doc("get all programs").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]Program{}).
		Returns(200, "OK", []Program{}))

	ws.Route(ws.GET("/{program-id}").To(u.FindProgram).
		// docs
		Doc("get a program").
		Param(ws.PathParameter("program-id", "identifier of the program").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(Program{}). // on the response
		Returns(200, "OK", Program{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("").To(u.CreateProgram).
		// docs
		Doc("create a program").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(Program{})) // from the request


	return ws
}