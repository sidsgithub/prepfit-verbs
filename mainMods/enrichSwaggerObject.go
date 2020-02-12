package mainMods
import(
	"github.com/go-openapi/spec"
)

func EnrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "ProgramService",
			Description: "Resource for managing Program",
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "programs",
		Description: "Managing programs"}}}
}