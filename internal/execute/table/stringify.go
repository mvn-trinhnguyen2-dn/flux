package table

import (
	"github.com/mvn-trinhnguyen2-dn/flux"
	"github.com/mvn-trinhnguyen2-dn/flux/execute/table"
)

// Stringify will read a table and turn it into a human-readable string.
func Stringify(t flux.Table) string {
	return table.Stringify(t)
}
