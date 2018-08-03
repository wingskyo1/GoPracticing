package config

//Config :結構
type Config struct {
	DB *DBConfig
}

//DBConfig : DB連線字串結構
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

//GetConfig :return 指向DB 結構的指標
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "123459_abcd",
			Name:     "go",
			Charset:  "utf8",
		},
	}
}
