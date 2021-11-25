package zsxg

import (
	"net/http"
	"time"
)

type Zsxg struct {
	HTTPClient *http.Client
}

func NewZsxg() Zsxg {
	hc := &http.Client{
		Timeout: time.Second * 60 * 5,
	}
	return Zsxg{
		HTTPClient: hc,
	}
}
