package ui

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestUI(t *testing.T) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	var default_dir = fmt.Sprintf("%s/Music", dirname)
	StartUI(default_dir, "", true)
}
