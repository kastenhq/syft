package cyclonedxxml

import (
	"github.com/CycloneDX/cyclonedx-go"

	"github.com/kastenhq/syft/syft/formats/common/cyclonedxhelpers"
	"github.com/kastenhq/syft/syft/sbom"
)

const ID sbom.FormatID = "cyclonedx-xml"

func Format() sbom.Format {
	return sbom.NewFormat(
		sbom.AnyVersion,
		encoder,
		cyclonedxhelpers.GetDecoder(cyclonedx.BOMFileFormatXML),
		cyclonedxhelpers.GetValidator(cyclonedx.BOMFileFormatXML),
		ID, "cyclonedx", "cyclone",
	)
}
