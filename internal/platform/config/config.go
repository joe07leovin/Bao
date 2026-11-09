package config

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ListenAddress, TLSCertFile, TLSKeyFile, DatabaseURL                                                               string
	DatabaseMaxConnections                                                                                            int32
	MaxHeaderBytes                                                                                                    int
	DatabaseConnectTimeout, ReadHeaderTimeout, ReadTimeout, WriteTimeout, IdleTimeout, HealthTimeout, ShutdownTimeout time.Duration
}

// Load deliberately has no fallback settings. Error messages never include values.
func Load(lookup func(string) string) (Config, error) {
	var c Config
	required := func(name string) (string, error) {
		value := strings.TrimSpace(lookup(name))
		if value == "" {
			return "", fmt.Errorf("%s is required", name)
		}
		return value, nil
	}
	for _, field := range []struct {
		name   string
		target *string
	}{
		{"BAO_LISTEN_ADDRESS", &c.ListenAddress}, {"BAO_TLS_CERT_FILE", &c.TLSCertFile}, {"BAO_TLS_KEY_FILE", &c.TLSKeyFile}, {"BAO_DATABASE_URL", &c.DatabaseURL},
	} {
		value, err := required(field.name)
		if err != nil {
			return c, err
		}
		*field.target = value
	}
	if _, port, err := net.SplitHostPort(c.ListenAddress); err != nil || port == "" {
		return c, errors.New("BAO_LISTEN_ADDRESS must contain a host and port")
	}
	u, err := url.Parse(c.DatabaseURL)
	if err != nil || (u.Scheme != "postgres" && u.Scheme != "postgresql") || u.Hostname() == "" || u.Query().Get("sslmode") != "verify-full" {
		return c, errors.New("BAO_DATABASE_URL must be a PostgreSQL URL with sslmode=verify-full")
	}
	for _, field := range []struct {
		name   string
		target *time.Duration
	}{
		{"BAO_DATABASE_CONNECT_TIMEOUT", &c.DatabaseConnectTimeout}, {"BAO_HTTP_READ_HEADER_TIMEOUT", &c.ReadHeaderTimeout}, {"BAO_HTTP_READ_TIMEOUT", &c.ReadTimeout}, {"BAO_HTTP_WRITE_TIMEOUT", &c.WriteTimeout}, {"BAO_HTTP_IDLE_TIMEOUT", &c.IdleTimeout}, {"BAO_HEALTH_TIMEOUT", &c.HealthTimeout}, {"BAO_SHUTDOWN_TIMEOUT", &c.ShutdownTimeout},
	} {
		value, err := required(field.name)
		if err != nil {
			return c, err
		}
		duration, err := time.ParseDuration(value)
		if err != nil || duration <= 0 {
			return c, fmt.Errorf("%s must be a positive duration", field.name)
		}
		*field.target = duration
	}
	maxConnections, err := required("BAO_DATABASE_MAX_CONNECTIONS")
	if err != nil {
		return c, err
	}
	connections, err := strconv.ParseInt(maxConnections, 10, 32)
	if err != nil || connections <= 0 {
		return c, errors.New("BAO_DATABASE_MAX_CONNECTIONS must be a positive int32")
	}
	c.DatabaseMaxConnections = int32(connections)
	maxHeaders, err := required("BAO_HTTP_MAX_HEADER_BYTES")
	if err != nil {
		return c, err
	}
	c.MaxHeaderBytes, err = strconv.Atoi(maxHeaders)
	if err != nil || c.MaxHeaderBytes <= 0 {
		return c, errors.New("BAO_HTTP_MAX_HEADER_BYTES must be a positive integer")
	}
	if c.HealthTimeout >= c.WriteTimeout {
		return c, errors.New("BAO_HEALTH_TIMEOUT must be less than BAO_HTTP_WRITE_TIMEOUT")
	}
	return c, nil
}
