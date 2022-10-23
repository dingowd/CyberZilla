package standartpush

import (
	"fmt"
	"github.com/dingowd/CyberZilla/test3/models"
	"io"
)

type StandartPush struct {
	out io.Writer
}

func New(output io.Writer) *StandartPush {
	return &StandartPush{
		out: output,
	}
}

func (p *StandartPush) Push(user models.User) {
	fmt.Fprintln(p.out, user)
}
