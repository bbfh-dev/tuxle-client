package widget

import (
	"os"
	"path/filepath"

	"github.com/emersion/go-appdir"
)

var Dir = appdir.New("tuxle").UserCache()

func GetCredentials() (string, error) {
	err := os.MkdirAll(Dir, 0755)
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(filepath.Join(Dir, "token.txt"))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func SetCredentials(token string) error {
	err := os.MkdirAll(Dir, 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(Dir, "token.txt"), []byte(token), 0755)
}
