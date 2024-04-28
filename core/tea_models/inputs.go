package teamodels

import (
	"github.com/anton-uvarenko/suqences_alg/core/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputModel struct {
	Display      bool
	Inputs       []textinput.Model
	FocusedIndex int
}

func (m *InputModel) HandleInputs(msg tea.KeyType) (tea.Cmd, []tea.Cmd) {
	if msg == tea.KeyEnter && m.FocusedIndex == len(m.Inputs) {
		return tea.Quit, nil
	}

	if msg == tea.KeyUp || msg == tea.KeyShiftTab {
		m.FocusedIndex--
	}

	if msg == tea.KeyDown || msg == tea.KeyTab {
		m.FocusedIndex++
	}

	cmds := make([]tea.Cmd, len(m.Inputs))
	for i := range m.Inputs {
		if i == m.FocusedIndex {
			cmds[i] = m.Inputs[i].Focus()
			m.Inputs[i].PromptStyle = styles.FocusedStyle
			m.Inputs[i].TextStyle = styles.FocusedStyle
			continue
		}

		m.Inputs[i].Blur()
		m.Inputs[i].PromptStyle = styles.NoStyle
		m.Inputs[i].TextStyle = styles.NoStyle
	}

	return nil, cmds
}

func (m *InputModel) UpdateInputs(msg tea.Msg) tea.Cmd {
	if !m.Display {
		return nil
	}
	cmds := make([]tea.Cmd, len(m.Inputs))
	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

func (m *InputModel) GetInputsData() []string {
	result := []string{}
	for _, v := range m.Inputs {
		result = append(result, v.Value())
	}
	return result
}

func (m *InputModel) ClearInputs() {
	for i := range m.Inputs {
		m.Inputs[i].Reset()
	}
}
