package main

import(
	"log"
	"net/http"
	"github.com/emicklei/go-restful"
	"github.com/user/prepFitness/mainMods"
	"github.com/emicklei/go-restful-openapi"
)

func main(){
	// Database()
	u := mainMods.ProgramResource{map[string]mainMods.Program{}}
	restful.DefaultContainer.Add(u.WebService())

	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: mainMods.EnrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))
	log.Printf("start listening on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}