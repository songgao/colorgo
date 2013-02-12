package main

import (
	"regexp"
	"strings"
)

func nocolor(in string) string {
	return in
}

var buildRE, _ = regexp.Compile(`^(.*\.go)\:(\d*)\:(.*)\n?$`)

func build(in string) (out string) {
	matches := buildRE.FindStringSubmatch(in)
	if len(matches) > 0 {
		out = sgrBoldBlue(matches[1]) + ":" + sgrBoldRed(matches[2]) + ":" + matches[3] + "\n"
	} else {
		if strings.HasSuffix(in, "\n") {
			out = in
		} else {
			out = in + "\n"
		}
	}
	return
}
