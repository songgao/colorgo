package main

import (
	"regexp"
	"strings"
)

var buildRE, _ = regexp.Compile(`^(.*\.go)\:(\d*)\:(.*)\n?$`)

func buildOrTest(in string) (out string) {
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
