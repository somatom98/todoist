package main

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/somatom98/todoist/todo"
)

type model struct {
	todoRepo todo.TodoRepo
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func newModel() *model {
	return &model{
		todoRepo: todo.NewMockRepo(),
		choices:  []string{"New"},
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m *model) Init() tea.Cmd {
	ctx := context.Background()

	todos, err := m.todoRepo.GetAll(ctx)
	if err != nil {
		panic("unable to get todo list")
	}

	for _, item := range todos {
		m.choices = append(m.choices, item.String())
	}

	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m *model) View() string {
	s := "Choices: \n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
