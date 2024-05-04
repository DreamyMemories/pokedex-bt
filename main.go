package main

import (
	"fmt"
	"os"

	"github.com/DreamyMemories/pokedex-bt/mainmenu"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := mainmenu.LoadModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
