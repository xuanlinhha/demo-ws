package global

type Config struct {
	Address string
	MySql   MysqlConfig
}

type MysqlConfig struct {
	Address   string
	Username  string
	Passsword string
	DB        string
}

var Cfg Config
