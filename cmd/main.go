package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/models"
	"github.com/somatom98/todoist/todo"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("dead %w", err)
	}
	defer f.Close()

	todoRepo := todo.NewMockRepo()

	p := tea.NewProgram(models.NewMain(todoRepo))

	if _, err := p.Run(); err != nil {
		log.Fatalf("dead %w", err)
		os.Exit(1)
	}
}
