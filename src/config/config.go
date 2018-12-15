package config

type Config struct {
	URL     string `json:"URL,omitempty"`
	TIMEOUT int32  `json:"TIMEOUT,omitempty"`
	HTTPS   bool   `json:"HTTPS,omitempty"`
}

var DefaultConfig = Config{
	URL:     "mainnet.infura.io/v3/c5b6f5890b644d71b4cbf5a544010b90",
	TIMEOUT: 10,
	HTTPS:   true,
}
