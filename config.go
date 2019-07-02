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

type Config struct {
	Root string
}

func (c *Config) Mkdir(filename string) {
	dirname := path.Dir(filename)
	if len(dirname) > 0 {
		os.MkdirAll(path.Join(c.Root, dirname), 0700)
	}
}

func (c *Config) PathIsNew(filename string) (string, bool) {
	filepath := path.Join(c.Root, filename)
	stat, err := os.Stat(filepath)
	notExist := os.IsNotExist(err)

	if notExist {
		c.Mkdir(filename)
	}

	return filepath, notExist || stat.Size() == 0
}

func (c *Config) Path(filename string) string {
	f, _ := c.PathIsNew(filename)
	return f
}

func (c *Config) Byte(filename string, init func() []byte) []byte {
	filepath, isNew := c.PathIsNew(filename)
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

func (c *Config) OneLineFunc(filename string, init func() string) string {
	filepath, isNew := c.PathIsNew(filename + ".1L")
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

func (c *Config) OneLine(filename string, init string) string {
	return c.OneLineFunc(filename, func() string { return init })
}

func (c *Config) Li(filename string, init []string) []string {

	filename += ".li"

	var li []string

	filepath, isNew := c.PathIsNew(filename)
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

func (c *Config) String(filename string, init func() string) string {
	return string(
		c.Byte(
			filename,
			func() []byte {
				return []byte(init())
			}))
}
