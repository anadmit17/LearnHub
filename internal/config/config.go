package config

import "fmt"

type Config struct {
	ServerHost string
	ServerPort string
	DBConn     string
}

func (c *Config) String() string {
	return fmt.Sprintf(
		"Host: %s, Port: %s, DBConn: %s\n",
		c.ServerHost,
		c.ServerPort,
		c.DBConn,
	)
}
