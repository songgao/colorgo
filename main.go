package main

import (
	"os"

	"github.com/songgao/go.pipeline"
)

var colorizers = map[string]pipeline.LineProcessor{
	"build": buildOrTest,
	"test":  buildOrTest,
}

func main() {
	var p *pipeline.Pipeline

	if len(os.Args) == 1 {
		p = pipeline.StartPipelineWithCommand("go")
	} else {
		args := make([]interface{}, len(os.Args[1:]))
		for i, v := range os.Args[1:] {
			args[i] = v
		}
		p = pipeline.StartPipelineWithCommand("go", args...).ChainLineProcessor(colorizers[os.Args[1]], colorizers[os.Args[1]])
	}
	p.PrintAll()
	errs := p.Errors()
	if len(errs) > 0 {
		os.Exit(1)
	}
}
