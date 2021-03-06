package config

import (
	"fmt"
	"strings"
	"time"
)

type Server struct {
	Host           string `required:"true"`
	Port           int    `envconfig:"PORT"`
	Scheme         string `default:"http"`
	AllowedDomains string `split_words:"true"`
}

type StaticServer struct {
	Host   string `required:"true"`
	Scheme string `default:"http"`
	Suffix string
}

type Application struct {
	server Server
	db     DB
	StaticServer
}

func AppAddress() string {
	return fmt.Sprintf("%s:%d", app.server.Host, app.server.Port)
}

func StaticServerDomain() string {
	cfg := app.StaticServer
	return fmt.Sprintf("%s://%s/%s", cfg.Scheme, cfg.Host, cfg.Suffix)
}

func AccessControlAllowOrigin() []string {
	domains := app.server.AllowedDomains
	if domains == "" {
		return nil
	}
	return strings.Split(domains, ",")
}

func Database() DB {
	return app.db
}

type DB struct {
	Driver            string `default:"postgres"`
	Host              string `required:"true"`
	User              string `required:"true"`
	Password          string `required:"true"`
	Port              int    `required:"true"`
	MaxIdleConns      int    `split_words:"true" default:"20"`
	MaxOpenConns      int    `split_words:"true" default:"30"`
	MaxConnLifetimeMs int    `split_words:"true" default:"1000"`
	Name              string `split_words:"true" required:"true"`
	SslMode           string `split_words:"true" default:"disable"`
}

func (db DB) MaxConnLifetime() time.Duration {
	return time.Millisecond * time.Duration(db.MaxConnLifetimeMs)
}

func (db DB) URL() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", db.User, db.Password, db.Host, db.Port, db.Name, db.SslMode)
}
