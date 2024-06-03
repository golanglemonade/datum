package bulkgen

import (
	_ "embed"
	"strings"
	"text/template"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
	"github.com/gertd/go-pluralize"
)

//go:embed bulk.gotpl
var bulkTemplate string

// New returns a new plugin
func New() plugin.Plugin {
	return &Plugin{}
}

// Plugin is a gqlgen plugin to generate bulk resolver functions used for mutations
type Plugin struct{}

// Name returns the name of the plugin
func (m *Plugin) Name() string {
	return "bulkgen"
}

// BulkResolverBuild is a struct to hold the objects for the bulk resolver
type BulkResolverBuild struct {
	// Objects is a list of objects to generate bulk resolvers for
	Objects []Object
}

// Object is a struct to hold the object name for the bulk resolver
type Object struct {
	// Name of the object
	Name string
	// PluralName of the object
	PluralName string
}

// GenerateCode generates the bulk resolver code
func (m *Plugin) GenerateCode(data *codegen.Data) error {
	if !data.Config.Resolver.IsDefined() {
		return nil
	}

	return m.generateSingleFile(*data)
}

// generateSingleFile generates the bulk resolver code, this is all done in a single file and
// used by the resolvergen plugin for each bulk resolver
func (m *Plugin) generateSingleFile(data codegen.Data) error {
	inputData := BulkResolverBuild{
		Objects: []Object{},
	}

	for _, f := range data.Schema.Mutation.Fields {
		lowerName := strings.ToLower(f.Name)

		// if the field is a bulk create mutation, add it to the list of objects
		// we skip csv bulk mutations because they will reuse the same functions
		if strings.Contains(lowerName, "bulk") && !strings.Contains(lowerName, "csv") {
			objectName := strings.Replace(f.Name, "createBulk", "", 1)

			inputData.Objects = append(inputData.Objects, Object{
				Name:       objectName,
				PluralName: pluralize.NewClient().Plural(objectName),
			})
		}
	}

	// render the bulk resolver template
	return templates.Render(templates.Options{
		PackageName: data.Config.Resolver.Package,            // use the resolver package
		Filename:    data.Config.Resolver.Dir() + "/bulk.go", // write to the resolver directory
		FileNotice:  `// THIS CODE IS REGENERATED BY github.com/datumforge/datum/pkg/gqlplugin. DO NOT EDIT.`,
		Data:        inputData,
		Funcs: template.FuncMap{
			"toLower": strings.ToLower,
		},
		Packages: data.Config.Packages,
		Template: bulkTemplate,
	})
}