// automatically generated by the FlatBuffers compiler, do not modify

package structs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Node struct {
	_tab flatbuffers.Table
}

func (rcv *Node) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Node) ID() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Node) CPUList() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func NodeStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func NodeAddID(builder *flatbuffers.Builder, ID int32) { builder.PrependInt32Slot(0, ID, 0) }
func NodeAddCPUList(builder *flatbuffers.Builder, CPUList flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(CPUList), 0) }
func NodeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }