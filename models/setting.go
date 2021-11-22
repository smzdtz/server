package models

type InstallForm struct {
	// DbType               string `json:"dbType" binding:"required,oneof=mysql postgres"`
	// DbHost               string `json:"dbHost" binding:"required,min=0,max=50"`
	// DbPort               int    `json:"dbPort" binding:"required,gt=0,lte=65535"`
	// DbUsername           string `json:"dbUsername" binding:"required,max=50"`
	// DbPassword           string `json:"dbPassword" binding:"required,max=30"`
	// DbName               string `json:"dbName" binding:"required,max=50"`
	// DbTablePrefix        string `json:"dbTablePrefix" binding:"max=20"`
	AdminUsername        string `json:"adminUsername" binding:"required,min=3"`
	AdminPassword        string `json:"adminPassword" binding:"required,min=6"`
	ConfirmAdminPassword string `json:"confirmAdminPassword" binding:"required,min=6"`
	AdminEmail           string `json:"adminEmail" binding:"required,email,max=50"`
}

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
