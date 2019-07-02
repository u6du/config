package user

import (
	"path"

	"github.com/u6du/config"
)

func UserFilename(filename string) string {
	return path.Join("user", USER, filename)
}

func UserPathIsNew(filename string) (string, bool) {
	return config.FilePathIsNew(UserFilename(filename))
}

func UserByte(filename string, init func() []byte) []byte {
	return config.FileByte(UserFilename(filename), init)
}

func UserString(filename string, init func() string) string {
	return config.FileString(UserFilename(filename), init)
}

func UserLi() {
	return config.FileLi()
}
