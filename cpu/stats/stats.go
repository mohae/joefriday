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

package stats

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

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

type Profiler struct {
	joe.Proc
	ClkTck int16
}

func New() (prof *Profiler, err error) {
	// if it hasn't been set, set it.
	if atomic.LoadInt32(&CLK_TCK) == 0 {
		err = ClkTck()
		if err != nil {
			return nil, err
		}
	}
	f, err := os.Open(procFile)
	if err != nil {
		return nil, err
	}
	return &Profiler{Proc: joe.Proc{File: f, Buf: bufio.NewReader(f)}, ClkTck: int16(atomic.LoadInt32(&CLK_TCK))}, nil
}

func (prof *Profiler) Get() (stats *Stats, err error) {
	prof.Reset()
	prof.Lock()
	defer prof.Unlock()
	if CLK_TCK == 0 {
		err := ClkTck()
		if err != nil {
			return stats, err
		}
	}
	var (
		name                     string
		i, j, pos, val, fieldNum int
		v                        byte
		stop                     bool
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
		// Get everything up to the first space, this is the key.  Not all keys are processed.
		for i, v = range prof.Line {
			if v == 0x20 {
				name = string(prof.Line[:i])
				pos = i + 1
				break
			}
		}
		// skip the intr line
		if name == "intr" {
			continue
		}
		if name[:3] == "cpu" {
			j = 0
			// skip over any remaining spaces
			for i, v = range prof.Line[pos:] {
				if v != 0x20 {
					break
				}
				j++
			}
			stat := Stat{ID: name}
			fieldNum = 0
			pos, j = j+pos, j+pos
			// space is the field separator
			for i, v = range prof.Line[pos:] {
				if v == '\n' {
					stop = true
				}
				if v == 0x20 || stop {
					fieldNum++
					val, err = strconv.Atoi(string(prof.Line[j : pos+i]))
					if err != nil {
						return stats, joe.Error{Type: "cpu stat", Op: "convert cpu data", Err: err}
					}
					j = pos + i + 1
					if fieldNum < 4 {
						if fieldNum == 1 {
							stat.User = int64(val)
						} else if fieldNum == 2 {
							stat.Nice = int64(val)
						} else if fieldNum == 3 {
							stat.System = int64(val)
						}
					} else if fieldNum < 7 {
						if fieldNum == 4 {
							stat.Idle = int64(val)
						} else if fieldNum == 5 {
							stat.IOWait = int64(val)
						} else if fieldNum == 6 {
							stat.IRQ = int64(val)
						}
					} else if fieldNum < 10 {
						if fieldNum == 7 {
							stat.SoftIRQ = int64(val)
						} else if fieldNum == 8 {
							stat.Steal = int64(val)
						} else if fieldNum == 9 {
							stat.Quest = int64(val)
						}
					} else if fieldNum == 10 {
						stat.QuestNice = int64(val)
					}
				}
			}
			stats.CPU = append(stats.CPU, stat)
			stop = false
			continue
		}
		if name == "ctxt" {
			// rest of the line is the data
			val, err = strconv.Atoi(string(prof.Line[pos : len(prof.Line)-1]))
			if err != nil {
				return stats, joe.Error{Type: "cpu stat", Op: "convert ctxt data", Err: err}
			}
			stats.Ctxt = int64(val)
			continue
		}
		if name == "btime" {
			// rest of the line is the data
			val, err = strconv.Atoi(string(prof.Line[pos : len(prof.Line)-1]))
			if err != nil {
				return stats, joe.Error{Type: "cpu stat", Op: "convert btime data", Err: err}
			}
			stats.BTime = int64(val)
			continue
		}
		if name == "processes" {
			// rest of the line is the data
			val, err = strconv.Atoi(string(prof.Line[pos : len(prof.Line)-1]))
			if err != nil {
				return stats, joe.Error{Type: "cpu stat", Op: "convert processes data", Err: err}
			}
			stats.Processes = int64(val)
			continue
		}
	}
	return stats, nil
}

var std *Profiler
var stdMu sync.Mutex

// Get gets the output of /proc/stat.
func Get() (stat *Stats, err error) {
	stdMu.Lock()
	defer stdMu.Unlock()
	if std == nil {
		std, err = New()
		if err != nil {
			return nil, err
		}
	}
	return std.Get()
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
