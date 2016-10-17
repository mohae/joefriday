// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Processors struct {
	_tab flatbuffers.Table
}

func GetRootAsProcessors(buf []byte, offset flatbuffers.UOffsetT) *Processors {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Processors{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Processors) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Processors) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Processors) Count() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Processors) Chips(obj *Chip, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(Chip)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Processors) ChipsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ProcessorsStart(builder *flatbuffers.Builder) { builder.StartObject(3) }
func ProcessorsAddTimestamp(builder *flatbuffers.Builder, Timestamp int64) { builder.PrependInt64Slot(0, Timestamp, 0) }
func ProcessorsAddCount(builder *flatbuffers.Builder, Count int16) { builder.PrependInt16Slot(1, Count, 0) }
func ProcessorsAddChips(builder *flatbuffers.Builder, Chips flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Chips), 0) }
func ProcessorsStartChipsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func ProcessorsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
