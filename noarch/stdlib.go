package noarch

import (
	"strconv"
	"unicode"
)

func Atoi(a string) int {
	// TODO: It looks like atoi allow other non-digit characters. We need to
	// only pull off the digit characters before we can do the conversion.
	s := ""

	for _, c := range a {
		if !unicode.IsDigit(c) {
			break
		}

		s += string(c)
	}

	// TODO: Does it always return 0 on error?
	v, _ := strconv.Atoi(s)

	return v
}

func Strtol(a string, b string, c int) int32 {
	// TODO: This is a bad implementation
	return 65535
}

func Malloc(size uint32) []byte {
	return make([]byte, size)
}

// Free doesn't do anything since memory is managed by the Go garbage collector.
// However, I will leave it here as a placeholder for now.
func Free(anything interface{}) {
}
