package b

import (
	"github.com/pkumza/modt/c"
)

// Ver is version
func Ver() string {
	if c.Ver() != "v1.0.0" {
		panic("c not v1.0.0")
	}
	return "v2.0.0"
}
