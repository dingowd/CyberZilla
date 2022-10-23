package pusher

import "github.com/dingowd/CyberZilla/test3/models"

type Pusher interface {
	Push(user models.User)
}
