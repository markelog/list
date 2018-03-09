// Package list for interactive CLI list
package list

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/markelog/curse"
)

// Printer function type used for outputing
type Printer func(...interface{}) (int, error)

// Colors settings
type Colors struct {
	Head,
	Option,
	Highlight *color.Color
}

// DefaultColors which will be used by default
var DefaultColors = &Colors{
	Head:      color.New(color.FgWhite, color.Bold),
	Option:    color.New(color.FgWhite),
	Highlight: color.New(color.FgCyan),
}

const (
	// HideCursor ASCII sequence to hide cursor
	HideCursor = "\033[?25l"

	// ShowCursor ASCII sequence for show cursor
	ShowCursor = "\033[?25h"

	keyCtrlC = 3
	keyCtrlD = 4
	keyCtrlZ = 26
	keyEnter = 13
	keyUp    = 65
	keyDown  = 66
)

// List base struct
type List struct {
	Index   int
	Print   Printer
	Cursor  *curse.Cursor
	colors  *Colors
	name    string
	chooser string
	indent  string
	options []string
}

// GetWith will call .New(...); .Show() and returns .Get()
func GetWith(name string, options []string) string {
	l := New(name, options)
	l.Show()
	return l.Get()
}

// New returns a list initialized with Default theme.
func New(name string, options []string) *List {
	list := &List{}

	list.name = name
	list.options = options
	list.Index = 0

	list.Cursor = &curse.Cursor{}

	list.SetColors(DefaultColors)
	list.SetPrint(fmt.Print)
	list.SetChooser(" ❯ ")
	list.SetIndent(3)

	return list
}

// SetChooser sets chooser string i.e. " ❯ "
func (list *List) SetChooser(chooser string) {
	list.chooser = chooser
}

// SetIndent sets indent before options
func (list *List) SetIndent(indent int) {
	list.indent = "" + strings.Repeat(" ", indent)
}

// SetColors sets colors for the output
func (list *List) SetColors(colors *Colors) {
	list.colors = colors
}

// SetPrint set print function
func (list *List) SetPrint(print Printer) {
	list.Print = print
}

// PrintHighlight prints highlighted list element
func (list *List) PrintHighlight(element string) int {
	newIndent := len(list.indent) - 3
	indent := ""

	if newIndent > 1 {
		indent = strings.Repeat(" ", newIndent)
	}

	list.colors.Highlight.Set()
	bytes, _ := list.Print(indent + list.chooser + element)
	color.Unset()

	return bytes
}

// PrintOption prints list option
func (list *List) PrintOption(option string) int {
	list.colors.Option.Set()
	bytes, _ := list.Print(list.indent + option)
	color.Unset()

	return bytes
}

// PrintHead prints list header
func (list *List) PrintHead() int {
	list.colors.Head.Set()
	bytes, _ := list.Print(list.name)
	color.Unset()

	return bytes
}

// PrintResult prints list header and choosen option
func (list *List) PrintResult(result string) int {
	var bytes int

	bytes = list.PrintHead()

	list.colors.Highlight.Set()
	printBytes, _ := list.Print(" ", result)
	color.Unset()

	return bytes + printBytes + list.Println()
}

// Println just prints "\n"
func (list *List) Println() int {
	result, _ := list.Print("\n")

	return result
}

// Show list
func (list *List) Show() int {
	list.Print(HideCursor)

	result := list.PrintHead() + list.Println()

	return result + list.ShowOptions()
}

// ShowOptions shows options
func (list *List) ShowOptions() int {
	result := 0

	for Index, element := range list.options {
		if Index == 0 {
			result += list.PrintHighlight(element) + list.Println()
			continue
		}

		result += list.PrintOption(element) + list.Println()
	}

	list.Index = 1
	list.Cursor.MoveUp(len(list.options))

	return result
}

// Get result
func (list *List) Get() string {

	// Listen for the user input
	for {
		ascii, keyCode, err := getChar()
		if err != nil {
			list.Print(err)
			os.Exit(1)
		}

		switch keyCode {
		case keyUp:
			list.HighlightUp()
		case keyDown:
			list.HighlightDown()
		}

		switch ascii {
		case keyCtrlC:
			list.Exit()
			return ""
		case keyCtrlZ:
			list.Exit()
			return ""
		case keyCtrlD:
			list.Exit()
			return ""
		case keyEnter:
			return list.Enter()
		}
	}
}

// ClearOptions clears options from console
func (list *List) ClearOptions() {
	length := len(list.options)
	diff := length - list.Index

	// Move to the last line.
	if diff != 0 {
		list.Cursor.MoveDown(diff)
	}

	list.Index = length

	// Erase options
	for list.Index != 0 {
		list.Index--
		list.Cursor.EraseCurrentLine()
		list.Cursor.MoveUp(1)
	}
}

// Enter key handler
func (list *List) Enter() string {
	result := list.options[list.Index-1]

	list.ClearOptions()
	list.PrintResult(result)
	list.Print(ShowCursor)

	return result
}

// Exit Ctrl + (C | D | Z) and alike handler
func (list *List) Exit() {

	// Should go down to the last option
	for list.Index != len(list.options) {
		list.Index++
		list.Cursor.MoveDown(1)
	}

	list.Println()
	list.Println()
	list.Print(ShowCursor)

	os.Exit(1)
}

// HighlightUp highlights option above
func (list *List) HighlightUp() *List {

	// If there is no where to go
	if 0 >= list.Index-1 {
		return list
	}

	list.Cursor.EraseCurrentLine()
	list.PrintOption(list.options[list.Index-1])

	list.Cursor.MoveUp(1)

	list.Cursor.EraseCurrentLine()
	list.Index--
	list.PrintHighlight(list.options[list.Index-1])

	return list
}

// HighlightDown highlights option below
func (list *List) HighlightDown() *List {

	// If there is no where to go
	if len(list.options) < list.Index+1 {
		return list
	}

	list.Cursor.EraseCurrentLine()
	list.PrintOption(list.options[list.Index-1])

	list.Cursor.MoveDown(1)

	list.Cursor.EraseCurrentLine()
	list.PrintHighlight(list.options[list.Index])

	list.Index++

	return list
}
