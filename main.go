package main

import "puleeno/tarantula/cli"

func main() {
	tarantula := cli.Tarantula{}

	tarantula.ParseArgs()
	tarantula.InitDownloader()
}
