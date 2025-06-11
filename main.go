package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/charmbracelet/lipgloss"
)

var (
	REGEX_STYLE    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("51"))
	DEBUG_STYLE    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
	INFO_STYLE     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86"))
	WARN_STYLE     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("192"))
	ERROR_STYLE    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("204"))
	FATAL_STYLE    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("134"))
	CRITICAL_STYLE = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("134"))

	DEBUG_COMPILE, _    = regexp.Compile("(Debug|debug|DEBUG)")
	INFO_COMPILE, _     = regexp.Compile("(Info|info|INFO)")
	WARN_COMPILE, _     = regexp.Compile("(Warn|warn|WARN)")
	ERROR_COMPILE, _    = regexp.Compile("(Error|error|ERROR)")
	FATAL_COMPILE, _    = regexp.Compile("(Fatal|fatal|FATAL)")
	CRITICAL_COMPILE, _ = regexp.Compile("(Critical|critical|CRITICAL)")
)

func match(str string, re *regexp.Regexp) string {

	if DEBUG_COMPILE.MatchString(str) {
		return DEBUG_STYLE.Render(str)
	}
	if INFO_COMPILE.MatchString(str) {
		return INFO_STYLE.Render(str)
	}
	if WARN_COMPILE.MatchString(str) {
		return WARN_STYLE.Render(str)
	}
	if ERROR_COMPILE.MatchString(str) {
		return ERROR_STYLE.Render(str)
	}
	if FATAL_COMPILE.MatchString(str) {
		return FATAL_STYLE.Render(str)
	}
	if CRITICAL_COMPILE.MatchString(str) {
		return CRITICAL_STYLE.Render(str)
	}
	if re != nil && re.MatchString(str) {
		return REGEX_STYLE.Render(str)
	}
	return str
}

func flags() *regexp.Regexp {
	var pattern string
	flag.StringVar(&pattern, "pattern", "", "regex pattern to search for")
	flag.Parse()
	if pattern != "" {
		return regexp.MustCompile(pattern)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pattern := flags()
	for scanner.Scan() {
		fmt.Println(match(scanner.Text(), pattern))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
