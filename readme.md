# List [![Build Status](https://travis-ci.org/markelog/list.svg)](https://travis-ci.org/markelog/list) [![GoDoc](https://godoc.org/github.com/markelog/list?status.svg)](https://godoc.org/github.com/markelog/list) [![Go Report Card](https://goreportcard.com/badge/github.com/markelog/list)](https://goreportcard.com/report/github.com/markelog/list)

> Terminal interactive list

![](./example.gif)

```go
func GetAnimal() string {
	options := []string{"Gangsta panda", "Sexy turtle", "Killa gorilla",}

	l := list.New("Which animal is the coolest?", options)
	l.Show()

	// Waiting for the user input
	result := l.Get()

	return result // "Gangsta panda"
}
```

