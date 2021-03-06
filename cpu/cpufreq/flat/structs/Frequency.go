// automatically generated by the FlatBuffers compiler, do not modify

package structs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Frequency struct {
	_tab flatbuffers.Table
}

func GetRootAsFrequency(buf []byte, offset flatbuffers.UOffsetT) *Frequency {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Frequency{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Frequency) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Frequency) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Frequency) Sockets() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Frequency) CPU(obj *CPU, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(CPU)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Frequency) CPULength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func FrequencyStart(builder *flatbuffers.Builder) { builder.StartObject(3) }
func FrequencyAddTimestamp(builder *flatbuffers.Builder, Timestamp int64) { builder.PrependInt64Slot(0, Timestamp, 0) }
func FrequencyAddSockets(builder *flatbuffers.Builder, Sockets int32) { builder.PrependInt32Slot(1, Sockets, 0) }
func FrequencyAddCPU(builder *flatbuffers.Builder, CPU flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(CPU), 0) }
func FrequencyStartCPUVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func FrequencyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
