package models

type Config struct {
	Server struct {
		Port  int
		Debug bool
	}
	Auth struct {
		Username string
		Password string
	}
	DB struct {
		Hostname string
		Port     int
		Username string
		Password string
		DBName   string
	}
	Discord struct {
		Bot string
	}
	ReCAPTCHA struct {
		SecretKey string
	}
}
