package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kagundajm/gobuffalo_examples/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
