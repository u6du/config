package config

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/u6du/ex"
)

func FilePathIsNew(filename string) (string, bool) {
	filepath := path.Join(ROOT, filename)
	stat, err := os.Stat(filepath)
	notExist := os.IsNotExist(err)

	if notExist {
		Mkdir(filename)
	}

	return filepath, notExist || stat.Size() == 0
}

func FilePath(filename string) string {
	f, _ := FilePathIsNew(filename)
	return f
}

func FileByte(filename string, init func() []byte) []byte {
	filepath, isNew := FilePathIsNew(filename)
	var txt []byte
	if isNew {
		txt = init()
		ioutil.WriteFile(filepath, txt, 0600)
	} else {
		var err error
		txt, err = ioutil.ReadFile(filepath)
		ex.Panic(err)
	}
	return txt
}

func FileOneLineFunc(filename string, init func() string) string {
	filepath, isNew := FilePathIsNew(filename + ".1L")
	var txt string
	if isNew {
		txt = init()
		ioutil.WriteFile(filepath, []byte(txt), 0600)
	} else {
		var err error
		b, err := ioutil.ReadFile(filepath)
		ex.Panic(err)
		txt = string(bytes.TrimSpace(b))
	}
	return txt
}

func FileOneLine(filename string, init string) string {
	FileOneLineFunc(filename, func() string { return init })
}

func FileLi(filename string, init []string) []string {

	filename += ".li"

	var li []string

	filepath, isNew := FilePathIsNew(filename)
	if !isNew {
		file, err := os.Open(filepath)
		ex.Panic(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			txt := strings.TrimSpace(scanner.Text())
			if len(txt) > 0 {
				li = append(li, txt)
			}
		}
		ex.Panic(scanner.Err())
	}

	if isNew || len(li) == 0 {
		li = init
		ioutil.WriteFile(filepath, []byte(strings.Join(li, "\n")), 0600)
	}

	return li
}

func FileString(filename string, init func() string) string {
	return string(
		FileByte(
			filename,
			func() []byte {
				return []byte(init())
			}))
}
