package cli

import (
	"os"

	"github.com/bobziuchkovski/writ"
	"github.com/puleeno/tarantula/cli/command"
)

type Tarantula struct {
	HelpFlag bool             `flag:"help" description:"Display this help message and exit"`
	Download command.Download `command:"download" alias:"dl" description:"Download a stream video URL"`

	command    *writ.Command
	path       writ.Path
	positional []string
}

func (t *Tarantula) ParseArgs() {
	t.command = writ.New("tarantula", t)

	path, positional, err := t.command.Decode(os.Args[1:])

	if err != nil {
		path.Last().ExitHelp(err)
	}

	t.path = path
	t.positional = positional

	// Create the help messages
	t.setupHelp()
}

func (t *Tarantula) setupHelp() {
	t.command.Help.Usage = "Usage: tarantula COMMAND [OPTION]... [ARG]..."
	t.command.Subcommand("download").Help.Usage = "Usage: tarantula download <Stream URL> <Output file name> [--live]"
}

func (t *Tarantula) Execute() {
	switch t.path.String() {
	case "tarantula download":
		t.Download.Run(t.positional, &t.path)
		break
	default:
		t.command.ExitHelp(nil)
		break
	}
}
