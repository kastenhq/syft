package syft

import (
	"github.com/kastenhq/syft/syft/formats"
	"github.com/kastenhq/syft/syft/formats/cyclonedxjson"
	"github.com/kastenhq/syft/syft/formats/cyclonedxxml"
	"github.com/kastenhq/syft/syft/formats/github"
	"github.com/kastenhq/syft/syft/formats/spdxjson"
	"github.com/kastenhq/syft/syft/formats/spdxtagvalue"
	"github.com/kastenhq/syft/syft/formats/syftjson"
	"github.com/kastenhq/syft/syft/formats/table"
	"github.com/kastenhq/syft/syft/formats/template"
	"github.com/kastenhq/syft/syft/formats/text"
	"github.com/kastenhq/syft/syft/sbom"
)

// these have been exported for the benefit of API users
// TODO: deprecated: now that the formats package has been moved to syft/formats, will be removed in v1.0.0
const (
	JSONFormatID          = syftjson.ID
	TextFormatID          = text.ID
	TableFormatID         = table.ID
	CycloneDxXMLFormatID  = cyclonedxxml.ID
	CycloneDxJSONFormatID = cyclonedxjson.ID
	GitHubFormatID        = github.ID
	SPDXTagValueFormatID  = spdxtagvalue.ID
	SPDXJSONFormatID      = spdxjson.ID
	TemplateFormatID      = template.ID
)

// TODO: deprecated, moved to syft/formats/formats.go. will be removed in v1.0.0
func FormatIDs() (ids []sbom.FormatID) {
	return formats.AllIDs()
}

// TODO: deprecated, moved to syft/formats/formats.go. will be removed in v1.0.0
func FormatByID(id sbom.FormatID) sbom.Format {
	return formats.ByNameAndVersion(string(id), "")
}

// TODO: deprecated, moved to syft/formats/formats.go. will be removed in v1.0.0
func FormatByName(name string) sbom.Format {
	return formats.ByName(name)
}

// TODO: deprecated, moved to syft/formats/formats.go. will be removed in v1.0.0
func IdentifyFormat(by []byte) sbom.Format {
	return formats.Identify(by)
}
