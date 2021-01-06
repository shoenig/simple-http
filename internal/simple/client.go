package simple

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/subcommands"
	"gophers.dev/pkgs/extractors/env"
	"gophers.dev/pkgs/loggy"
)

func Client() subcommands.Command {
	return &clientCmd{
		log: loggy.New("client-cli"),
	}
}

type clientCmd struct {
	log     loggy.Logger
	address string
	port    int
}

func (c *clientCmd) Name() string {
	return "client"
}

func (c *clientCmd) Synopsis() string {
	return "Run a simple http client"
}

func (c *clientCmd) Usage() string {
	return "simple-http client"
}

func (c *clientCmd) SetFlags(set *flag.FlagSet) {
	set.StringVar(&c.address, "address", "127.0.0.1", "set request address")
	set.IntVar(&c.port, "port", 8999, "set request port")
}

func (c *clientCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if err := env.ParseOS(env.Schema{
		"ADDRESS": env.String(&c.address, false),
		"PORT":    env.Int(&c.port, false),
	}); err != nil {
		c.log.Errorf("failed to parse environment: %v", err)
		return subcommands.ExitFailure
	}

	if err := runClient(c.address, c.port); err != nil {
		c.log.Errorf("failed to run client: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}

func runClient(address string, port int) error {
	log := loggy.New("client")
	log.Infof("starting client, address: %s, port: %d", address, port)

	for range time.Tick(2 * time.Second) {
		u := fmt.Sprintf("http://%s:%d", address, port)
		log.Infof("sending request to %s ...", u)

		response, err := http.Get(u)
		if err != nil {
			log.Errorf(" -> GET error: %v", err)
			continue
		}
		log.Infof(" -> GET response code: (%d)", response.StatusCode)
		if b, err := ioutil.ReadAll(response.Body); err != nil {
			log.Errorf(" -> GET response error: %v", err)
		} else {
			log.Infof(" -> GET response: %s", string(b))
		}
	}
	return nil
}
