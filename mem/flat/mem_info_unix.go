// Copyright 2016 Joel Scoble and The JoeFriday authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package flat handles Flatbuffer based processing of /proc/meminfo.
// Instead of returning a Go struct, it returns Flatbuffer serialized bytes.
// A function to deserialize the Flatbuffer serialized bytes into a
// mem.Info struct is provided.  After the first use, the flatbuffer
// builder is reused.
package flat

import (
	"io"
	"sync"
	"time"

	"github.com/SermoDigital/helpers"
	fb "github.com/google/flatbuffers/go"
	joe "github.com/mohae/joefriday"
	"github.com/mohae/joefriday/mem"
)

// Profiler is used to process the /proc/meminfo file using Flatbuffers.
type Profiler struct {
	*mem.Profiler
	*fb.Builder
}

// Initializes and returns a mem info profiler that utilizes FlatBuffers.
func NewProfiler() (prof *Profiler, err error) {
	p, err := mem.NewProfiler()
	if err != nil {
		return nil, err
	}
	return &Profiler{Profiler: p, Builder: fb.NewBuilder(0)}, nil
}

// Get returns the current meminfo as Flatbuffer serialized bytes.
func (prof *Profiler) Get() ([]byte, error) {
	inf, err := prof.Profiler.Get()
	if err != nil {
		return nil, err
	}
	return prof.Serialize(inf), nil
}

var std *Profiler
var stdMu sync.Mutex //protects standard to preven data race on checking/instantiation

// Get returns the current meminfo as Flatbuffer serialized bytes using the
// package's global Profiler.
func Get() (p []byte, err error) {
	stdMu.Lock()
	defer stdMu.Unlock()
	if std == nil {
		std, err = NewProfiler()
		if err != nil {
			return nil, err
		}
	}
	return std.Get()
}

// Serialize mem.Info using Flatbuffers.
func (prof *Profiler) Serialize(inf *mem.Info) []byte {
	// ensure the Builder is in a usable state.
	prof.Builder.Reset()
	InfoStart(prof.Builder)
	InfoAddTimestamp(prof.Builder, inf.Timestamp)
	InfoAddActive(prof.Builder, inf.Active)
	InfoAddActiveAnon(prof.Builder, inf.ActiveAnon)
	InfoAddActiveFile(prof.Builder, inf.ActiveFile)
	InfoAddAnonHugePages(prof.Builder, inf.AnonHugePages)
	InfoAddAnonPages(prof.Builder, inf.AnonPages)
	InfoAddBounce(prof.Builder, inf.Bounce)
	InfoAddBuffers(prof.Builder, inf.Buffers)
	InfoAddCached(prof.Builder, inf.Cached)
	InfoAddCommitLimit(prof.Builder, inf.CommitLimit)
	InfoAddCommittedAS(prof.Builder, inf.CommittedAS)
	InfoAddDirectMap4K(prof.Builder, inf.DirectMap4K)
	InfoAddDirectMap2M(prof.Builder, inf.DirectMap2M)
	InfoAddDirty(prof.Builder, inf.Dirty)
	InfoAddHardwareCorrupted(prof.Builder, inf.HardwareCorrupted)
	InfoAddHugePagesFree(prof.Builder, inf.HugePagesFree)
	InfoAddHugePagesRsvd(prof.Builder, inf.HugePagesRsvd)
	InfoAddHugePagesSize(prof.Builder, inf.HugePagesSize)
	InfoAddHugePagesSurp(prof.Builder, inf.HugePagesSurp)
	InfoAddHugePagesTotal(prof.Builder, inf.HugePagesTotal)
	InfoAddInactive(prof.Builder, inf.Inactive)
	InfoAddInactiveAnon(prof.Builder, inf.InactiveAnon)
	InfoAddInactiveFile(prof.Builder, inf.InactiveFile)
	InfoAddKernelStack(prof.Builder, inf.KernelStack)
	InfoAddMapped(prof.Builder, inf.Mapped)
	InfoAddMemAvailable(prof.Builder, inf.MemAvailable)
	InfoAddMemFree(prof.Builder, inf.MemFree)
	InfoAddMemTotal(prof.Builder, inf.MemTotal)
	InfoAddMlocked(prof.Builder, inf.Mlocked)
	InfoAddNFSUnstable(prof.Builder, inf.NFSUnstable)
	InfoAddPageTables(prof.Builder, inf.PageTables)
	InfoAddShmem(prof.Builder, inf.Shmem)
	InfoAddSlab(prof.Builder, inf.Slab)
	InfoAddSReclaimable(prof.Builder, inf.SReclaimable)
	InfoAddSUnreclaim(prof.Builder, inf.SUnreclaim)
	InfoAddSwapCached(prof.Builder, inf.SwapCached)
	InfoAddSwapFree(prof.Builder, inf.SwapFree)
	InfoAddSwapTotal(prof.Builder, inf.SwapTotal)
	InfoAddUnevictable(prof.Builder, inf.Unevictable)
	InfoAddVmallocChunk(prof.Builder, inf.VmallocChunk)
	InfoAddVmallocTotal(prof.Builder, inf.VmallocTotal)
	InfoAddVmallocUsed(prof.Builder, inf.VmallocUsed)
	InfoAddWriteback(prof.Builder, inf.Writeback)
	InfoAddWritebackTmp(prof.Builder, inf.WritebackTmp)
	prof.Builder.Finish(InfoEnd(prof.Builder))
	p := prof.Builder.Bytes[prof.Builder.Head():]
	// copy them (otherwise gets lost in reset)
	tmp := make([]byte, len(p))
	copy(tmp, p)
	return tmp
}

