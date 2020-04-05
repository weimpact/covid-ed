package config

import "fmt"

type Server struct {
	Host   string `required:"true"`
	Port   int    `envconfig:"PORT"`
	Scheme string `default:"http"`
}

type Application struct {
	server Server
}

func AppAddress() string {
	return fmt.Sprintf("%s:%d", app.server.Host, app.server.Port)
}
