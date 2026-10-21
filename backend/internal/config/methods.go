package config

import (
	"fmt"
	"net/url"
)

func (c *Config) GetApiPort() string {
	return c.ApiConfig.port
}

func (c *Config) GetDbDsn() string {
	db := c.DbConfig
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.host, db.port, db.user, db.password, db.name, db.sslMode,
	)
}

func (c *Config) GetDbUrl() string {
	db := c.DbConfig
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(db.user, db.password),
		Host:     db.host + ":" + db.port,
		Path:     db.name,
		RawQuery: "sslmode=" + db.sslMode,
	}
	return u.String()
}