// Serialize mem.Info using Flatbuffers with the package global Profiler.
func Serialize(inf *mem.Info) (p []byte, err error) {
	stdMu.Lock()
	defer stdMu.Unlock()
	if std == nil {
		std, err = NewProfiler()
		if err != nil {
			return nil, err
		}
	}
	return std.Serialize(inf), nil
}

// Deserialize takes some Flatbuffer serialized bytes and deserialize's them
// as mem.Info.
func Deserialize(p []byte) *mem.Info {
	infoFlat := GetRootAsInfo(p, 0)
	info := &mem.Info{}
	info.Timestamp = infoFlat.Timestamp()
	info.Active = infoFlat.Active()
	info.ActiveAnon = infoFlat.ActiveAnon()
	info.ActiveFile = infoFlat.ActiveFile()
	info.AnonHugePages = infoFlat.AnonHugePages()
	info.AnonPages = infoFlat.AnonPages()
	info.Bounce = infoFlat.Bounce()
	info.Buffers = infoFlat.Buffers()
	info.Cached = infoFlat.Cached()
	info.CommitLimit = infoFlat.CommitLimit()
	info.CommittedAS = infoFlat.CommittedAS()
	info.DirectMap4K = infoFlat.DirectMap4K()
	info.DirectMap2M = infoFlat.DirectMap2M()
	info.Dirty = infoFlat.Dirty()
	info.HardwareCorrupted = infoFlat.HardwareCorrupted()
	info.HugePagesFree = infoFlat.HugePagesFree()
	info.HugePagesRsvd = infoFlat.HugePagesRsvd()
	info.HugePagesSize = infoFlat.HugePagesSize()
	info.HugePagesSurp = infoFlat.HugePagesSurp()
	info.HugePagesTotal = infoFlat.HugePagesTotal()
	info.Inactive = infoFlat.Inactive()
	info.InactiveAnon = infoFlat.InactiveAnon()
	info.InactiveFile = infoFlat.InactiveFile()
	info.KernelStack = infoFlat.KernelStack()
	info.Mapped = infoFlat.Mapped()
	info.MemAvailable = infoFlat.MemAvailable()
	info.MemFree = infoFlat.MemFree()
	info.MemTotal = infoFlat.MemTotal()
	info.Mlocked = infoFlat.Mlocked()
	info.NFSUnstable = infoFlat.NFSUnstable()
	info.PageTables = infoFlat.PageTables()
	info.Shmem = infoFlat.Shmem()
	info.Slab = infoFlat.Slab()
	info.SReclaimable = infoFlat.SReclaimable()
	info.SUnreclaim = infoFlat.SUnreclaim()
	info.SwapCached = infoFlat.SwapCached()
	info.SwapFree = infoFlat.SwapFree()
	info.SwapTotal = infoFlat.SwapTotal()
	info.Unevictable = infoFlat.Unevictable()
	info.VmallocChunk = infoFlat.VmallocChunk()
	info.VmallocTotal = infoFlat.VmallocTotal()
	info.VmallocUsed = infoFlat.VmallocUsed()
	info.Writeback = infoFlat.Writeback()
	info.WritebackTmp = infoFlat.WritebackTmp()
	return info
}

