// automatically generated, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Info struct {
	_tab flatbuffers.Table
}

func GetRootAsInfo(buf []byte, offset flatbuffers.UOffsetT) *Info {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Info{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *Info) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Info) Timestamp() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Active() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) ActiveAnon() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) ActiveFile() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) AnonHugePages() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) AnonPages() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Bounce() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Buffers() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Cached() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) CommitLimit() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) CommittedAS() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) DirectMap4K() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) DirectMap2M() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Dirty() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) HardwareCorrupted() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) HugePagesFree() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) HugePagesRsvd() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) HugePagesSize() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) HugePagesSurp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) HugePagesTotal() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Inactive() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) InactiveAnon() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) InactiveFile() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) KernelStack() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Mapped() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) MemAvailable() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) MemFree() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) MemTotal() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Mlocked() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) NFSUnstable() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) PageTables() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Shmem() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Slab() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(68))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) SReclaimable() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(70))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) SUnreclaim() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(72))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) SwapCached() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(74))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) SwapFree() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(76))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) SwapTotal() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(78))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Unevictable() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(80))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) VmallocChunk() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(82))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) VmallocTotal() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(84))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) VmallocUsed() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(86))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) Writeback() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(88))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Info) WritebackTmp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(90))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func InfoStart(builder *flatbuffers.Builder) { builder.StartObject(44) }
func InfoAddTimestamp(builder *flatbuffers.Builder, Timestamp int64) { builder.PrependInt64Slot(0, Timestamp, 0) }
func InfoAddActive(builder *flatbuffers.Builder, Active uint64) { builder.PrependUint64Slot(1, Active, 0) }
func InfoAddActiveAnon(builder *flatbuffers.Builder, ActiveAnon uint64) { builder.PrependUint64Slot(2, ActiveAnon, 0) }
func InfoAddActiveFile(builder *flatbuffers.Builder, ActiveFile uint64) { builder.PrependUint64Slot(3, ActiveFile, 0) }
func InfoAddAnonHugePages(builder *flatbuffers.Builder, AnonHugePages uint64) { builder.PrependUint64Slot(4, AnonHugePages, 0) }
func InfoAddAnonPages(builder *flatbuffers.Builder, AnonPages uint64) { builder.PrependUint64Slot(5, AnonPages, 0) }
func InfoAddBounce(builder *flatbuffers.Builder, Bounce uint64) { builder.PrependUint64Slot(6, Bounce, 0) }
func InfoAddBuffers(builder *flatbuffers.Builder, Buffers uint64) { builder.PrependUint64Slot(7, Buffers, 0) }
func InfoAddCached(builder *flatbuffers.Builder, Cached uint64) { builder.PrependUint64Slot(8, Cached, 0) }
func InfoAddCommitLimit(builder *flatbuffers.Builder, CommitLimit uint64) { builder.PrependUint64Slot(9, CommitLimit, 0) }
func InfoAddCommittedAS(builder *flatbuffers.Builder, CommittedAS uint64) { builder.PrependUint64Slot(10, CommittedAS, 0) }
func InfoAddDirectMap4K(builder *flatbuffers.Builder, DirectMap4K uint64) { builder.PrependUint64Slot(11, DirectMap4K, 0) }
func InfoAddDirectMap2M(builder *flatbuffers.Builder, DirectMap2M uint64) { builder.PrependUint64Slot(12, DirectMap2M, 0) }
func InfoAddDirty(builder *flatbuffers.Builder, Dirty uint64) { builder.PrependUint64Slot(13, Dirty, 0) }
func InfoAddHardwareCorrupted(builder *flatbuffers.Builder, HardwareCorrupted uint64) { builder.PrependUint64Slot(14, HardwareCorrupted, 0) }
func InfoAddHugePagesFree(builder *flatbuffers.Builder, HugePagesFree uint64) { builder.PrependUint64Slot(15, HugePagesFree, 0) }
func InfoAddHugePagesRsvd(builder *flatbuffers.Builder, HugePagesRsvd uint64) { builder.PrependUint64Slot(16, HugePagesRsvd, 0) }
func InfoAddHugePagesSize(builder *flatbuffers.Builder, HugePagesSize uint64) { builder.PrependUint64Slot(17, HugePagesSize, 0) }
func InfoAddHugePagesSurp(builder *flatbuffers.Builder, HugePagesSurp uint64) { builder.PrependUint64Slot(18, HugePagesSurp, 0) }
func InfoAddHugePagesTotal(builder *flatbuffers.Builder, HugePagesTotal uint64) { builder.PrependUint64Slot(19, HugePagesTotal, 0) }
func InfoAddInactive(builder *flatbuffers.Builder, Inactive uint64) { builder.PrependUint64Slot(20, Inactive, 0) }
func InfoAddInactiveAnon(builder *flatbuffers.Builder, InactiveAnon uint64) { builder.PrependUint64Slot(21, InactiveAnon, 0) }
func InfoAddInactiveFile(builder *flatbuffers.Builder, InactiveFile uint64) { builder.PrependUint64Slot(22, InactiveFile, 0) }
func InfoAddKernelStack(builder *flatbuffers.Builder, KernelStack uint64) { builder.PrependUint64Slot(23, KernelStack, 0) }
func InfoAddMapped(builder *flatbuffers.Builder, Mapped uint64) { builder.PrependUint64Slot(24, Mapped, 0) }
func InfoAddMemAvailable(builder *flatbuffers.Builder, MemAvailable uint64) { builder.PrependUint64Slot(25, MemAvailable, 0) }
func InfoAddMemFree(builder *flatbuffers.Builder, MemFree uint64) { builder.PrependUint64Slot(26, MemFree, 0) }
func InfoAddMemTotal(builder *flatbuffers.Builder, MemTotal uint64) { builder.PrependUint64Slot(27, MemTotal, 0) }
func InfoAddMlocked(builder *flatbuffers.Builder, Mlocked uint64) { builder.PrependUint64Slot(28, Mlocked, 0) }
func InfoAddNFSUnstable(builder *flatbuffers.Builder, NFSUnstable uint64) { builder.PrependUint64Slot(29, NFSUnstable, 0) }
func InfoAddPageTables(builder *flatbuffers.Builder, PageTables uint64) { builder.PrependUint64Slot(30, PageTables, 0) }
func InfoAddShmem(builder *flatbuffers.Builder, Shmem uint64) { builder.PrependUint64Slot(31, Shmem, 0) }
func InfoAddSlab(builder *flatbuffers.Builder, Slab uint64) { builder.PrependUint64Slot(32, Slab, 0) }
func InfoAddSReclaimable(builder *flatbuffers.Builder, SReclaimable uint64) { builder.PrependUint64Slot(33, SReclaimable, 0) }
func InfoAddSUnreclaim(builder *flatbuffers.Builder, SUnreclaim uint64) { builder.PrependUint64Slot(34, SUnreclaim, 0) }
func InfoAddSwapCached(builder *flatbuffers.Builder, SwapCached uint64) { builder.PrependUint64Slot(35, SwapCached, 0) }
func InfoAddSwapFree(builder *flatbuffers.Builder, SwapFree uint64) { builder.PrependUint64Slot(36, SwapFree, 0) }
func InfoAddSwapTotal(builder *flatbuffers.Builder, SwapTotal uint64) { builder.PrependUint64Slot(37, SwapTotal, 0) }
func InfoAddUnevictable(builder *flatbuffers.Builder, Unevictable uint64) { builder.PrependUint64Slot(38, Unevictable, 0) }
func InfoAddVmallocChunk(builder *flatbuffers.Builder, VmallocChunk uint64) { builder.PrependUint64Slot(39, VmallocChunk, 0) }
func InfoAddVmallocTotal(builder *flatbuffers.Builder, VmallocTotal uint64) { builder.PrependUint64Slot(40, VmallocTotal, 0) }
func InfoAddVmallocUsed(builder *flatbuffers.Builder, VmallocUsed uint64) { builder.PrependUint64Slot(41, VmallocUsed, 0) }
func InfoAddWriteback(builder *flatbuffers.Builder, Writeback uint64) { builder.PrependUint64Slot(42, Writeback, 0) }
func InfoAddWritebackTmp(builder *flatbuffers.Builder, WritebackTmp uint64) { builder.PrependUint64Slot(43, WritebackTmp, 0) }
func InfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
