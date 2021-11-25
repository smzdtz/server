package model

type Setting struct {
	Db struct {
		Engine       string
		Host         string
		Port         int
		User         string
		Password     string
		Database     string
		Prefix       string
		Charset      string
		MaxIdleConns int
		MaxOpenConns int
	}
	AllowIps      string
	AppName       string
	ApiKey        string
	ApiSecret     string
	ApiSignEnable bool

	EnableTLS bool
	CAFile    string
	CertFile  string
	KeyFile   string

	ConcurrencyQueue int
	AuthSecret       string
}
