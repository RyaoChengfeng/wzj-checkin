package config

var (
	C *Config
)

type Config struct {
	CheckIn checkin `yaml:"checkin"`
	LogConf logConf `yaml:"logConf"`
	Proxy   string  `yaml:"proxy"`
	Debug   bool    `yaml:"debug"`
}

type checkin struct {
	CheckInUrl string  `yaml:"checkin_url"`
	OpenID     string  `yaml:"openid"`
	Lon        float64 `yaml:"lon"`
	Lat        float64 `yaml:"lat"`
}

type logConf struct {
	LogPath string `yaml:"log_path"`
	LogFile string `yaml:"log_file"`
}
