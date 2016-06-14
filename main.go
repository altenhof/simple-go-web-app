package main

import (
	"os"
	"github.com/pivotal-golang/lager"
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "Hello world"
	}
	logger := lager.NewLogger("simple-go-web-app")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.INFO))
	m := martini.Classic()
	m.Get("/", func(req *http.Request) string {
		logger.Info("Handling request", lager.Data {"request": req.URL})
		return message
	})
	m.Run()
}
