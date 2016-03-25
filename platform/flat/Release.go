// automatically generated, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Release struct {
	_tab flatbuffers.Table
}

func GetRootAsRelease(buf []byte, offset flatbuffers.UOffsetT) *Release {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Release{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Release) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Release) ID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Release) IDLike() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Release) PrettyName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Release) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Release) VersionID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Release) HomeURL() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Release) BugReportURL() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ReleaseStart(builder *flatbuffers.Builder) { builder.StartObject(7) }
func ReleaseAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ID), 0) }
func ReleaseAddIDLike(builder *flatbuffers.Builder, IDLike flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(IDLike), 0) }
func ReleaseAddPrettyName(builder *flatbuffers.Builder, PrettyName flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(PrettyName), 0) }
func ReleaseAddVersion(builder *flatbuffers.Builder, Version flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(Version), 0) }
func ReleaseAddVersionID(builder *flatbuffers.Builder, VersionID flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(VersionID), 0) }
func ReleaseAddHomeURL(builder *flatbuffers.Builder, HomeURL flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(HomeURL), 0) }
func ReleaseAddBugReportURL(builder *flatbuffers.Builder, BugReportURL flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(BugReportURL), 0) }
func ReleaseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
