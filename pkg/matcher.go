package pkg

import (
	"log"
	"regexp"

	"github.com/charmbracelet/lipgloss"
)

var (
	REGEX_STYLE = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("51"))
)

type Colour struct {
	name    string
	pattern *regexp.Regexp
	style   lipgloss.Style
}
type RorColouiser struct {
	Patterns []Colour
}

func NewRorColuriser() RorColouiser {
	return RorColouiser{
		Patterns: []Colour{
			{
				name:    "PASS",
				pattern: ignoreError("pass", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("47")),
			},
			{
				name:    "FAIL",
				pattern: ignoreError("fail", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("160")),
			},
			{
				name:    "DEBUG",
				pattern: ignoreError("debug", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("105")),
			},
			{
				name:    "INFO",
				pattern: ignoreError("info", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86")),
			},
			{
				name:    "WARN",
				pattern: ignoreError("warn", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("192")),
			},
			{
				name:    "ERROR",
				pattern: ignoreError("error", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("204")),
			},
			{
				name:    "FATAL",
				pattern: ignoreError("fatal", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("134")),
			},
			{
				name:    "CRITICAL",
				pattern: ignoreError("critical", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("134")),
			},
		},
	}
}

func (ror *RorColouiser) ApplyStyle(str string, re *regexp.Regexp) string {
	if re != nil && re.MatchString(str) {
		return REGEX_STYLE.Render(str)
	}
	for _, colour := range ror.Patterns {
		if colour.pattern.MatchString(str) {
			return colour.style.Render(str)
		}
	}
	return str
}

func ignoreError[X, Y any](x X, f func(X) (Y, error)) Y {
	ret, err := f(x)
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	return ret
}
