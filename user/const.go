package user

import (
	osUser "os/user"
	"path"

	"github.com/u6du/config"
)

var USER string

func init() {
	USER = config.Global.OneLineFunc("user", func() string {
		user, err := osUser.Current()
		if err != nil {
			return "root"
		} else {
			return user.Name
		}
	})
}

var User = config.Config{Root: path.Join(config.ROOT, USER)}
