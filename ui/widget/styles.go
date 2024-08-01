package widget

import "github.com/charmbracelet/lipgloss"

var InputStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), true).
	Padding(0, 1)

var IncomingStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#9C9C9C")).
	Border(lipgloss.ThickBorder(), false, false, false, true).
	BorderForeground(lipgloss.Color("#CFCFCF")).
	AlignHorizontal(lipgloss.Left).
	Padding(0, 1)

var OutcomingStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff")).
	Border(lipgloss.ThickBorder(), false, false, false, true).
	BorderForeground(lipgloss.Color("#31F038")).
	AlignHorizontal(lipgloss.Left).
	Padding(0, 1)

var ErrorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffff")).
	Border(lipgloss.ThickBorder(), false, false, false, true).
	BorderForeground(lipgloss.Color("#EC4848")).
	AlignHorizontal(lipgloss.Left).
	Padding(0, 1)

var TimeStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#434343")).
	AlignHorizontal(lipgloss.Left)

var HeaderStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#ffffff")).
	Foreground(lipgloss.Color("#000000")).
	Padding(0, 1)
