package main

import (
	"os"
	"os/signal"
	"context"
	"net/http"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func startService(addr string, ctx context.Context) error {
	serv := http.Server{
		Addr:    addr,
		Handler: nil,
	}

	go func() {
		<-ctx.Done()
		_ = serv.Shutdown(context.Background())
	}()

	return serv.ListenAndServe()
}

func main() {
	signalCh := make(chan os.Signal)

	egroup, ctx := errgroup.WithContext(context.Background())
	signal.Notify(signalCh, os.Interrupt)

	egroup.Go(func() error {
		select {
		case <-signalCh:
			return errors.New("stop service")
		case <-ctx.Done():
			return errors.New("start service error")
		}
	})

	egroup.Go(func() error {
		return startService("0.0.0.0:8080", ctx)
	})

	egroup.Go(func() error {
		return startService("0.0.0.0:8081", ctx)
	})

	if err := egroup.Wait(); err != nil {
		return
	}

}
