package model

import (
	"github.com/kastenhq/syft/syft/file"
)

type Secrets struct {
	Location file.Coordinates    `json:"location"`
	Secrets  []file.SearchResult `json:"secrets"`
}
