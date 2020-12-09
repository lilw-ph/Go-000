package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func startHttpServer(ctx context.Context, url string, name string) error {

	smux := http.NewServeMux()
	smux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "http server %s\n", name)
	})

	srv := http.Server{
		Addr:    url,
		Handler: smux,
	}

	go func() {
		<-ctx.Done()
		fmt.Printf("[server %s] done!\n", name)
		srv.Shutdown(context.Background())
	}()

	return srv.ListenAndServe()
}

func listenSig(ctx context.Context, sig ...os.Signal) error {
	s := make(chan os.Signal, 1)

	signal.Notify(s, sig...)

	for {
		select {
		case <-ctx.Done():
			return errors.New("[signal goroutine] done!")
		case d := <-s:
			fmt.Println("signal is ", d)
			return errors.New("[signal goroutine] get signal")
		}
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return startHttpServer(ctx, "127.0.0.1:8080", "1")
	})

	g.Go(func() error {
		return startHttpServer(ctx, "127.0.0.1:8081", "2")
	})

	g.Go(func() error {
		return listenSig(ctx, syscall.SIGINT, syscall.SIGILL)
	})

	g.Wait()

	fmt.Println("finished")
}
