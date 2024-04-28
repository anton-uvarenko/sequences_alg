package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/anton-uvarenko/suqences_alg/core"
	"github.com/anton-uvarenko/suqences_alg/core/styles"
	"github.com/anton-uvarenko/suqences_alg/core/tea_models"
	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

type Model struct {
	PathsModel *teamodels.PathsModel
	NodeModel  *teamodels.NodeModel
	InputModel *teamodels.InputModel
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type.String() {
		case tea.KeyCtrlC.String():
			return m, tea.Quit
		case tea.KeyEnter.String(), tea.KeyTab.String(), tea.KeyShiftTab.String(), tea.KeyUp.String(), tea.KeyDown.String():
			if !m.InputModel.Display {
				m.PathsModel.SwitchSelected(msg)
				return m, nil
			}

			_, cmds := m.InputModel.HandleInputs(msg.Type)
			if cmds == nil {
				m.createNewPath()
				m.InputModel.Display = false
				m.PathsModel.Main.Focus()

				return m, nil
			}

			return m, tea.Batch(cmds...)
		case tea.KeyCtrlA.String():
			m.InputModel.Display = true
			m.InputModel.Inputs[0].Focus()
			m.InputModel.FocusedIndex = 0
			return m, nil
		case tea.KeyRunes.String():
			switch msg.String() {
			case "r":
			}
		}

	default:
	}

	cmds := make([]tea.Cmd, 2)
	cmds = append(cmds, m.InputModel.UpdateInputs(msg))
	cmds = append(cmds, m.PathsModel.UpdateTable(msg))

	return m, tea.Batch(cmds...)
}

func (m Model) createNewPath() {
	inputsData := m.InputModel.GetInputsData()
	rawPosition := strings.Split(inputsData[0], " ")
	row, _ := strconv.Atoi(rawPosition[0])
	column, _ := strconv.Atoi(rawPosition[1])
	startNum, _ := strconv.Atoi(inputsData[1])
	endNum, _ := strconv.Atoi(inputsData[2])

	m.PathsModel.AddRow(teamodels.TableRow{
		PathId:     core.CurrentId,
		Position:   m.InputModel.Inputs[0].Value(),
		StartValue: startNum,
		EndValue:   endNum,
	})

	core.GetNodeAt(m.NodeModel.Nodes, row, column).Walk(int(math.Abs(float64(startNum)-float64(endNum))), startNum, core.CurrentId)
	core.CurrentId++
	m.InputModel.ClearInputs()
}

func (m Model) View() string {
	result := viewField(m.NodeModel.Nodes) + "\n"
	result += m.PathsModel.ViewStarts() + "\n"

	if m.InputModel.Display {
		result += viewSeqParams(m.InputModel.Inputs) + "\n"
	}

	return result
}

func viewField(nodes []*core.Node) string {
	result := ""
	for i, v := range nodes {
		if i%6 == 0 {
			result += "\n"
		}

		if len(v.IsMarked) != 0 {
			colorId := v.IsMarked[0].PathId
			if len(v.IsMarked) > 1 {
				colorId = 4
			}

			result += fmt.Sprintf("%v ", styles.PatchIdToColor[colorId].Render(fmt.Sprintf("%v", v.Value)))
		} else {
			result += fmt.Sprintf("%v ", styles.PatchIdToColor[0].Render(fmt.Sprintf("%v", v.Value)))
		}

		if v.Value/10 == 0 {
			result += " "
		}
	}

	return result
}

func viewSeqParams(inputs []textinput.Model) string {
	var b strings.Builder

	for i := range inputs {
		b.WriteString(inputs[i].View())
		if i < len(inputs)-1 {
			b.WriteRune('\n')
		}
	}

	return b.String()
}

func addSuqenceParams(model *Model) {
	var t textinput.Model

	for range 3 {
		model.InputModel.Inputs = append(model.InputModel.Inputs, textinput.Model{})
	}

	for i := range model.InputModel.Inputs {
		t = textinput.New()
		switch i {
		case 0:
			t.Placeholder = "position"
			t.Focus()
			t.PromptStyle = styles.FocusedStyle
			t.TextStyle = styles.FocusedStyle
		case 1:
			t.Placeholder = "from number"
		case 2:
			t.Placeholder = "to number"
		}

		model.InputModel.Inputs[i] = t
	}
}
