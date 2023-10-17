package main

import (
	tea "github.com/charmbracelet/bubbletea"
	. "github.com/nopfault/fatfingers/bubbles"
)

func main() {
	p := tea.NewProgram(InitMainmenu())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
