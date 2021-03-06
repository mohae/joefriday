// automatically generated by the FlatBuffers compiler, do not modify

package structs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type CPU struct {
	_tab flatbuffers.Table
}

func (rcv *CPU) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CPU) ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CPU) User() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) Nice() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) System() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) Idle() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) IOWait() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) IRQ() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) SoftIRQ() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) Steal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) Quest() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CPU) QuestNice() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func CPUStart(builder *flatbuffers.Builder) { builder.StartObject(11) }
func CPUAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ID), 0) }
func CPUAddUser(builder *flatbuffers.Builder, User int64) { builder.PrependInt64Slot(1, User, 0) }
func CPUAddNice(builder *flatbuffers.Builder, Nice int64) { builder.PrependInt64Slot(2, Nice, 0) }
func CPUAddSystem(builder *flatbuffers.Builder, System int64) { builder.PrependInt64Slot(3, System, 0) }
func CPUAddIdle(builder *flatbuffers.Builder, Idle int64) { builder.PrependInt64Slot(4, Idle, 0) }
func CPUAddIOWait(builder *flatbuffers.Builder, IOWait int64) { builder.PrependInt64Slot(5, IOWait, 0) }
func CPUAddIRQ(builder *flatbuffers.Builder, IRQ int64) { builder.PrependInt64Slot(6, IRQ, 0) }
func CPUAddSoftIRQ(builder *flatbuffers.Builder, SoftIRQ int64) { builder.PrependInt64Slot(7, SoftIRQ, 0) }
func CPUAddSteal(builder *flatbuffers.Builder, Steal int64) { builder.PrependInt64Slot(8, Steal, 0) }
func CPUAddQuest(builder *flatbuffers.Builder, Quest int64) { builder.PrependInt64Slot(9, Quest, 0) }
func CPUAddQuestNice(builder *flatbuffers.Builder, QuestNice int64) { builder.PrependInt64Slot(10, QuestNice, 0) }
func CPUEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
