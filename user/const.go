package user

import (
	osUser "os/user"

	"github.com/u6du/config"
)

var USER string

func init() {
	USER = config.FileOneLineFunc("user", func() string {
		user, err := osUser.Current()
		if err != nil {
			return "root"
		} else {
			return user.Name
		}
	})
}
