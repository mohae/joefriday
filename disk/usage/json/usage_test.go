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

package json

import (
	"testing"
	"time"

	"github.com/mohae/joefriday/disk/structs"
)

func TestGet(t *testing.T) {
	p, err := New()
	if err != nil {
		t.Errorf("got %s, want nil", err)
		return
	}
	b, err := p.Get()
	if err != nil {
		t.Errorf("got %s, want nil", err)
		return
	}
	u, err := Unmarshal(b)
	if err != nil {
		t.Errorf("got %s, want nil", err)
		return
	}
	checkUsage("get", u, t)
	t.Logf("%#v\n", u)
}

func TestTicker(t *testing.T) {
	out := make(chan []byte)
	done := make(chan struct{})
	errs := make(chan error)
	go Ticker(time.Duration(400)*time.Millisecond, out, done, errs)
	var err error
	var u *structs.Usage
	for i := 0; i < 1; i++ {
		select {
		case b := <-out:
			u, err = Unmarshal(b)
			if err != nil {
				t.Errorf("got %s, want nil", err)
				return
			}
			checkUsage("ticker", u, t)
		case err := <-errs:
			t.Errorf("unexpected error: %s", err)
		}
	}
	t.Logf("%#v\n", u)
}

func checkUsage(n string, u *structs.Usage, t *testing.T) {
	if u.Timestamp == 0 {
		t.Errorf("%s: Timestamp: wanted non-zero value; got 0", n)
	}
	if u.TimeDelta == 0 {
		t.Errorf("%s: TimeDelta: wanted non-zero value; got 0", n)
	}
	if len(u.Devices) == 0 {
		t.Errorf("%s: expected there to be devices; didn't get any", n)
	}
	for i := 0; i < len(u.Devices); i++ {
		if u.Devices[i].Major == 0 {
			t.Errorf("%s: Device %d: Major: wanted a non-zero value, was 0", n, i)
		}
		if u.Devices[i].Name == "" {
			t.Errorf("%s: Device %d: Name: wanted a non-empty value; was empty", n, i)
		}
	}
}