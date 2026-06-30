package escpaper

import (
	"errors"
	"strings"
)

func Escape(input string) (string, error) {
	var escaped string
	if strings.Contains(input, "\x1b") {
		return "", errors.New("i think the strings already escaped")
	}
	if len(input) <= 1 {
		return "", nil
	}
	//escapin
	for i := range input {
		if input[i] == '\\' {
			if i == 0 {
				escaped += "\x1b"
			} else if escaped[i-1] != '\x1b' {
				escaped += "\x1b"
			} else {
				escaped += "\\"
			}
		} else {
			escaped += string(input[i])
		}
	}
	return escaped, nil
}
func SubString(input string, delim rune) (string, error) {
	var next, last int
	var returned string
	if len(input) <= 1 {
		return "", errors.New("not formated string")
	} else if input[0] == byte(delim) && input[1] == byte(delim) {
		return string(delim) + string(delim), nil
	}
	for {
		next = strings.Index(input[last:], string(delim))
		if next == -1 {
			return "", errors.New("string with no end")
		}
		escapedsection, err := Escape(string(input[last : last+next]))
		if err != nil {
			return escapedsection, err
		}
		returned += escapedsection
		last += next

		if input[last-1] != '\x1b' {
			last++
			break
		} else {
			last++
		}
	}
	return returned, nil

}
