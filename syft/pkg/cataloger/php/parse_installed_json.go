package php

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/kastenhq/syft/syft/artifact"
	"github.com/kastenhq/syft/syft/file"
	"github.com/kastenhq/syft/syft/pkg"
	"github.com/kastenhq/syft/syft/pkg/cataloger/generic"
)

var _ generic.Parser = parseComposerLock

// Note: composer version 2 introduced a new structure for the installed.json file, so we support both
type installedJSONComposerV2 struct {
	Packages []parsedData `json:"packages"`
}

func (w *installedJSONComposerV2) UnmarshalJSON(data []byte) error {
	type compv2 struct {
		Packages []parsedData `json:"packages"`
	}
	compv2er := new(compv2)
	err := json.Unmarshal(data, &compv2er)
	if err != nil {
		// If we had an err	or, we may be dealing with a composer v.1 installed.json
		// which should be all arrays
		var packages []parsedData
		err := json.Unmarshal(data, &packages)
		if err != nil {
			return err
		}
		w.Packages = packages
		return nil
	}
	w.Packages = compv2er.Packages
	return nil
}

// parseInstalledJSON is a parser function for Composer.lock contents, returning "Default" php packages discovered.
func parseInstalledJSON(_ file.Resolver, _ *generic.Environment, reader file.LocationReadCloser) ([]pkg.Package, []artifact.Relationship, error) {
	var pkgs []pkg.Package
	dec := json.NewDecoder(reader)

	for {
		var lock installedJSONComposerV2
		if err := dec.Decode(&lock); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, nil, fmt.Errorf("failed to parse installed.json file: %w", err)
		}
		for _, pkgMeta := range lock.Packages {
			pkgs = append(
				pkgs,
				newComposerLockPackage(pkgMeta,
					reader.Location.WithAnnotation(pkg.EvidenceAnnotationKey, pkg.PrimaryEvidenceAnnotation),
				),
			)
		}
	}

	return pkgs, nil, nil
}
