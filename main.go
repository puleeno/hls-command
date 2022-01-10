package main

import (
	"github.com/puleeno/tarantula/cli"
)

func main() {
	tarantula := cli.Tarantula{}

	tarantula.ParseArgs()
	tarantula.Execute()
}
