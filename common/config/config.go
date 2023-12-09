package config

import "time"

type Config struct {
	AppConfig      AppConfig
	ServerConfig   ServerConfig
	DatabaseConfig DatabaseConfig
}

type AppConfig struct {
	RunMode      string `ini:"RUN_MODE"`
	PageSize     int    `ini:"PAGE_SIZE"`
	JwtSecret    string `ini:"JWT_SECRET"`
	JwtIssuer    string `ini:"JWT_ISSUER"`
	JwtExpiresat int64  `ini:"JWT_EXPIRESAT"`
}

type ServerConfig struct {
	HttpPort     int           `ini:"HTTP_PORT"`
	ReadTimeout  time.Duration `ini:"READ_TIMEOUT"`
	WriteTimeout time.Duration `ini:"WRITE_TIMEOUT"`
}

type DatabaseConfig struct {
	Type        string `ini:"TYPE"`
	User        string `ini:"USER"`
	Password    string `ini:"PASSWORD"`
	Host        string `ini:"HOST"`
	Name        string `ini:"NAME"`
	TablePrefix string `ini:"TABLE_PREFIX"`
}
