# List [![Build Status](https://travis-ci.org/markelog/list.svg)](https://travis-ci.org/markelog/list) [![GoDoc](https://godoc.org/github.com/markelog/list?status.svg)](https://godoc.org/github.com/markelog/list) [![Go Report Card](https://goreportcard.com/badge/github.com/markelog/list)](https://goreportcard.com/report/github.com/markelog/list)

> Terminal interactive list

This

```go
options := []string{"Gangsta panda", "Sexy turtle", "Killa gorilla",}

// returns user choice i.e. "Gangsta panda"
list.GetWith("Which animal is the coolest?", options)
```

Will get you

![](./example.gif)

## A bit of flexibility

```go
options := []string{"Gangsta panda", "Sexy turtle", "Killa gorilla",}

l := list.New("Which animal is the coolest?", options)

// Set your own printer
l.SetPrint(func(args ...interface{}) (n int, err error) {
	return fmt.Print(args...)
})

// Set your own colors
// With github.com/fatih/color package
l.SetColors(&list.Colors{
	Head:      color.New(color.BgRed),
	Option:    color.New(color.FgGreen, color.Underline),
	Highlight: color.New(color.FgCyan, color.Bold),
})

// Show the list
l.Show()

// Waiting for the user input
l.Get()
```

