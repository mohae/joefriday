// automatically generated by the FlatBuffers compiler, do not modify

package structs

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

func (rcv *Processors) CPUs() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Processors) Sockets() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Processors) Socket(obj *Processor, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(Processor)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Processors) SocketLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ProcessorsStart(builder *flatbuffers.Builder) { builder.StartObject(4) }
func ProcessorsAddTimestamp(builder *flatbuffers.Builder, Timestamp int64) { builder.PrependInt64Slot(0, Timestamp, 0) }
func ProcessorsAddCPUs(builder *flatbuffers.Builder, CPUs int64) { builder.PrependInt64Slot(1, CPUs, 0) }
func ProcessorsAddSockets(builder *flatbuffers.Builder, Sockets int32) { builder.PrependInt32Slot(2, Sockets, 0) }
func ProcessorsAddSocket(builder *flatbuffers.Builder, Socket flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(Socket), 0) }
func ProcessorsStartSocketVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func ProcessorsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
