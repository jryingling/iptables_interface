package ui

import (
  "fmt"

  tea "charm.land/bubbletea/v2"
  
  //	"iptables_interface/model"
	//  "iptables_interface/render"
)

type model struct {
	choices []string
	cursor int
	selected map[int]struct{}
}

func InitialModel() model {
	return model {
		choices: []string{"Desktop", "Server", "Router", "Custom"},
		selected: make(map[int]struct{}),

	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// from bubbletea docs, looks like it works here too
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyPressMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        // The "enter" key and the space bar toggle the selected state
        // for the item that the cursor is pointing at.
        case "enter", "space":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() tea.View {
	s := "Profile Presets\n\n"

	for i, choice := range m.choices {
		//from bubbletea docs
		cursor := " " 
		if m.cursor == i {
			cursor = ">" // we see an arrow
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)

	}
	
	s += "\nPress q to quit.\n"
	return tea.NewView(s)
}

