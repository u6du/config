package user

import (
	osUser "os/user"
	"path"

	"github.com/u6du/config"
)

var USER string

func init() {
	USER = config.Path.OneLineFunc("user", func() string {
		user, err := osUser.Current()
		if err != nil {
			return "root"
		} else {
			return user.Name
		}
	})
}

var File = config.Config{Root: path.Join(config.ROOT, USER)}
