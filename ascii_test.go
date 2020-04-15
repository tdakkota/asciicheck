package asciicheck

import (
	"fmt"
	"testing"
	"unicode"
)

func TestIsASCII(t *testing.T) {
	t.Run("ascii", func(t *testing.T) {
		s := generateString()
		ch, ok := isASCII(s)

		if !ok {
			t.Error("expected that string contains only ASCII symbols")
		}

		if ch != 0 {
			t.Error("expected that ch is zero")
		}
	})

	t.Run("non-ascii", func(t *testing.T) {
		s := "привет!"
		ch, ok := isASCII(s)

		if ok {
			t.Error("expected that string contains non-ASCII symbols")
		}

		fmt.Println()
		if ch != 'п' {
			t.Error("expected that ch is equal to first letter")
		}
	})
}

func generateString() string {
	s := make([]byte, unicode.MaxASCII+1)
	for i := range s {
		s[i] = byte(i)
	}
	return string(s)
}
