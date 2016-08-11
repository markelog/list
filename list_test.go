package list_test

import (
	"fmt"

	"github.com/bouk/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/markelog/list"
)

var _ = Describe("list", func() {
	var l *list.List
	var prints []string

	BeforeEach(func() {
		prints = []string{}

		l = list.New("test", []string{"one", "two"})

		// That's what curse uses.
		// We have do it, otherwise their output will interfere with ours
		monkey.Patch(fmt.Printf, func(str string, a ...interface{}) (n int, err error) {
			return 0, nil
		})

		l.SetPrint(func(args ...interface{}) (int, error) {
			for _, element := range args {
				prints = append(prints, element.(string))
			}

			return 0, nil
		})
	})

	AfterEach(func() {
		prints = []string{}
	})

	Describe("Show", func() {
		It("should show elements", func() {
			l.Show()
			Expect(prints[0]).To(Equal("\033[?25l"))
			Expect(prints[1]).To(Equal("test"))
			Expect(prints[2]).To(Equal("\n"))
			Expect(prints[3]).To(Equal(" ❯ "))
			Expect(prints[4]).To(Equal("one"))
			Expect(prints[5]).To(Equal("\n"))
			Expect(prints[6]).To(Equal("   "))
			Expect(prints[7]).To(Equal("two"))
		})
	})

	Describe("ClearOptions", func() {
		It("should clear options", func() {
			l.Show()

			Expect(l.Index).To(Equal(1))
			Expect(l.Cursor.Position.X).To(Equal(0))
			Expect(l.Cursor.Position.Y).To(Equal(-2))

			l.ClearOptions()
			Expect(l.Index).To(Equal(0))
			Expect(l.Cursor.Position.X).To(Equal(1))
			Expect(l.Cursor.Position.Y).To(Equal(-3))
		})
	})

	Describe("Enter", func() {
		It("should get result after enter", func() {
			l.Show()
			result := l.Enter()

			Expect(result).To(Equal("one"))
			Expect(l.Index).To(Equal(0))
			Expect(l.Cursor.Position.X).To(Equal(1))
			Expect(l.Cursor.Position.Y).To(Equal(-3))
		})
	})

	Describe("HighlightDown", func() {
		BeforeEach(func() {
			l.Show()
			l.HighlightDown()
		})

		It("should go down", func() {
			Expect(l.Index).To(Equal(2))
			Expect(l.Cursor.Position.X).To(Equal(1))
			Expect(l.Cursor.Position.Y).To(Equal(-1))
		})

		It("should not go down again", func() {
			l.HighlightDown()

			Expect(l.Index).To(Equal(2))
			Expect(l.Cursor.Position.X).To(Equal(1))
			Expect(l.Cursor.Position.Y).To(Equal(-1))
		})
	})

	Describe("HighlightUp", func() {
		BeforeEach(func() {
			l.Show()
			l.HighlightDown().HighlightUp()
		})

		It("should go up", func() {
			Expect(l.Index).To(Equal(1))
			Expect(l.Cursor.Position.X).To(Equal(1))
			Expect(l.Cursor.Position.Y).To(Equal(-2))
		})

		It("should not go up again", func() {
			l.HighlightUp()

			Expect(l.Index).To(Equal(1))
			Expect(l.Cursor.Position.X).To(Equal(1))
			Expect(l.Cursor.Position.Y).To(Equal(-2))
		})
	})

	Describe("Exit", func() {
		BeforeEach(func() {
			l.Exit()
		})

		It("should print exiting chars", func() {
			Expect(prints[0]).To(Equal("\n"))
			Expect(prints[1]).To(Equal("\n"))

			// Show cursor
			Expect(prints[2]).To(Equal("\033[?25h"))
		})

		It("should point to the last", func() {
			Expect(l.Index).To(Equal(2))
		})

		It("should set right cursor position", func() {
			Expect(l.Cursor.Position.X).To(Equal(0))
			Expect(l.Cursor.Position.Y).To(Equal(2))
		})
	})
})
