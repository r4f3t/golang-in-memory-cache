package config

type (
	Configuration struct {
		PostgreDbSettings DbSettings `json:"PostgreDbSettings"`
	}

	DbSettings struct {
		Host     string `json:"Host"`
		User     string `json:"User"`
		Password string `json:"Password"`
		DbName   string `json:"DbName"`
		Port     string `json:"Port"`
	}
)
