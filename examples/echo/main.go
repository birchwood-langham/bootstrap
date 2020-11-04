package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"

	"github.com/birchwood-langham/bootstrap/v1/pkg/cmd"
	"github.com/birchwood-langham/bootstrap/v1/pkg/logger"
	"github.com/birchwood-langham/bootstrap/v1/pkg/service"
)

const (
	usage = `echo`
	short = `echo service`
	long  = `echo service echos whatever it is given`
)

func initHttp(ctx context.Context, state service.StateStore) error {
	log := logger.New(logger.ApplicationLogLevel(), logger.ConfiguredLumberjackLogger())

	log.Info("Starting server")

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer, middleware.Timeout(time.Minute))

	r.Get("/echo/{message}", func(w http.ResponseWriter, r *http.Request) {
		message := chi.URLParam(r, "message")
		if message == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	go func() {
		log.Info("Starting echo server on port: 8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Error("HTTP Server terminated", zap.Error(err))
		}
	}()

	return nil
}

func main() {
	app := service.NewApplication().
		AddInitFunc(initHttp)

	app.SetProperties(usage, short, long)

	cmd.Execute(context.Background(), app, service.NewStateStore())
}
