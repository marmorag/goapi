package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/marmorag/goapi/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
