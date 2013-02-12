package main

var colorizers = map[string]colorizer{
	"build": getColorizer(build),
}

type colorizer func() (chan<- string, chan<- string, <-chan string, <-chan string)
type lineProcessor func(string) string

func getColorizer(c lineProcessor) colorizer {
	return func() (chan<- string, chan<- string, <-chan string, <-chan string) {
		ch_stdout := make(chan string, 4)
		ch_stderr := make(chan string, 4)
		ch_colorout := make(chan string, 4)
		ch_colorerr := make(chan string, 4)
		h := func(ch_in <-chan string, ch_out chan<- string) {
			ok := true
			var line string
			for ok {
				line, ok = <-ch_in
				if ok {
					ch_out <- c(line)
				}
			}
			close(ch_out)
		}
		go h(ch_stdout, ch_colorout)
		go h(ch_stderr, ch_colorerr)
		return ch_stdout, ch_stderr, ch_colorout, ch_colorerr
	}
}
