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
				name:    "PASS_COMPILE",
				pattern: ignoreError("(Pass|PASS|pass)", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("47")),
			},
			{
				name:    "FAIL_COMPILE",
				pattern: ignoreError("(Fail|FAIL|fail)", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("160")),
			},
			{
				name:    "DEBUG_COMPILE",
				pattern: ignoreError("(Debug|debug|DEBUG)", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("105")),
			},
			{
				name:    "INFO_COMPILE",
				pattern: ignoreError("(Info|info|INFO)", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86")),
			},
			{
				name:    "WARN_COMPILE",
				pattern: ignoreError("(Warn|warn|WARN)", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("192")),
			},
			{
				name:    "ERROR_COMPILE",
				pattern: ignoreError("(Error|error|ERROR)", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("204")),
			},
			{
				name:    "FATAL_COMPILE",
				pattern: ignoreError("(Fatal|fatal|FATAL)", regexp.Compile),
				style:   lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("134")),
			},
			{
				name:    "CRITICAL_COMPILE",
				pattern: ignoreError("(Critical|critical|CRITICAL)", regexp.Compile),
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
