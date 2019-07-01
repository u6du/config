package config

import (
	"io/ioutil"

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
