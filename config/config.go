package config

type T struct {
	Mysql struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	} `json:"mysql"`
	Redis struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"redis"`
}
