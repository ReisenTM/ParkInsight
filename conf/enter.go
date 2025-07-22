package conf

type Config struct {
	DB  []DB `yaml:"db"` //连接的数据库
	Log Log  `yaml:"log"`
}
