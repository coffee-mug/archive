package notes

import (
	"testing"
)

func TestNotes(t *testing.T) {
	t.Run("Can input a character", func(t *testing.T) {
		archive := NewArchive()

		letter := "a"

		archive.Type(letter)

		got := archive.buffer

		if got.String() != letter {
			t.Errorf("Got %s Expect %s", got.String(), letter)
		}
	})
}
