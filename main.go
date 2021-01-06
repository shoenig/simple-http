package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"gophers.dev/cmds/simple-http/internal/simple"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "about")
	subcommands.Register(subcommands.FlagsCommand(), "about")
	subcommands.Register(subcommands.CommandsCommand(), "about")

	subcommands.Register(simple.Server(), "")
	subcommands.Register(simple.Client(), "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
