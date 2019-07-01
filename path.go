package config

import (
	"os"
	osUser "os/user"
	"path"

	"github.com/u6du/ex"
)

var ROOT string
var USER string

/*

首先从sqlite读取数据
如果sqlite中没有数据
从配置文件导入数据
*/

func init() {
	ROOT = os.Getenv("_" + PROJECT + "_ROOT")

	if len(ROOT) == 0 {
		var home string
		user, err := osUser.Current()
		if err != nil {
			home, err = os.UserHomeDir()
			ex.Panic(err)
		} else {
			home = user.HomeDir
		}
		ROOT = path.Join(home, ".config", PROJECT)
	}
	os.MkdirAll(ROOT, 0700)

	USER = FileString("user.txt", func() string {
		user, err := osUser.Current()
		if err != nil {
			return "root"
		} else {
			return user.Name
		}
	})
}

func Mkdir(filename string) {
	dirname := path.Dir(filename)
	if len(dirname) > 0 {
		os.MkdirAll(path.Join(ROOT, dirname), 0700)
	}
}

func FilePathIsNew(filename string) (string, bool) {
	filepath := path.Join(ROOT, filename)
	stat, err := os.Stat(filepath)
	notExist := os.IsNotExist(err)

	if notExist {
		Mkdir(filename)
	}

	return filepath, notExist || stat.Size() == 0
}

func UserFilename(filename string) string {
	return path.Join("user", USER, filename)
}

func UserPathIsNew(filename string) (string, bool) {
	return FilePathIsNew(UserFilename(filename))
}

func FilePath(filename string) string {
	f, _ := FilePathIsNew(filename)
	return f
}

func UserPath(filename string) string {
	f, _ := UserPathIsNew(filename)
	return f
}

/*
func Li(filename string, init string) []string {
	var li []string

	bli := LiByte(filename, init)

	for i := 0; i < len(bli); i++ {
		li = append(li, string(bli[i]))
	}

	return li
}

func LiByte(filename string, init string) [][]byte {

	filepath := path.Join(ROOT, filename+".txt")

	var txt []byte
	stat, err := os.Stat(filepath)
	if os.IsNotExist(err) || stat.Size() == 0 {
		txt = []byte(init)
		ioutil.WriteFile(filepath, txt, 0600)
	} else {
		txt, err = ioutil.ReadFile(filepath)
		ex.Panic(err)
	}

	li := make([][]byte, 0)

	scanner := bufio.NewScanner(bytes.NewReader(txt))

	rewrite := false

	for scanner.Scan() {
		line := scanner.Bytes()
		ex.Panic(err)

		trim := bytes.TrimSpace(line)
		trimLen := len(trim)

		if trimLen != len(line) {
			rewrite = true
		}
		if trimLen > 0 {
			li = append(li, trim)
		} else {
			rewrite = true
		}
	}

	if rewrite {
		ioutil.WriteFile(filepath, []byte(bytes.Join(li, []byte("\n"))), 0600)
	}

	return li

}
*/
