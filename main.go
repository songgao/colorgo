package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func start(cmd *exec.Cmd, stdout chan<- string, stderr chan<- string) {
	outio, _ := cmd.StdoutPipe()
	errio, _ := cmd.StderrPipe()
	out := bufio.NewReader(outio)
	err := bufio.NewReader(errio)
	cmd.Start()
	h := func(ch chan<- string, in *bufio.Reader) {
		var line string
		var err error
		for err == nil {
			line, err = in.ReadString('\n')
			if err == nil {
				ch <- line
			}
		}
		close(ch)
	}
	go h(stdout, out)
	go h(stderr, err)
}

func findColorizer(command string) colorizer {
	c, ok := colorizers[command]
	if !ok {
		return getColorizer(nocolor)
	}
	return c
}

func main() {
	var stdout, stderr chan<- string
	var colorout, colorerr <-chan string

	if len(os.Args) == 1 {
		stdout, stderr, colorout, colorerr = findColorizer("")()
		start(exec.Command("go"), stdout, stderr)
	} else {
		stdout, stderr, colorout, colorerr = findColorizer(os.Args[1])()
		start(exec.Command("go", os.Args[1:]...), stdout, stderr)
	}

	var wg sync.WaitGroup
	h := func(ch <-chan string, out *os.File) {
		ok := true
		var line string
		for ok {
			line, ok = <-ch
			if ok {
				fmt.Fprint(out, line)
			}
		}
		wg.Done()
	}
	wg.Add(2)
	go h(colorout, os.Stdout)
	go h(colorerr, os.Stderr)
	wg.Wait()
}
