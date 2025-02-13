package pkg

import "github.com/kastenhq/syft/syft/artifact"

func NewRelationships(catalog *Collection) []artifact.Relationship {
	rels := RelationshipsByFileOwnership(catalog)
	rels = append(rels, RelationshipsEvidentBy(catalog)...)
	return rels
}
