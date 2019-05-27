package generator

import (
	"testing"

	"github.com/cv21/gen/pkg"
)

func TestMock_Generate(t *testing.T) {
	testCases := []pkg.TestCase{
		{
			Name: "base",
			GeneratedFilePaths: []string{
				"./DESCRIPTION.md",
				"./main.go",
			},
			Params: `
{
	"repository": "github.com/cv21/gen-generator-base@v1.0.0",
	"params_structure_name": "generatorParams"
}`,
		},
	}

	pkg.RunTestCases(t, testCases, NewGenerator())
}
