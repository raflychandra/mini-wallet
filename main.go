package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-mini-wallet/message"
	"go-mini-wallet/router"
	"go-mini-wallet/service"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var (
	Environment *string
	Http        *bool
)

func init() {
	Environment = flag.String("config_name", "local", "define environment")
	Http = flag.Bool("http", false, "define environment")
	flag.Parse()
}

func main() {
	mainHandler := service.MakeHandler(Environment, Http)
	var wg sync.WaitGroup

	if *Http {
		wg.Add(1)
		go func() {
			defer wg.Done()
			RunWithHTTP(mainHandler)
		}()
	}

	wg.Wait()
}

func RunWithHTTP(mainHandler *service.HandlerSetup) {
	handlerRouter := router.NewHandlerRouter(mainHandler)
	r := handlerRouter.ListRouter()
	messageRun := fmt.Sprintf("go-mini-wallet run on %s environment", mainHandler.Env.Application.Env)
	message.Log(logrus.InfoLevel, messageRun, "SETUP ENV HTTP")

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", mainHandler.Env.Application.Port),
		Handler: r,
	}

	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		message.Log(logrus.InfoLevel, "We received an interrupt signal, shut down.", "SUCCESS STOPPING HTTP go-mini-wallet")
		if err := srv.Shutdown(context.Background()); err != nil {
			message.Log(logrus.ErrorLevel, fmt.Sprintf("HTTP server Shutdown: %v", err), "ERROR STOPPING HTTP go-mini-wallet")
		}
		close(idleConsClosed)
		message.Log(logrus.InfoLevel, "Bye.", "SUCCESS STOPPING HTTP go-mini-wallet")
	}()

	message.Log(logrus.InfoLevel, fmt.Sprintf("Listening on port %s", mainHandler.Env.Application.Port), "SUCCESS RUNNING HTTP go-mini-wallet")

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		message.Log(logrus.FatalLevel, err.Error(), "ERROR START HTTP go-mini-wallet")
	}
	<-idleConsClosed
}
