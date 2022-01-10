package command

import (
	"errors"
	"github.com/puleeno/tarantula/core"
	"net/url"

	"github.com/bobziuchkovski/writ"
)

type Download struct {
	HelpFlag bool `flag:"h, help" description:""`
	LiveFlag bool `flag:"l, live" description:"Support download live stream video"`
}

func (t *Download) Run(positional []string, path *writ.Path) {
	argsCounter := len(positional)
	if argsCounter < 1 {
		path.Last().ExitHelp(errors.New("URL parameter is missing"))
	}
	if argsCounter < 2 {
		positional = append(positional, "output.mp4")
	}

	// Create variables preparing to download the stream
	downloadUrl := positional[0]
	//outputFileName := positional[1]

	u, err := url.Parse(downloadUrl)

	if err != nil {
		panic(err)
	}
	if len(u.Host) <= 0 {
		panic(errors.New("the URL of stream is invalid format"))
	}

	detector := core.Detector{}
	format := detector.DetectFromURL(downloadUrl)
}
