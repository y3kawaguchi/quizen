package location

import (
	"time"
)

var (
	jp *time.Location
)

func init() {
	jp = time.FixedZone("JST", 9*60*60)
}

// JP ...
func JP() *time.Location {
	return jp
}
