package ui

import (
	"fmt"

	"github.com/bbfh-dev/tuxle-client/ui/widget"
)

func (model *Model) Header() []string {
	var header string

	if model.Connection == nil {
		header = widget.HeaderStyle.Width(model.Width).Render("  Not connected")
	} else {
		header = widget.HeaderStyle.Width(model.Width).Render(fmt.Sprintf(
			"󰌘  Connected to %s://%s",
			model.Connection.RemoteAddr().Network(),
			model.Connection.RemoteAddr().String(),
		))
	}

	return []string{header, ""}
}
