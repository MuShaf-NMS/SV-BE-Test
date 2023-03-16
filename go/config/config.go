package config

type Config struct {
	App_Mode string
	App_Port string
	DB_User  string
	DB_Pass  string
	DB_Name  string
	DB_Host  string
	DB_Port  string
}

func GetConfig() *Config {
	return &Config{
		App_Mode: getVariable("APP_MODE"),
		App_Port: getVariable("APP_PORT"),
		DB_User:  getVariable("DB_USER"),
		DB_Pass:  getVariable("DB_PASS"),
		DB_Name:  getVariable("DB_NAME"),
		DB_Host:  getVariable("DB_HOST"),
		DB_Port:  getVariable("DB_PORT"),
	}
}
