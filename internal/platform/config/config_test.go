package config

import (
	"strings"
	"testing"
)

func valid() map[string]string {
	return map[string]string{
		"BAO_LISTEN_ADDRESS": "localhost:8443", "BAO_TLS_CERT_FILE": "cert.pem", "BAO_TLS_KEY_FILE": "key.pem", "BAO_DATABASE_URL": "postgres://bao:secret@database.example/bao?sslmode=verify-full",
		"BAO_DATABASE_MAX_CONNECTIONS": "5", "BAO_DATABASE_CONNECT_TIMEOUT": "10s", "BAO_HTTP_READ_HEADER_TIMEOUT": "5s", "BAO_HTTP_READ_TIMEOUT": "15s", "BAO_HTTP_WRITE_TIMEOUT": "15s", "BAO_HTTP_IDLE_TIMEOUT": "60s", "BAO_HTTP_MAX_HEADER_BYTES": "16384", "BAO_HEALTH_TIMEOUT": "2s", "BAO_SHUTDOWN_TIMEOUT": "10s",
	}
}
func TestLoad(t *testing.T) {
	values := valid()
	if _, err := Load(func(k string) string { return values[k] }); err != nil {
		t.Fatal(err)
	}
}
func TestRequiredSettings(t *testing.T) {
	for key := range valid() {
		t.Run(key, func(t *testing.T) {
			values := valid()
			delete(values, key)
			if _, err := Load(func(k string) string { return values[k] }); err == nil {
				t.Fatal("missing setting accepted")
			}
		})
	}
}
func TestRejectUnsafeSettings(t *testing.T) {
	for _, tc := range []struct{ key, value string }{
		{"BAO_DATABASE_URL", "postgres://bao:secret@database.example/bao?sslmode=disable"},
		{"BAO_DATABASE_URL", "postgres://bao:secret@database.example/bao?sslmode=require"},
		{"BAO_DATABASE_URL", "postgres://bao:secret@database.example/bao?sslmode=verify-ca"},
		{"BAO_DATABASE_MAX_CONNECTIONS", "0"}, {"BAO_DATABASE_MAX_CONNECTIONS", "2147483648"},
		{"BAO_HEALTH_TIMEOUT", "0s"}, {"BAO_SHUTDOWN_TIMEOUT", "-1s"},
		{"BAO_HTTP_MAX_HEADER_BYTES", "-1"}, {"BAO_LISTEN_ADDRESS", "localhost"},
		{"BAO_HEALTH_TIMEOUT", "16s"},
	} {
		t.Run(tc.key+tc.value, func(t *testing.T) {
			values := valid()
			values[tc.key] = tc.value
			_, err := Load(func(k string) string { return values[k] })
			if err == nil {
				t.Fatal("unsafe setting accepted")
			}
			if strings.Contains(err.Error(), "secret") {
				t.Fatal("credential leaked")
			}
		})
	}
}
