package user

import (
	osUser "os/user"
	"path"

	"github.com/u6du/config"
)

var USER string
var File config.Config

func init() {
	USER = config.File.OneLineFunc("user", func() string {
		user, err := osUser.Current()
		if err != nil {
			return "6du"
		} else {
			return user.Name
		}
	})
	File = config.Config{Root: path.Join(config.ROOT, "user", USER)}
}
