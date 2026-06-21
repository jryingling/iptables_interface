package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"iptables_interface/ui"
)

func main() {
    p := tea.NewProgram(ui.InitialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
