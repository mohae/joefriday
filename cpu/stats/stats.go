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

// Package stat handles processing of the /proc/stats file: information about
// kernel activity.
package stats

import (
	"bytes"
	"io"
	"os/exec"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/SermoDigital/helpers"
	joe "github.com/mohae/joefriday"
)

const procFile = "/proc/stat"

var CLK_TCK int32    // the ticks per clock cycle
var tckMu sync.Mutex //protects CLK_TCK

// Set CLK_TCK.
func ClkTck() error {
	tckMu.Lock()
	defer tckMu.Unlock()
	var out bytes.Buffer
	cmd := exec.Command("getconf", "CLK_TCK")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return joe.Error{Type: "cpu", Op: "get conf CLK_TCK", Err: err}
	}
	b, err := out.ReadBytes('\n')
	if err != nil {
		return joe.Error{Type: "cpu", Op: "read conf CLK_TCK output", Err: err}
	}
	v, err := strconv.Atoi(string(b[:len(b)-1]))
	if err != nil {
		return joe.Error{Type: "cpu", Op: "processing conf CLK_TCK output", Err: err}
	}
	atomic.StoreInt32(&CLK_TCK, int32(v))
	return nil
}

// Stats holds the /proc/stat information
type Stats struct {
	ClkTck    int16  `json:"clk_tck"`
	Timestamp int64  `json:"timestamp"`
	Ctxt      int64  `json:"ctxt"`
	BTime     int64  `json:"btime"`
	Processes int64  `json:"processes"`
	CPU       []Stat `json:"cpu"`
}

// Stat is for capturing the CPU information of /proc/stat.
type Stat struct {
	ID        string `json:"ID"`
	User      int64  `json:"user"`
	Nice      int64  `json:"nice"`
	System    int64  `json:"system"`
	Idle      int64  `json:"idle"`
	IOWait    int64  `json:"io_wait"`
	IRQ       int64  `json:"irq"`
	SoftIRQ   int64  `json:"soft_irq"`
	Steal     int64  `json:"steal"`
	Quest     int64  `json:"quest"`
	QuestNice int64  `json:"quest_nice"`
}

// Profiler is used to process the /proc/stats file.
type Profiler struct {
	*joe.Proc
	ClkTck int16
}

// Returns an initialized Profiler; ready to use.
func NewProfiler() (prof *Profiler, err error) {
	// if it hasn't been set, set it.
	if atomic.LoadInt32(&CLK_TCK) == 0 {
		err = ClkTck()
		if err != nil {
			return nil, err
		}
	}
	proc, err := joe.New(procFile)
	if err != nil {
		return nil, err
	}
	return &Profiler{Proc: proc, ClkTck: int16(atomic.LoadInt32(&CLK_TCK))}, nil
}

// Get returns information about current kernel activity.
func (prof *Profiler) Get() (stats *Stats, err error) {
	err = prof.Reset()
	if err != nil {
		return nil, err
	}
	var (
		i, j, pos, fieldNum int
		n                   uint64
		v                   byte
		stop                bool
	)

	stats = &Stats{Timestamp: time.Now().UTC().UnixNano(), ClkTck: prof.ClkTck, CPU: make([]Stat, 0, 2)}

	// read each line until eof
	for {
		prof.Line, err = prof.Buf.ReadSlice('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, joe.Error{Type: "cpu stat", Op: "reading /proc/stat output", Err: err}
		}
		prof.Val = prof.Val[:0]
		// Get everything up to the first space, this is the key.  Not all keys are processed.
		for i, v = range prof.Line {
			if v == 0x20 {
				prof.Val = prof.Line[:i]
				pos = i + 1
				break
			}
		}
		// skip the intr line
		if prof.Val[0] == 'i' {
			continue
		}
		if prof.Val[0] == 'c' {
			if prof.Val[1] == 'p' { // process CPU
				stat := Stat{ID: string(prof.Val[:])}
				j = 0
				// skip over any remaining spaces
				for i, v = range prof.Line[pos:] {
					if v != 0x20 {
						break
					}
					j++
				}
				fieldNum = 0
				pos, j = j+pos, j+pos
				// space is the field separator
				for i, v = range prof.Line[pos:] {
					if v == '\n' {
						stop = true
					}
					if v == 0x20 || stop {
						fieldNum++
						n, err = helpers.ParseUint(prof.Line[j : pos+i])
						if err != nil {
							return stats, joe.Error{Type: "cpu stat", Op: "convert cpu data", Err: err}
						}
						j = pos + i + 1
						if fieldNum < 6 {
							if fieldNum < 4 {
								if fieldNum == 1 {
									stat.User = int64(n)
									continue
								}
								if fieldNum == 2 {
									stat.Nice = int64(n)
									continue
								}
								stat.System = int64(n) // 3
								continue
							}
							if fieldNum == 4 {
								stat.Idle = int64(n)
								continue
							}
							stat.IOWait = int64(n) // 5
							continue
						}
						if fieldNum < 8 {
							if fieldNum == 6 {
								stat.IRQ = int64(n)
								continue
							}
							stat.SoftIRQ = int64(n) // 7
							continue
						}
						if fieldNum == 8 {
							stat.Steal = int64(n)
							continue
						}
						if fieldNum == 9 {
							stat.Quest = int64(n)
							continue
						}
						stat.QuestNice = int64(n) // 10
					}
				}
				stats.CPU = append(stats.CPU, stat)
				stop = false
				continue
			}
			// Otherwise it's ctxt info; rest of the line is the data.
			n, err = helpers.ParseUint(prof.Line[pos : len(prof.Line)-1])
			if err != nil {
				return stats, joe.Error{Type: "cpu stat", Op: "convert ctxt data", Err: err}
			}
			stats.Ctxt = int64(n)
			continue
		}
		if prof.Val[0] == 'b' {
			// rest of the line is the data
			n, err = helpers.ParseUint(prof.Line[pos : len(prof.Line)-1])
			if err != nil {
				return stats, joe.Error{Type: "cpu stat", Op: "convert btime data", Err: err}
			}
			stats.BTime = int64(n)
			continue
		}
		if prof.Val[0] == 'p' && prof.Val[4] == 'e' { // processes info
			// rest of the line is the data
			n, err = helpers.ParseUint(prof.Line[pos : len(prof.Line)-1])
			if err != nil {
				return stats, joe.Error{Type: "cpu stat", Op: "convert processes data", Err: err}
			}
			stats.Processes = int64(n)
			continue
		}
	}
	return stats, nil
}

var std *Profiler
var stdMu sync.Mutex

// Get returns the current kernal activity information using the package's
// global Profiler.
func Get() (stat *Stats, err error) {
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

// Ticker delivers the system's memory information at intervals.
type Ticker struct {
	*joe.Ticker
	Data chan *Stats
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
	t := Ticker{Ticker: joe.NewTicker(d), Data: make(chan *Stats), Profiler: p}
	go t.Run()
	return &t, nil
}

// Run runs the ticker.
func (t *Ticker) Run() {
	// ticker
	for {
		select {
		case <-t.Done:
			return
		case <-t.Ticker.C:
			s, err := t.Profiler.Get()
			if err != nil {
				t.Errs <- err
				continue
			}
			t.Data <- s
		}
	}
}

// Close closes the ticker resources.
func (t *Ticker) Close() {
	t.Ticker.Close()
	close(t.Data)
}
