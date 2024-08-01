package widget

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var CodeRoute = lipgloss.NewStyle().Foreground(lipgloss.Color("#F73F3F"))
var CodeNormal = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
var CodeParamKey = lipgloss.NewStyle().Foreground(lipgloss.Color("#4BB0F4"))
var CodeParamValue = lipgloss.NewStyle().Foreground(lipgloss.Color("#55F44B"))

func Highlight(body string) string {
	if strings.HasPrefix(body, "/") {
		return body
	}

	var result strings.Builder
	partitions := strings.SplitN(body, "\n\n", 2)
	lines := strings.Split(partitions[0], "\n")

	headerParts := strings.SplitN(lines[0], " ", 2)
	result.WriteString(CodeRoute.Render(headerParts[0]))
	if len(headerParts) == 2 {
		result.WriteRune(' ')
		result.WriteString(CodeNormal.Render(headerParts[1]))
	}

	if len(lines) > 1 {
		result.WriteRune('\n')
		for i, line := range lines[1:] {
			parts := strings.SplitN(line, "=", 2)
			result.WriteString(CodeParamKey.Render(parts[0]))
			if len(parts) == 2 {
				result.WriteString(CodeNormal.Render("="))
				result.WriteString(CodeParamValue.Render(parts[1]))
				if i < len(lines[1:])-1 {
					result.WriteRune('\n')
				}
			}
		}
	}

	if len(partitions) == 2 {
		result.WriteString("\n\n")
		result.WriteString(CodeNormal.Render(partitions[1]))
	}

	return result.String()
}
