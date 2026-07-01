package escpaper_test

import (
	"fmt"
	"github.com/message-256/escpaper"
	"testing"
	"errors"
)

type output struct {
	final string
	err   error
}

func TestEscaper(t *testing.T) {
	fmt.Println("escaper");

	var outputs = map[string]output{

		`test\\\"text`: {final: "test\x1b\\\x1b\"text", err: nil},
		`test\\`:       {final: "test\x1b\\", err: nil},
		`\\test`:       {final: "\x1b\\test", err: nil},
		"test":         {final: "test", err: nil},
	}
	for input := range outputs {
		fmt.Println("input = ", input)
		ret, err := escpaper.Escape(input)
		if ret != outputs[input].final && err != outputs[input].err {
			t.Errorf("ret not right value")
		}
	}

}
func TestDelim(t *testing.T) {
	fmt.Println("delim");
	var outputs = map[string]output{

		`test\\\"text"`: {final: "test\x1b\\\x1b\"text", err: nil},
		`test\\"`:       {final: "test\x1b\\", err: nil},
		`\\test"`:       {final: "\x1b\\test", err: nil},
		"test\"":        {final: "test", err: nil},
		`test\\\"text",other stuff`: {final: "test\x1b\\\x1b\"text", err: nil},
		`test\\",other stuff`:       {final: "test\x1b\\", err: nil},
		`\\test",other stuff`:       {final: "\x1b\\test", err: nil},
		"test\",others stuff":        {final: "test", err: nil},
		`test`: {final:"",err:errors.New("string with no end")},
		"":{final:"",err:nil},
		"\"":{final:"",err:nil},
	}
	for input := range outputs {
		fmt.Println("input = ", input)
		ret, err := escpaper.SubString(input, '"')
		if ret != outputs[input].final && err != outputs[input].err {
			t.Errorf("ret not right value")
		}
	}

}
