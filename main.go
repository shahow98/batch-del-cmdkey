package main

import (
	"fmt"
	"os"
	"shahow98/batch-del-cmdkey/core"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("程序奔溃了: %v", err)
		os.Exit(1)
	}
}

type model struct {
	tip      string
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	m := model{
		tip:      "",
		choices:  []string{},
		selected: make(map[int]struct{}),
	}
	m.reset()
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "d":
			m.delete()
		case "a":
			if len(m.selected) == len(m.choices) {
				m.selected = make(map[int]struct{})
			} else {
				for i := 0; i < len(m.choices); i++ {
					m.selected[i] = struct{}{}
				}
			}
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "选择需要删除的凭证?\n\n"
	if len(m.tip) > 0 {
		s += m.tip + "\n\n"
	}
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\n[enter]:选择 [↑+↓]:移动 [a]:全选 [d]:删除 [ctrl+c]:退出\n"
	return s
}

func (m *model) reset() {
	list, err := core.GetAllCmdKeys()
	if err != nil {
		m.tip = "获取凭证列表失败: " + err.Error()
	}
	m.choices = list
	m.selected = make(map[int]struct{})
	if len(m.choices) == 0 {
		m.tip = "当前凭证列表为空"
	}
}

func (m *model) delete() {
	selected := []string{}
	for i := 0; i < len(m.choices); i++ {
		if _, ok := m.selected[i]; ok {
			selected = append(selected, m.choices[i])
		}
	}
	err := core.DelCmdkeys(selected)
	if err != nil {
		m.tip = "删除凭证失败: " + err.Error()
	}
	m.reset()
}
