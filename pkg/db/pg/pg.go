package pg

import "fmt"

type Configuration struct {
	host     string
	port     int
	user     string
	password string
	database string
	sslMode  string
}

func (c Configuration) PgConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.user, c.password, c.host, c.port, c.database, c.sslMode)
}

func NewConfiguration(host string, port int, user string, password string, database string, sslMode string) Configuration {
	return Configuration{
		host, port, user, password, database, sslMode,
	}
}

func (c Configuration) Host() string {
	return c.host
}

func (c Configuration) Port() int {
	return c.port
}

func (c Configuration) User() string {
	return c.user
}

func (c Configuration) Password() string {
	return c.password
}

func (c Configuration) Database() string {
	return c.database
}

func (c Configuration) SslMode() string {
	return c.sslMode
}
