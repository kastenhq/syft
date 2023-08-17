package dotnet

import (
	"github.com/kastenhq/syft/syft/pkg/cataloger/generic"
)

// NewDotnetDepsCataloger returns a new Dotnet cataloger object base on deps json files.
func NewDotnetDepsCataloger() *generic.Cataloger {
	return generic.NewCataloger("dotnet-deps-cataloger").
		WithParserByGlobs(parseDotnetDeps, "**/*.deps.json")
}

func NewDotnetPortableExecutableCataloger() *generic.Cataloger {
	return generic.NewCataloger("dotnet-portable-executable-cataloger").
		WithParserByGlobs(parseDotnetPortableExecutable, "**/*.dll", "**/*.exe")
}
