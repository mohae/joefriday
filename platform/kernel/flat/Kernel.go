// automatically generated, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Kernel struct {
	_tab flatbuffers.Table
}

func GetRootAsKernel(buf []byte, offset flatbuffers.UOffsetT) *Kernel {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Kernel{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Kernel) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Kernel) OS() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Kernel) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Kernel) CompileUser() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Kernel) GCC() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Kernel) OSGCC() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Kernel) Type() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Kernel) CompileDate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Kernel) Arch() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func KernelStart(builder *flatbuffers.Builder) { builder.StartObject(8) }
func KernelAddOS(builder *flatbuffers.Builder, OS flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(OS), 0) }
func KernelAddVersion(builder *flatbuffers.Builder, Version flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Version), 0) }
func KernelAddCompileUser(builder *flatbuffers.Builder, CompileUser flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(CompileUser), 0) }
func KernelAddGCC(builder *flatbuffers.Builder, GCC flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(GCC), 0) }
func KernelAddOSGCC(builder *flatbuffers.Builder, OSGCC flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(OSGCC), 0) }
func KernelAddType(builder *flatbuffers.Builder, Type flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(Type), 0) }
func KernelAddCompileDate(builder *flatbuffers.Builder, CompileDate flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(CompileDate), 0) }
func KernelAddArch(builder *flatbuffers.Builder, Arch flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(Arch), 0) }
func KernelEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
