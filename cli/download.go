package cli

type Download struct {
	HelpFlag bool `flag:"h, help" description:""`
}

func (t Download) Run() {
}
