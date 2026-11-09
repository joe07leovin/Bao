package main

import (
	"context"
	"crypto/tls"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joe07leovin/Bao/internal/health/httpapi"
	"github.com/joe07leovin/Bao/internal/health/query"
	"github.com/joe07leovin/Bao/internal/platform/config"
	"github.com/joe07leovin/Bao/internal/platform/postgres"
)

func main() {
	if err := run(); err != nil {
		slog.Error("BAO stopped", "reason", err)
		os.Exit(1)
	}
}
func run() error {
	c, err := config.Load(os.Getenv)
	if err != nil {
		return err
	}
	certificate, err := tls.LoadX509KeyPair(c.TLSCertFile, c.TLSKeyFile)
	if err != nil {
		return errors.New("unable to load HTTPS certificate and key")
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	pool, err := postgres.Open(ctx, c)
	if err != nil {
		return err
	}
	defer pool.Close()
	server := &http.Server{Addr: c.ListenAddress, Handler: httpapi.Handler(query.Readiness{Database: pool}, c.HealthTimeout), ReadHeaderTimeout: c.ReadHeaderTimeout, ReadTimeout: c.ReadTimeout, WriteTimeout: c.WriteTimeout, IdleTimeout: c.IdleTimeout, MaxHeaderBytes: c.MaxHeaderBytes, TLSConfig: &tls.Config{MinVersion: tls.VersionTLS13, Certificates: []tls.Certificate{certificate}}}
	result := make(chan error, 1)
	go func() { result <- server.ListenAndServeTLS("", "") }()
	select {
	case err := <-result:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return errors.New("HTTPS server failed")
	case <-ctx.Done():
		shutdown, cancel := context.WithTimeout(context.Background(), c.ShutdownTimeout)
		defer cancel()
		if err := server.Shutdown(shutdown); err != nil {
			_ = server.Close()
			return errors.New("HTTPS shutdown deadline exceeded")
		}
		return nil
	}
}
