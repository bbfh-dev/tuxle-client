package widget

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func withTimestamp(timestamp time.Time) string {
	diff := time.Now().Sub(timestamp).Round(time.Second)
	return TimeStyle.Render(fmt.Sprintf(" %s ago", diff))
}

type Bubble interface {
	Render(int) string
}

type OutcomingBubble struct {
	Timestamp time.Time
	Body      string
}

func (bubble OutcomingBubble) Render(width int) string {
	return lipgloss.JoinVertical(
		0,
		withTimestamp(bubble.Timestamp),
		IncomingStyle.Width(int(float64(width)/1.5)).
			Render(bubble.Body),
	)
}

type IncomingBubble struct {
	Timestamp time.Time
	Body      string
	IsError   bool
}

func (bubble IncomingBubble) Render(width int) string {
	if bubble.IsError {
		return lipgloss.JoinVertical(
			0,
			withTimestamp(bubble.Timestamp),
			ErrorStyle.Width(int(float64(width)/1.5)).
				Render(bubble.Body),
		)
	}

	return lipgloss.JoinVertical(
		0,
		withTimestamp(bubble.Timestamp),
		OutcomingStyle.Width(int(float64(width)/1.5)).
			Render(bubble.Body),
	)
}
