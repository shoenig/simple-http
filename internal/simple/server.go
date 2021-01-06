package simple

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/google/subcommands"
	"gophers.dev/pkgs/extractors/env"
	"gophers.dev/pkgs/loggy"
)

func Server() subcommands.Command {
	return &serverCmd{
		log: loggy.New("server-cli"),
	}
}

type serverCmd struct {
	log  loggy.Logger
	bind string
	port int
}

func (s *serverCmd) Name() string {
	return "server"
}

func (s *serverCmd) Synopsis() string {
	return "Run a simple http server"
}

func (s *serverCmd) Usage() string {
	return "simple-http server"
}

func (s *serverCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&s.bind, "bind", "127.0.0.1", "set server bind address")
	set.IntVar(&s.port, "port", 8999, "set server port")
}

func (s *serverCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if err := env.ParseOS(env.Schema{
		"BIND": env.String(&s.bind, false),
		"PORT": env.Int(&s.port, false),
	}); err != nil {
		s.log.Errorf("failed to parse environment: %v", err)
		return subcommands.ExitFailure
	}

	if err := runServer(s.bind, s.port); err != nil {
		s.log.Errorf("failed to run server: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func runServer(bind string, port int) error {
	log := loggy.New("server")
	log.Infof("starting server, bind: %s, port: %d", bind, port)
	return (&http.Server{
		Addr:           fmt.Sprintf("%s:%d", bind, port),
		Handler:        handler(log),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}).ListenAndServe()
}

func handler(log loggy.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("got request from %s (%s)", r.RemoteAddr, r.Method)
		now := time.Now().Format(time.Kitchen)
		w.WriteHeader(http.StatusOK)
		response := []byte(fmt.Sprintf("the time is %s", now))
		_, _ = w.Write(response)
	}
}
