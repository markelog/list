package list

import "github.com/pkg/term"

func getChar() (ascii int, keyCode int, err error) {
	bytes := make([]byte, 3)

	t, err := term.Open("/dev/tty")
	if err != nil {
		return
	}

	term.RawMode(t)

	numRead, err := t.Read(bytes)
	if err != nil {
		return
	}

	// Three character control sequence, beginning with "ESC-["
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {

		// Arrows - up, down, left, right; kinda stuff
		keyCode = int(bytes[2])
	} else if numRead == 1 {

		// Ctrl + (C | D | Z) kinda stuff
		ascii = int(bytes[0])
	}

	t.Restore()
	t.Close()

	return
}
