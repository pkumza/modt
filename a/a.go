package a

import (
	"github.com/pkumza/modt/c"
)

// Ver is version
func Ver() string {
	if c.Ver() != "v2.0.0" {
		panic("c not v2.0.0")
	}
	return "v2.0.0"
}
