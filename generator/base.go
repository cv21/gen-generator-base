package generator

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cv21/gen/pkg"
	"github.com/pkg/errors"
	"github.com/vetcher/go-astra/types"
)

const (
	mainFilePath        = "./main.go"
	descriptionFilePath = "./DESCRIPTION.md"
	moduleRepository    = "github.com/cv21/gen-generator-base"
	moduleQuery         = "1.0.0"

	prefixExample = "// Example: "

	lineBreak     = "\n"
	prefixComment = "// "
)

type (
	generatorParams struct {
		// Readme headline.
		// Example: Mock
		GeneratorName string `json:"generator_name"`

		// It is necessary for plugin registration.
		// Also it useful for building json config example.
		// Example: github.com/cv21/gen-generator-mock
		ModuleRepository string `json:"module_repository"`

		// It is necessary for plugin registration along with ModuleRepository.
		// Example: 1.0.0
		ModuleQuery string `json:"module_query"`

		// It is name of params structure which holds all generator params.
		// Example: generatorParams
		ParamsStructureName string `json:"params_structure_name"`
	}

	baseGenerator struct {
	}
)

func NewGenerator() pkg.Generator {
	return &baseGenerator{}
}

// This is convenient way to generate files.
func (m *baseGenerator) Generate(p *pkg.GenerateParams) (*pkg.GenerateResult, error) {
	params := &generatorParams{}
	err := json.Unmarshal(p.Params, params)
	if err != nil {
		return nil, err
	}

	parsedGeneratorParams := pkg.FindStructure(p.File, params.ParamsStructureName)
	if parsedGeneratorParams == nil {
		return nil, errors.New("could not find params structure")
	}

	parsedGenerateMethod := pkg.FindMethod(p.File, "Generate")
	if parsedGenerateMethod == nil {
		return nil, errors.New("could not find generate method")
	}

	return &pkg.GenerateResult{
		Files: []pkg.GenerateResultFile{
			{
				Path:    mainFilePath,
				Content: m.generateMainFile(params),
			},
			{
				Path:    descriptionFilePath,
				Content: m.generateDescFile(params, parsedGeneratorParams, parsedGenerateMethod),
			},
		},
	}, nil
}

func (m *baseGenerator) generateMainFile(p *generatorParams) []byte {
	return []byte(fmt.Sprintf(
		`// File generated by gen. DO NOT EDIT.
// Generator plugin %s %s
package main

import (
	"%s/generator"

	"github.com/cv21/gen/pkg"
	plugin "github.com/hashicorp/go-plugin"
)

const (
	pluginRepoURL = "%s"
)

func main() {
	pkg.RegisterGobTypes()
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pkg.DefaultHandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginRepoURL: &pkg.NetRPCWorker{Impl: generator.NewGenerator()},
		},
	})
}
`, moduleRepository, moduleQuery, p.ModuleRepository, buildRepositoryWithQuery(p.ModuleRepository, p.ModuleQuery)))
}

func (m *baseGenerator) generateDescFile(p *generatorParams, parsedGeneratorParams *types.Struct, parsedGeneratorMethod *types.Method) []byte {
	f := fmt.Sprintf(`# %s
%s

#### Generator Parameters
| Parameter Key | Example | Description |
| --- | --- | --- |
`, p.GeneratorName, strings.Join(removeStringsPrefix(GetStringsWithoutPrefix(parsedGeneratorMethod.Docs, prefixComment), prefixExample), lineBreak))

	for _, field := range parsedGeneratorParams.Fields {
		f += generateDescTablePart(&field)
	}

	f += fmt.Sprintf(`
#### Config Example:

`+"```"+`json
{
    "files": [
        {
            "path": "...",
            "generators": [
                {
                    "repository": "%s",
                    "version": "%s",
                    "params": {
                        %s
                    }
                }
            ]
        }
    ]
}
`+"```"+`

`, p.ModuleRepository, p.ModuleQuery, buildJsonConfigExample(parsedGeneratorParams))

	return []byte(f)
}

func GetStringByPrefix(docs []string, prefix string) string {
	for _, doc := range docs {
		if strings.HasPrefix(doc, prefix) {
			return doc
		}
	}

	return ""
}

func GetStringsWithoutPrefix(docs []string, prefix string) []string {
	// Filtering without additional allocation.
	b := docs[:0]
	for _, doc := range docs {
		if !strings.HasPrefix(doc, prefix) {
			b = append(b, doc)
		}
	}

	return b
}

func buildJsonConfigExample(s *types.Struct) (r string) {
	for i, f := range s.Fields {
		p := fmt.Sprintf(`"%s":"%s"`, f.Tags["json"][0], strings.TrimPrefix(GetStringByPrefix(f.Docs, prefixExample), prefixExample))
		if i < len(s.Fields)-1 {
			p += "," + lineBreak
		}

		if i != 0 {
			p = "                        " + p
		}

		r += p
	}

	return
}

func removeStringsPrefix(s []string, prefix string) (r []string) {
	for _, str := range s {
		r = append(r, strings.TrimPrefix(str, prefix))
	}

	return
}

func generateDescTablePart(field *types.StructField) string {
	return fmt.Sprintf(
		"|%s|%s|%s|\n",
		field.Tags["json"][0],
		strings.TrimPrefix(GetStringByPrefix(field.Docs, prefixExample), prefixExample),
		strings.Join(
			removeStringsPrefix(
				GetStringsWithoutPrefix(field.Docs, prefixExample),
				prefixComment,
			), lineBreak,
		),
	)
}

func buildRepositoryWithQuery(repo, query string) string {
	return fmt.Sprintf("%s@%s", repo, query)
}