// Ticker delivers the system's memory information at intervals.
type Ticker struct {
	*joe.Ticker
	Data chan []byte
	*Profiler
}

// NewTicker returns a new Ticker continaing a Data channel that delivers
// the data at intervals and an error channel that delivers any errors
// encountered.  Stop the ticker to signal the ticker to stop running; it
// does not close the Data channel.  Close the ticker to close all ticker
// channels.
func NewTicker(d time.Duration) (joe.Tocker, error) {
	p, err := NewProfiler()
	if err != nil {
		return nil, err
	}
	t := Ticker{Ticker: joe.NewTicker(d), Data: make(chan []byte), Profiler: p}
	go t.Run()
	return &t, nil
}

// Run runs the ticker.
func (t *Ticker) Run() {
	// predeclare some vars
	var (
		i, pos, nameLen int
		v               byte
		n               uint64
		err             error
	)
	// ticker
Tick:
	for {
		select {
		case <-t.Done:
			return
		case <-t.C:
			t.Builder.Reset()
			err = t.Profiler.Profiler.Reset()
			if err != nil {
				t.Errs <- err
				continue
			}
			InfoStart(t.Builder)
			InfoAddTimestamp(t.Builder, time.Now().UTC().UnixNano())
			for {
				t.Val = t.Val[:0]
				t.Line, err = t.Buf.ReadSlice('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					// An error results in sending error message and stop processing of this tick.
					t.Errs <- &joe.ReadError{Err: err}
					continue Tick
				}
				// first grab the key name (everything up to the ':')
				for i, v = range t.Line {
					if v == 0x3A {
						t.Val = t.Line[:i]
						pos = i + 1 // skip the :
						break
					}
				}
				nameLen = len(t.Val)
				// skip all spaces
				for i, v = range t.Line[pos:] {
					if v != 0x20 {
						pos += i
						break
					}
				}

				// grab the numbers
				for _, v = range t.Line[pos:] {
					if v == 0x20 || v == '\n' {
						break
					}
					t.Val = append(t.Val, v)
				}
				// any conversion error results in 0
				n, err = helpers.ParseUint(t.Val[nameLen:])
				if err != nil {
					t.Errs <- &joe.ParseError{Info: string(t.Val[:nameLen]), Err: err}
					continue
				}
				v = t.Val[0]
				if v == 'A' {
					if t.Val[5] == 'e' {
						if nameLen == 6 {
							InfoAddActive(t.Builder, n)
							continue
						}
						if t.Val[7] == 'a' {
							InfoAddActiveAnon(t.Builder, n)
							continue
						}
						InfoAddActiveFile(t.Builder, n)
						continue
					}
					if nameLen == 9 {
						InfoAddAnonPages(t.Builder, n)
						continue
					}
					InfoAddAnonHugePages(t.Builder, n)
					continue
				}
				if v == 'C' {
					if nameLen == 6 {
						InfoAddCached(t.Builder, n)
						continue
					}
					if nameLen == 11 {
						InfoAddCommitLimit(t.Builder, n)
						continue
					}
					InfoAddCommittedAS(t.Builder, n)
					continue
				}
				if v == 'D' {
					if nameLen == 5 {
						InfoAddDirty(t.Builder, n)
						continue
					}
					if t.Val[10] == 'k' {
						InfoAddDirectMap4K(t.Builder, n)
						continue
					}
					InfoAddDirectMap2M(t.Builder, n)
					continue
				}
				if v == 'H' {
					if nameLen == 14 {
						if t.Val[10] == 'F' {
							InfoAddHugePagesFree(t.Builder, n)
							continue
						}
						if t.Val[10] == 'R' {
							InfoAddHugePagesRsvd(t.Builder, n)
							continue
						}
						InfoAddHugePagesSurp(t.Builder, n)
					}
					if t.Val[1] == 'a' {
						InfoAddHardwareCorrupted(t.Builder, n)
						continue
					}
					if t.Val[9] == 'i' {
						InfoAddHugePagesSize(t.Builder, n)
						continue
					}
					InfoAddHugePagesTotal(t.Builder, n)
					continue
				}
				if v == 'I' {
					if nameLen == 8 {
						InfoAddInactive(t.Builder, n)
						continue
					}
					if t.Val[9] == 'a' {
						InfoAddInactiveAnon(t.Builder, n)
						continue
					}
					InfoAddInactiveFile(t.Builder, n)
				}
				if v == 'M' {
					v = t.Val[3]
					if nameLen < 8 {
						if v == 'p' {
							InfoAddMapped(t.Builder, n)
							continue
						}
						if v == 'F' {
							InfoAddMemFree(t.Builder, n)
							continue
						}
						InfoAddMlocked(t.Builder, n)
						continue
					}
					if v == 'A' {
						InfoAddMemAvailable(t.Builder, n)
						continue
					}
					InfoAddMemTotal(t.Builder, n)
					continue
				}
				if v == 'S' {
					v = t.Val[1]
					if v == 'w' {
						if t.Val[4] == 'C' {
							InfoAddSwapCached(t.Builder, n)
							continue
						}
						if t.Val[4] == 'F' {
							InfoAddSwapFree(t.Builder, n)
							continue
						}
						InfoAddSwapTotal(t.Builder, n)
						continue
					}
					if v == 'h' {
						InfoAddShmem(t.Builder, n)
						continue
					}
					if v == 'l' {
						InfoAddSlab(t.Builder, n)
						continue
					}
					if v == 'R' {
						InfoAddSReclaimable(t.Builder, n)
						continue
					}
					InfoAddSUnreclaim(t.Builder, n)
					continue
				}
				if v == 'V' {
					if t.Val[8] == 'C' {
						InfoAddVmallocChunk(t.Builder, n)
						continue
					}
					if t.Val[8] == 'T' {
						InfoAddVmallocTotal(t.Builder, n)
						continue
					}
					InfoAddVmallocUsed(t.Builder, n)
					continue
				}
				if v == 'W' {
					if nameLen == 9 {
						InfoAddWriteback(t.Builder, n)
						continue
					}
					InfoAddWritebackTmp(t.Builder, n)
					continue
				}
				if v == 'B' {
					if nameLen == 6 {
						InfoAddBounce(t.Builder, n)
						continue
					}
					InfoAddBuffers(t.Builder, n)
					continue
				}
				if v == 'K' {
					InfoAddKernelStack(t.Builder, n)
					continue
				}
				if v == 'N' {
					InfoAddNFSUnstable(t.Builder, n)
					continue
				}
				if v == 'P' {
					InfoAddPageTables(t.Builder, n)
				}
				InfoAddUnevictable(t.Builder, n)
			}
			t.Builder.Finish(InfoEnd(t.Builder))
			t.Data <- t.Profiler.Builder.Bytes[t.Builder.Head():]
		}
	}
}

// Close closes the ticker resources.
func (t *Ticker) Close() {
	t.Ticker.Close()
	close(t.Data)
}
