package styles

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// ANSI codes reference- https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797

	// TextStyle is the style for the text
	TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("254"))
	// SucessStyle is the style for the success message
	SucessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	// SpinnerStyle is the style for the spinner
	SpinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	// HelpStyle is the style for the help message
	HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	// ErrorStyle is the style for the error message
	ErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))

	// WarnStyle is the style for the warning message
	WarnStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("226"))

	// TitleStyle is the style for the title
	TitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	// OptionStyle is the style for the options
	OptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	// SelectedOptionStyle is the style for the selected option
	SelectedOptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Bold(true)
	// CursorOptionStyle is the style for the cursor
	CursorOptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Italic(true)

	// HighlightStyle is the style for the highlight
	HighlightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Bold(true)
	// HintStyle is the style for the hint
	HintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))

	// DocStyle styling for viewports
	DocStyle = lipgloss.NewStyle().Margin(0, 2)

	// BaseStyle is the base style for the table
	BaseStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
)

var (
	// WindowSize store the size of the terminal window
	WindowSize tea.WindowSizeMsg
)
