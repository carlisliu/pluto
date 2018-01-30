package config

func DefaultConfig() Option {
	return Option{
		Name:    "pluto",
		Version: "0.0.1",
		Db: DB{
			UserName: "carlis",
			Password: "carlis",
			Ip:       "192.168.1.2",
			Port:     3300,
		},
	}
}
