package args

import (
	"flag"
)

type Args struct {
	Limit int
}

func (a *Args) ParseArgs() {
	flag.IntVar(&a.Limit, "limit", 1, "limit of jobs. Default 1")
	flag.Parse()
}
