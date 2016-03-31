// automatically generated, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Stat struct {
	_tab flatbuffers.Table
}

func (rcv *Stat) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Stat) ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Stat) User() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) Nice() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) System() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) Idle() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) IOWait() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) IRQ() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) SoftIRQ() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) Steal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) Quest() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) QuestNice() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func StatStart(builder *flatbuffers.Builder) { builder.StartObject(11) }
func StatAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ID), 0) }
func StatAddUser(builder *flatbuffers.Builder, User int64) { builder.PrependInt64Slot(1, User, 0) }
func StatAddNice(builder *flatbuffers.Builder, Nice int64) { builder.PrependInt64Slot(2, Nice, 0) }
func StatAddSystem(builder *flatbuffers.Builder, System int64) { builder.PrependInt64Slot(3, System, 0) }
func StatAddIdle(builder *flatbuffers.Builder, Idle int64) { builder.PrependInt64Slot(4, Idle, 0) }
func StatAddIOWait(builder *flatbuffers.Builder, IOWait int64) { builder.PrependInt64Slot(5, IOWait, 0) }
func StatAddIRQ(builder *flatbuffers.Builder, IRQ int64) { builder.PrependInt64Slot(6, IRQ, 0) }
func StatAddSoftIRQ(builder *flatbuffers.Builder, SoftIRQ int64) { builder.PrependInt64Slot(7, SoftIRQ, 0) }
func StatAddSteal(builder *flatbuffers.Builder, Steal int64) { builder.PrependInt64Slot(8, Steal, 0) }
func StatAddQuest(builder *flatbuffers.Builder, Quest int64) { builder.PrependInt64Slot(9, Quest, 0) }
func StatAddQuestNice(builder *flatbuffers.Builder, QuestNice int64) { builder.PrependInt64Slot(10, QuestNice, 0) }
func StatEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
