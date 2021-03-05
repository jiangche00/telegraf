package random

import (
	"math/rand"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

//Random ...
type Random struct {
	x int
}

// RandomConfig ...
var RandomConfig = `
 ## set x
 #x = 1
`

//SampleConfig ...
func (r *Random) SampleConfig() (ret string) {
	ret = RandomConfig
	return
}

//Description ...
func (r *Random) Description() (ret string) {
	ret = "generate random number for demostration purposes"
	return
}

//Gather ...
func (r *Random) Gather(acc telegraf.Accumulator) error {
	if r.x == 0 {
		RandomNumber := 10
		fields := map[string]interface{}{
			"RandomNumber": RandomNumber,
		}
		acc.AddFields("random", fields, nil)
	} else {
		rand.Seed(time.Now().UnixNano())
		RandomNumber := 10 + rand.Intn(50)
		fields := map[string]interface{}{
			"RandomNumber": RandomNumber + r.x,
		}
		acc.AddFields("random", fields, nil)
	}
	return nil
}

func init() {
	inputs.Add("random", func() telegraf.Input { return &Random{x: 1} })
}
