package bubbles

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type MainMenu struct {
	items  []string
	cursor int
}

func InitMainmenu() MainMenu {
	return MainMenu{items: []string{"Start", "Score", "Quit"}, cursor: 0}
}
func (m MainMenu) Init() tea.Cmd {
	return nil
}

func (m MainMenu) View() string {
	s := "Main Menu\n\n"

	for i, menu := range m.items {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, menu)
	}

	s += "\nPress q to quit.\n"

	return s
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case "enter", " ":
			switch m.items[m.cursor] {
			case "Start":
				g := tea.NewProgram(InitGame())
				if err := g.Start(); err != nil {
					panic(err)
				}
			case "Quit":
				os.Exit(0)
			}
		}
	}

	return m, nil
}
