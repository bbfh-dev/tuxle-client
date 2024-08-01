package ui

import (
	"bufio"
	"io"
	"net"
	"strings"
	"time"

	"github.com/bbfh-dev/tuxle-client/ui/widget"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Width  int
	Height int

	Connection net.Conn

	InputArr     []string
	InputCurrent int

	Bubbles    []widget.Bubble
	shouldQuit bool

	variables map[string]string
}

func NewModel() *Model {
	return &Model{
		InputArr:  []string{""},
		variables: map[string]string{},
	}
}

func (model *Model) Init() tea.Cmd {
	go model.readMessage()
	return tick
}

func (model *Model) Update(raw tea.Msg) (tea.Model, tea.Cmd) {
	if model.shouldQuit {
		model.CloseConnection()
		return model, tea.Quit
	}

	switch msg := raw.(type) {

	case tea.WindowSizeMsg:
		model.Width, model.Height = msg.Width, msg.Height

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			model.CloseConnection()
			return model, tea.Quit
		}
		model.HandleInput(msg.String())

	case tickMsg:
		return model, tick
	}

	return model, nil
}

func (model *Model) View() string {
	if model.Height == 0 {
		return ""
	}

	header := model.Header()
	body := model.Body()
	input := model.Input()

	viewport := model.Height - len(header) - len(input) - 1
	overflow := len(body) - viewport
	if overflow > 0 {
		body = body[overflow : overflow+viewport]
	}

	content := append(header, body...)
	content = append(content, strings.Repeat("\n", max(0, -overflow)))
	content = append(content, input...)

	return strings.Join(content, "\n")
}

type tickMsg bool

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 250)
	return tickMsg(true)
}

func (model *Model) readMessage() {
	for {
		if model.Connection == nil {
			continue
		}

		data, err := bufio.NewReader(model.Connection).ReadString('\r')
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		body := strings.TrimSpace(string(data[:len(data)-1]))
		model.NewIncomingBubble(true, body)

		if strings.HasPrefix(body, "SET") {
			parts := strings.SplitN(body, "\n\n", 2)
			words := strings.SplitN(parts[0], " ", 2)
			model.variables[words[1]] = parts[1]
			if words[1] == "token" {
				err := widget.SetCredentials(parts[1])
				if err != nil {
					model.NewErrorBubble(err.Error())
				}
			}
		}
	}
}
