/*
Package golang provides a concrete Cataloger implementation for go.mod files.
*/
package golang

import (
	"github.com/kastenhq/syft/internal"
	"github.com/kastenhq/syft/syft/artifact"
	"github.com/kastenhq/syft/syft/event/monitor"
	"github.com/kastenhq/syft/syft/file"
	"github.com/kastenhq/syft/syft/pkg"
	"github.com/kastenhq/syft/syft/pkg/cataloger/generic"
)

// NewGoModFileCataloger returns a new Go module cataloger object.
func NewGoModFileCataloger(opts GoCatalogerOpts) pkg.Cataloger {
	c := goModCataloger{
		licenses: newGoLicenses(opts),
	}
	return &progressingCataloger{
		progress: c.licenses.progress,
		cataloger: generic.NewCataloger("go-mod-file-cataloger").
			WithParserByGlobs(c.parseGoModFile, "**/go.mod"),
	}
}

// NewGoModuleBinaryCataloger returns a new Golang cataloger object.
func NewGoModuleBinaryCataloger(opts GoCatalogerOpts) pkg.Cataloger {
	c := goBinaryCataloger{
		licenses: newGoLicenses(opts),
	}
	return &progressingCataloger{
		progress: c.licenses.progress,
		cataloger: generic.NewCataloger("go-module-binary-cataloger").
			WithParserByMimeTypes(c.parseGoBinary, internal.ExecutableMIMETypeSet.List()...),
	}
}

type progressingCataloger struct {
	progress  *monitor.CatalogerTask
	cataloger *generic.Cataloger
}

func (p *progressingCataloger) Name() string {
	return p.cataloger.Name()
}

func (p *progressingCataloger) Catalog(resolver file.Resolver) ([]pkg.Package, []artifact.Relationship, error) {
	defer p.progress.SetCompleted()
	return p.cataloger.Catalog(resolver)
}
