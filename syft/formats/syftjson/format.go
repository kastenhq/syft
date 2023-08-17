package syftjson

import (
	"github.com/kastenhq/syft/internal"
	"github.com/kastenhq/syft/syft/sbom"
)

const ID sbom.FormatID = "syft-json"

func Format() sbom.Format {
	return sbom.NewFormat(
		internal.JSONSchemaVersion,
		encoder,
		decoder,
		validator,
		ID, "json", "syft",
	)
}
