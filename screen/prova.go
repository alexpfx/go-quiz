package screen

import (
	"fmt"
	"github.com/alexpfx/go-quiz/prova"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"strings"
)

func NewScreen(engine prova.Engine) *tea.Program {
	return tea.NewProgram(initmodel(engine))
}

var (
	verticalDistance = 1
	appStyle         = lipgloss.NewStyle().Padding(1, 2)
	titleStyle       = lipgloss.NewStyle().Bold(true).
				Foreground(lipgloss.Color("115")).PaddingBottom(verticalDistance)
	enunciadoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("011")).PaddingBottom(verticalDistance + 1)
	questaoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("129")).PaddingBottom(verticalDistance)
)
var multipleChoiceArray = []string{"A", "B", "C", "D", "E"}
var selMap = make(map[string]int)

func init() {
	for i, l := range multipleChoiceArray {
		selMap[l] = i
	}
}

type model struct {
	titulo    string
	enunciado string
	choices   []mq
	cursor    int
	chosen    int
	showAnswer bool
	engine    prova.Engine
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit

		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			m.chosen = m.cursor
		case "enter":
			m.showAnswer = true
			return m, nil
		case "esc":
			m.chosen = -1
		default:
			if i, ok := selMap[strings.ToUpper(msg.String())]; ok {
				m.chosen = i
			}
		}

	}

	return m, nil
}

func (m model) View() string {
	t := titleStyle.Render(m.titulo)
	e := enunciadoStyle.Render(m.enunciado)
	q := questaoStyle.Render(m.questao())

	return appStyle.Render(lipgloss.JoinVertical(lipgloss.Top, t, e, q))
}

func (m model) questao() string {
	s := ""
	for i, q := range m.choices {
		alt := multipleChoiceArray[i]
		cursor := " "
		choice := " "
		if m.cursor == i {
			cursor = ">"
		}
		if m.chosen == i {
			choice = "x"
		}
		if m.showAnswer {
			if q.correct {
				alt = "V"
			}
		}

		s += fmt.Sprintf("%s) %s [%s] %s\n", alt, cursor, choice, q.text)
	}
	return s
}

type mq struct {
	correct bool
	text  string
}

func initmodel(engine prova.Engine) model {

	return toModel(engine)
}

func toModel(e prova.Engine) model {
	q := e.Get()
	choices := make([]mq, 0)
	for _, a := range q.Alternativas {
		choices = append(choices, mq{
			correct: a == q.Correta,
			text:  a.Texto,
		})
	}
	log.Println(choices)
	return model{
		titulo:    "Quiz",
		enunciado: q.Enunciado,
		choices:   choices,
		chosen:    -1,
		cursor:    2,
		engine:    e,
	}
}
