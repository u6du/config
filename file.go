package config

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/u6du/ex"
)

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

func UserByte(filename string, init func() []byte) []byte {
	return FileByte(UserFilename(filename), init)
}

func UserString(filename string, init func() string) string {
	return FileString(UserFilename(filename), init)
}
