package teamodels

import (
	"fmt"
	"strconv"

	"github.com/anton-uvarenko/suqences_alg/core"
	"github.com/anton-uvarenko/suqences_alg/core/styles"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PathsModel struct {
	Main          table.Model
	SelectedIndex int
	Starts        []*core.Node
}

type TableRow struct {
	PathId     int
	Position   string
	StartValue int
	EndValue   int
}

func (m *PathsModel) InitModel() {
	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Position", Width: 10},
		{Title: "StartValie", Width: 10},
		{Title: "EndValue", Width: 10},
	}
	m.Main = table.New(
		table.WithColumns(columns),
		table.WithHeight(0),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	m.Main.SetStyles(s)
}

func (m *PathsModel) SwitchSelected(msg tea.KeyMsg) {
	switch msg.Type {
	case tea.KeyTab:
		m.SelectedIndex++
	case tea.KeyShiftTab:
		m.SelectedIndex--
	}

	m.Main.SetCursor(m.SelectedIndex)
}

func (m *PathsModel) AddRow(r TableRow) {
	current := m.Main.Rows()
	current = append(current, table.Row{
		fmt.Sprint(r.PathId),
		r.Position,
		fmt.Sprint(r.StartValue),
		fmt.Sprint(r.EndValue),
	})
	m.Main.SetRows(current)
	m.Main.SetHeight(m.Main.Height() + 1)
}

func (m *PathsModel) UpdateTable(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.Main, cmd = m.Main.Update(msg)
	return cmd
}

func (m *PathsModel) RemoveRow(id int) {
	current := m.Main.Rows()
	index := findIndexById(current, id)
	current = append(current[:index], current[index+1:]...)
	m.Main.SetRows(current)
}

func findIndexById(rows []table.Row, id int) int {
	for i, v := range rows {
		rowID, _ := strconv.Atoi(v[0])
		if rowID == id {
			return i
		}
	}

	return -1
}

func (m *PathsModel) ViewStarts() string {
	return styles.TableStyle.Render(m.Main.View())
}
