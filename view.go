package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)

	faint = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)

	redStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Faint(true)

	greenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#39ff14")).Faint(true)

	enumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (m model) View() string {
	s := appNameStyle.Render("NOTES APP") + "\n\n"

	if m.state == titleView {
		s += "Note title:\n\n"
		s += m.textinput.View() + "\n\n"
		s += greenStyle.Render("enter - save")
		s += faint.Render(", esc - discard")
	}

	if m.state == bodyView {
		s += "Note:\n\n"
		s += m.textarea.View() + "\n\n"
		s += greenStyle.Render("ctrl+s - save")
		s += faint.Render(", esc - discard")
	}

	if m.state == deleteView {
		s += "Note:\n\n"
		s += m.textarea.View() + "\n\n"
		s += redStyle.Render("y - delete")
		s += faint.Render(", esc - discard")
	}

	if m.state == listView {
		for i, n := range m.notes {
			prefix := ""
			if i == m.listIndex {
				prefix = ">"
			}
			shortBody := strings.ReplaceAll(n.Body, "\n", "")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}
			s += enumeratorStyle.Render(prefix) + n.Title + " | " + faint.Render(shortBody) + "\n\n"
		}
		s += greenStyle.Render("n - new")
		s += faint.Render(", q - quit")
	}
	return s
}
