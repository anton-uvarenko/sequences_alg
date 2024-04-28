package styles

import "github.com/charmbracelet/lipgloss"

var PatchIdToColor = map[int]lipgloss.Style{
	0: lipgloss.NewStyle().Foreground(lipgloss.Color("#fcfcfc")),
	2: lipgloss.NewStyle().Foreground(lipgloss.Color("#60b347")),
	3: lipgloss.NewStyle().Foreground(lipgloss.Color("#c23ac0")),
	1: lipgloss.NewStyle().Foreground(lipgloss.Color("#c72020")),
	4: lipgloss.NewStyle().Foreground(lipgloss.Color("#f5e10c")),
}

var (
	NoStyle      = lipgloss.NewStyle()
	FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	TableStyle   = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
)
