package mainmenu

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type CommandItem struct {
	title       string
	description string
}

func (i CommandItem) Title() string       { return i.title }
func (i CommandItem) Description() string { return i.description }
func (i CommandItem) FilterValue() string { return i.title }

type Model struct {
	List   list.Model
	Choice string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Key Press
	case tea.KeyMsg:
		// What key pressed
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			i, ok := m.List.SelectedItem().(CommandItem)
			if ok {
				m.Choice = i.title
			}
			return m, nil
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	// return lipgloss.NewStyle().Render(m.list.View())
	return m.List.View()
}

func LoadModel() Model {
	items := []list.Item{
		CommandItem{title: "Explore", description: "Explore the area to return list of pokemons"},
		CommandItem{title: "Map", description: "Shows area"},
		CommandItem{title: "Catch", description: "Catch a pokemon"},
		CommandItem{title: "Inspect", description: "Inspect a pokemon"},
		CommandItem{
			title:       "Quit",
			description: "Exit the program",
		},
	}
	model := Model{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	model.List.Title = "Pokedex"
	return model
}
