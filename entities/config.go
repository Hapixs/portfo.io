package entities

type Config struct {
	ListenPort string `json:"listenPort"`
	ListenIp   string `json:"listenIp"`
	Database   struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
	DefaultApiKey string `json:"defaultApiKey"`
}
