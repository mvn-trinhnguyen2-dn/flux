//go:build !go1.12
// +build !go1.12

package runtime

import (
	"github.com/mvn-trinhnguyen2-dn/flux/values"
)

func Version() (values.Value, error) {
	return nil, errBuildInfoNotPresent
}
