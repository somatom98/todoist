package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/controllers"
	"github.com/somatom98/todoist/db"
	"github.com/somatom98/todoist/models"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("dead %w", err)
	}
	defer f.Close()

	conn := db.Init()
	defer conn.Close()

	todoRepo := controllers.NewRepo(conn)
	paneSelector := controllers.NewPaneSelector()

	p := tea.NewProgram(models.NewMain(todoRepo, paneSelector))

	if _, err := p.Run(); err != nil {
		log.Fatalf("dead %w", err)
		os.Exit(1)
	}
}
