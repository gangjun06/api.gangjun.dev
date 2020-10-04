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
}
