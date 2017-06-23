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

package release

import (
	"testing"

	r "github.com/mohae/joefriday/platform/release"
)

func TestSerializeDeserialize(t *testing.T) {
	p, err := Get()
	if err != nil {
		t.Errorf("Get(): got %s, want nil", err)
		return
	}
	inf, err := r.Get()
	if err != nil {
		t.Errorf("release.Get(): got %s, want nil", err)
		return
	}
	infD := Deserialize(p)
	if inf.Name != infD.Name {
		t.Errorf("Name: got %s; want %s", infD.Name, inf.Name)
	}
	if inf.ID != infD.ID {
		t.Errorf("ID: got %s; want %s", infD.ID, inf.ID)
	}
	if inf.IDLike != infD.IDLike {
		t.Errorf("IDLike: got %s; want %s", infD.IDLike, inf.IDLike)
	}
	if inf.PrettyName != infD.PrettyName {
		t.Errorf("PrettyName: got %s; want %s", infD.PrettyName, inf.PrettyName)
	}
	if inf.Version != infD.Version {
		t.Errorf("Version: got %s; want %s", infD.Version, inf.Version)
	}
	if inf.VersionID != infD.VersionID {
		t.Errorf("VersionID: got %s; want %s", infD.VersionID, inf.VersionID)
	}
	if inf.HomeURL != infD.HomeURL {
		t.Errorf("HomeURL: got %s; want %s", infD.HomeURL, inf.HomeURL)
	}
	if inf.BugReportURL != infD.BugReportURL {
		t.Errorf("BugReportURL: got %s; want %s", infD.BugReportURL, inf.BugReportURL)
	}
}

func BenchmarkGet(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := NewProfiler()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = p.Get()
	}
	_ = tmp
}

func BenchmarkSerialize(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := NewProfiler()
	k, _ := p.Profiler.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = Serialize(k)
	}
	_ = tmp
}

func BenchmarkDeserialize(b *testing.B) {
	var inf *r.Info
	b.StopTimer()
	p, _ := NewProfiler()
	tmp, _ := p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		inf = Deserialize(tmp)
	}
	_ = inf
}