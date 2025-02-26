package composites

import (
	b3 "github.com/golang-soft/behavior3go"
	. "github.com/golang-soft/behavior3go/core"
)

type MemSequence struct {
	Composite
}

/**
 * Open method.
 * @method open
 * @param {b3.Tick} tick A tick instance.
**/
func (this *MemSequence) OnOpen(tick *Tick) {
	tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), this.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *MemSequence) OnTick(tick *Tick) b3.Status {
	var child = tick.Blackboard.GetInt("runningChild", tick.GetTree().GetID(), this.GetID())
	for i := child; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)

		if status != b3.SUCCESS {
			if status == b3.RUNNING {
				tick.Blackboard.Set("runningChild", i, tick.GetTree().GetID(), this.GetID())
			}

			return status
		}
	}
	return b3.SUCCESS
}
