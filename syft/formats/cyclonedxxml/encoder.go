package cyclonedxxml

import (
	"io"

	"github.com/CycloneDX/cyclonedx-go"

	"github.com/kastenhq/syft/syft/formats/common/cyclonedxhelpers"
	"github.com/kastenhq/syft/syft/sbom"
)

func encoder(output io.Writer, s sbom.SBOM) error {
	bom := cyclonedxhelpers.ToFormatModel(s)
	enc := cyclonedx.NewBOMEncoder(output, cyclonedx.BOMFileFormatXML)
	enc.SetPretty(true)

	err := enc.Encode(bom)
	return err
}
