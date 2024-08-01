package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/bbfh-dev/tuxle-client/ui/widget"
)

func (model *Model) NewOutgoingBubble(body string, args ...any) {
	model.Bubbles = append(
		model.Bubbles,
		widget.OutcomingBubble{
			Timestamp: time.Now(),
			Body:      fmt.Sprintf(body, args...),
		},
	)
}

func (model *Model) NewIncomingBubble(isMessage bool, body string, args ...any) {
	var text = fmt.Sprintf(body, args...)
	if isMessage {
		text = widget.Highlight(fmt.Sprintf(body, args...))
	}

	model.Bubbles = append(
		model.Bubbles,
		widget.IncomingBubble{
			Timestamp: time.Now(),
			Body:      text,
			IsError:   false,
		},
	)
}

func (model *Model) NewErrorBubble(body string, args ...any) {
	model.Bubbles = append(
		model.Bubbles,
		widget.IncomingBubble{
			Timestamp: time.Now(),
			Body:      fmt.Sprintf(body, args...),
			IsError:   true,
		},
	)
}

func (model *Model) Body() []string {
	var lines []string
	for _, bubble := range model.Bubbles {
		lines = append(lines, strings.Split(bubble.Render(model.Width), "\n")...)
		lines = append(lines, "")
	}
	return lines
}
