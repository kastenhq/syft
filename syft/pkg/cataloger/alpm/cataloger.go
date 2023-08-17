package alpm

import (
	"github.com/kastenhq/syft/syft/pkg"
	"github.com/kastenhq/syft/syft/pkg/cataloger/generic"
)

const CatalogerName = "alpmdb-cataloger"

func NewAlpmdbCataloger() *generic.Cataloger {
	return generic.NewCataloger(CatalogerName).
		WithParserByGlobs(parseAlpmDB, pkg.AlpmDBGlob)
}
