package syntax

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

func Escape(input string) string {
	b := &bytes.Buffer{}
	for _, r := range input {
		escape(b, r, false)
	}
	return b.String()
}

const meta = `\.+*?()|[]{}^$#`

func escape(b *bytes.Buffer, r rune, force bool) {
	if unicode.IsPrint(r) {
		if strings.IndexRune(meta, r) >= 0 || force {
			b.WriteRune('\\')
		}
		b.WriteRune(r)
		return
	}

	switch r {
	case '\a':
		b.WriteString(`\a`)
	case '\f':
		b.WriteString(`\f`)
	case '\n':
		b.WriteString(`\n`)
	case '\r':
		b.WriteString(`\r`)
	case '\t':
		b.WriteString(`\t`)
	case '\v':
		b.WriteString(`\v`)
	default:
		if r < 0x100 {
			b.WriteString(`\x`)
			s := strconv.FormatInt(int64(r), 16)
			if len(s) == 1 {
				b.WriteRune('0')
			}
			b.WriteString(s)
			break
		}
		b.WriteString(`\x{`)
		b.WriteString(strconv.FormatInt(int64(r), 16))
		b.WriteString(`}`)
	}
}
