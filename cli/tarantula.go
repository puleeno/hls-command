package cli

import (
	"github.com/bobziuchkovski/writ"
	"github.com/puleeno/tarantula/core/downloader"
)

type Tarantula struct {
	HelpFlag bool     `flag:"help" description:"Display this help message and exit"`
	Download Download `command:"download" alias:"dl" description:"Download a stream via URL"`
	Command  *writ.Command
}

func (t Tarantula) ParseArgs() {
	t.Command = writ.New("Tarantula", &t)

	t.Command.Subcommand("download")

	_, _, err := t.Command.Decode([]string{})

	if err != nil || t.HelpFlag {
		t.Command.ExitHelp(err)
	}
}

func (t Tarantula) InitDownloader() downloader.Downloader {
	downloader := downloader.Downloader{}

	return downloader
}

func (t Tarantula) Run() {
}
