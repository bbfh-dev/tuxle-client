package main

import (
	"log"

	"github.com/bbfh-dev/tuxle-client/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}
