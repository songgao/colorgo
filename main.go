package main

import (
	"github.com/songgao/go.pipeline"
	"os"
)

var colorizers = map[string]pipeline.LineProcessor{
	"build": build,
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
		p = pipeline.StartPipelineWithCommand("go", args...).ChainLineProcessor(colorizers[os.Args[1]], nil)
	}
	p.PrintAll()
}
