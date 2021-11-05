package datacenter

import (
	"smzdtz-server/datacenter/sina"
)

var (
	Sina sina.Sina
)

func init() {
	Sina = sina.NewSina()
}
