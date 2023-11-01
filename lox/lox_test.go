package lox_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/Krawabbel/go-lox/lox"
)

func TestREPL(t *testing.T) {

	var w = new(test_writer)
	lox.STDOUT = w
	lox.STDERR = w

	var args = []struct {
		given string
		want  string
	}{
		{"1+2", "3"},
		{"1<2", "true"},
		{"1<=2", "true"},
		{"1>2", "false"},
		{"1==nil", "false"},
		{"nil==nil", "true"},
		{"!(5 - 4 > 3 * 2 == !nil)", "true"},
	}

	var src = ""
	var want = "> "
	for _, arg := range args {
		src += arg.given + "\n"
		want += arg.want + "\n> "
	}

	lox.STDIN = strings.NewReader(src)

	if err := lox.RunREPL(); err != nil {
		t.Fatal(err)
	}

	var have = string(w.data)

	if have != want {
		t.Fatalf("run repl(): have = %s, want = %s", have, want)
	}
}

func TestLox(t *testing.T) {

	err := filepath.Walk("..",

		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			is_lox_file, err := regexp.Match("\\.lox$", []byte(info.Name()))

			if err != nil {
				t.Error(err)
			}

			if is_lox_file {

				if err := test_script_helper(path, info); err != nil {
					t.Error(err)
				}

			}

			return nil
		})

	if err != nil {
		t.Fatal(err)
	}

}

func test_script_helper(path string, info os.FileInfo) error {

	var w = new(test_writer)
	lox.STDOUT = w
	lox.STDERR = w

	lox.RunScript(path)

	var got = w.data

	var name = info.Name()
	var base_name = "." + name[:len(name)-4] + "_base"

	var base_path = strings.Replace(path, info.Name(), base_name, -1)

	var want, err = os.ReadFile(base_path)

	if err != nil {
		fmt.Printf("test \"%s\": creating new baseline: \"%s\"\n", path, base_path)
		return os.WriteFile(base_path, got, 0700)
	}

	if !reflect.DeepEqual(got, want) {
		return fmt.Errorf("error running lox \"%s\": unexpected result\n%s", path, string(got))
	}

	fmt.Printf("lox \"%s\" successfully executed\n", path)

	return nil
}

type test_writer struct {
	data []byte
}

func (w *test_writer) Write(p []byte) (n int, err error) {
	if w.data == nil {
		w.data = make([]byte, 0)
	}
	w.data = append(w.data, p...)
	return len(p), nil
}
