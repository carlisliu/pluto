package config

type Option struct {
	Name    string
	Version string
	Db      DB
}

type DB struct {
	UserName string
	Password string
	Ip       string
	Port     int
}
