package cninfo

import (
	"net/http"
	"time"
)

type CnInfo struct {
	HTTPClient *http.Client
}

func NewCninfo() CnInfo {
	hc := &http.Client{
		Timeout: time.Second * 60 * 5,
	}
	return CnInfo{
		HTTPClient: hc,
	}
}
