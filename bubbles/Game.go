package bubbles

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Game struct {
	phrase       string
	counter      int
	actionPhrase string
	timeStart    time.Time
}

func InitGame() Game {
	return Game{
		phrase:       "small test to make some knowledge on bubbletea",
		actionPhrase: "",
		counter:      0,
	}
}

func (g Game) Init() tea.Cmd {
	return nil
}

func (g Game) View() string {
	txtgood := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#32814c"))

	txtfail := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#fafafa")).
		Background(lipgloss.Color("#e40c13"))

	str := ""
	for i := 0; i < g.counter; i++ {

		if g.phrase[i] == g.actionPhrase[i] {
			str += txtgood.Render(string(g.phrase[i]))
		} else {
			str += txtfail.Render(string(g.actionPhrase[i]))
		}

	}
	str = str + "|" + g.phrase[g.counter:]

	return str

}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			return g, tea.Quit
		case "backspace":
			g.counter -= 1
			g.actionPhrase = g.actionPhrase[:g.counter]
		default:
			// Start timer
			if int(g.timeStart.Unix()) <= 0 {
				g.timeStart = time.Now()
			}
			// Check basic validation
			if len(g.phrase) > g.counter {
				g.counter += 1
				g.actionPhrase += msg.String()
			} else {
				timeSince := time.Since(g.timeStart)
				// Scoreboard
				fmt.Printf("\n\nSugaisai: %.4f\n\n", timeSince.Minutes())
				return g, tea.Quit

			}
		}

	}
	return g, nil
}
