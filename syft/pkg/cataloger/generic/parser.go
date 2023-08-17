package generic

import (
	"github.com/kastenhq/syft/syft/artifact"
	"github.com/kastenhq/syft/syft/file"
	"github.com/kastenhq/syft/syft/linux"
	"github.com/kastenhq/syft/syft/pkg"
)

type Environment struct {
	LinuxRelease *linux.Release
}

type Parser func(file.Resolver, *Environment, file.LocationReadCloser) ([]pkg.Package, []artifact.Relationship, error)
