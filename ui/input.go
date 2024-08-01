package ui

import (
	"fmt"
	"strings"

	"github.com/bbfh-dev/tuxle-client/ui/widget"
)

const INPUT_CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 -_=+!@#$%^&*()[]{};:'\",./<>?`~"

func (model *Model) Input() []string {
	input := model.InputArr[model.InputCurrent]
	input = widget.Highlight(input)

	var lines []string
	for _, line := range strings.Split(input, "\n") {
		lines = append(lines, fmt.Sprintf("→ %s", line))
	}
	lines[len(lines)-1] += "󰗧"

	data := widget.InputStyle.Width(model.Width - 2).Render(strings.Join(lines, "\n"))
	return strings.Split(data, "\n")
}

func (model *Model) Write(str string) {
	model.InputArr[model.InputCurrent] += str
}

func (model *Model) Remove(length int) {
	slice := len(model.InputArr[model.InputCurrent]) - length
	model.InputArr[model.InputCurrent] = model.InputArr[model.InputCurrent][:max(0, slice)]
}

func (model *Model) Send(command string) {
	if model.Connection == nil {
		return
	}

	if token, ok := model.variables["token"]; ok {
		lines := strings.Split(command, "\n")
		lines = append([]string{lines[0]}, "Token="+token)
		strings.Join(append(lines, lines[1:]...), "\n")
	}

	_, err := model.Connection.Write([]byte(command + "\r"))
	if err != nil {
		model.NewErrorBubble(err.Error())
	}
}

func (model *Model) Perform(command string, args ...string) error {
	var err error

	switch command {
	case "connect", "conn", "c":
		if len(args) == 0 {
			err = model.NewConnection(DEFAULT_ADDR)
		} else {
			err = model.NewConnection(args[0])
		}
	case "disconnect", "dc":
		model.CloseConnection()
	case "reconnect", "rc":
		if model.Connection == nil {
			break
		}
		addr := model.Connection.RemoteAddr().String()
		model.CloseConnection()
		model.NewConnection(addr)
	case "authenticate", "auth":
		token, err := widget.GetCredentials()
		if err != nil {
			model.NewErrorBubble(err.Error())
		} else {
			model.variables["token"] = token
		}
	case "login":
		if len(args) != 2 {
			model.NewErrorBubble("/login [user_id] [password]")
			break
		}
		model.Send(fmt.Sprintf("ACCOUNT login\nId=%s\n\n%s", args[0], args[1]))
	case "quit", "q":
		model.shouldQuit = true
	default:
		err = fmt.Errorf("Unknown command: %q", command)
	}

	return err
}

func (model *Model) HandleInput(key string) {
	if strings.Contains(INPUT_CHARACTERS, key) {
		model.Write(key)
	}

	switch key {
	case "tab":
		model.Write("\n")
	case "backspace":
		model.Remove(1)
	case "delete":
		model.Remove(100000000)
	case "up":
		model.InputCurrent -= 1
		if model.InputCurrent < 0 {
			model.InputCurrent = 0
		}
	case "down":
		model.InputCurrent += 1
		if model.InputCurrent >= len(model.InputArr) {
			model.InputCurrent = len(model.InputArr) - 1
		}
	case "enter":
		command := strings.TrimSpace(model.InputArr[model.InputCurrent])
		model.NewOutgoingBubble(command)
		switch strings.HasPrefix(command, "/") {
		case true:
			var err error
			parts := strings.Split(command[1:], " ")
			if len(parts) == 1 {
				err = model.Perform(parts[0])
			} else {
				err = model.Perform(parts[0], parts[1:]...)
			}
			if err != nil {
				model.NewErrorBubble(err.Error())
			}
		case false:
			model.Send(command)
		}
		model.InputArr = append(model.InputArr, "")
		model.InputCurrent = len(model.InputArr) - 1
	}
}
