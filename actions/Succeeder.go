package actions

import (
	b3 "github.com/golang-soft/behavior3go"
	. "github.com/golang-soft/behavior3go/core"
)

type Succeeder struct {
	Action
}

func (this *Succeeder) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}
