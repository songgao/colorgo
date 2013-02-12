package main

const (
	RED  = "\x1b[38;05;1m"
	BLUE = "\x1b[38;05;4m"

	BOLD  = "\x1b[1m"
	RESET = "\x1b[0m"
)

func sgrBoldRed(text string) string {
	return RED + BOLD + text + RESET
}

func sgrBoldBlue(text string) string {
	return BLUE + BOLD + text + RESET
}
