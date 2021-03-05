package random

import (
	"math/rand"
	"testing"
	"time"

	"github.com/influxdata/telegraf/testutil"
)

func TestRandom(t *testing.T) {
	s := Random{
		x: 0,
	}

	var acc testutil.Accumulator
	s.Gather(&acc)

	rand.Seed(time.Now().UnixNano())
	RandomNumber := 10

	fields := map[string]interface{}{
		"RandomNumber": RandomNumber,
	}

	acc.AssertContainsFields(t, "random", fields)
}
