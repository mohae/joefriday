package main

import (
	"testing"

	"github.com/mohae/benchutil"
	"github.com/mohae/joefriday/disk/stats"
	sfb "github.com/mohae/joefriday/disk/stats/flat"
	sjson "github.com/mohae/joefriday/disk/stats/json"
	"github.com/mohae/joefriday/disk/structs"
	"github.com/mohae/joefriday/disk/usage"
	ufb "github.com/mohae/joefriday/disk/usage/flat"
	ujson "github.com/mohae/joefriday/disk/usage/json"
)

const (
	DiskStats = "Disk Stats"
	DiskUsage = "Disk Usage"
)

func BenchDiskGetStats(b *testing.B) {
	var stts *structs.Stats
	b.StopTimer()
	p, _ := stats.NewProfiler()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		stts, _ = p.Get()
	}
	_ = stts
}

func DiskGetStats() benchutil.Bench {
	bench := benchutil.NewBench("Get")
	bench.Group = DiskStats
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskGetStats))
	return bench
}

func BenchDiskGetStatsFB(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := sfb.NewProfiler()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = p.Get()
	}
	_ = tmp
}

func DiskGetStatsFB() benchutil.Bench {
	bench := benchutil.NewBench("flat.Get")
	bench.Group = DiskStats
	bench.Desc = Flat
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskGetStatsFB))
	return bench
}

func BenchDiskStatsSerializeFB(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := stats.NewProfiler()
	sts, _ := p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = sfb.Serialize(sts)
	}
	_ = tmp
}

func DiskStatsSerializeFB() benchutil.Bench {
	bench := benchutil.NewBench("flat.Serialize")
	bench.Group = DiskStats
	bench.Desc = Flat
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskStatsSerializeFB))
	return bench
}

func BenchDiskStatsDeserializeFB(b *testing.B) {
	var sts *structs.Stats
	b.StopTimer()
	p, _ := sfb.NewProfiler()
	tmp, _ := p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sts = sfb.Deserialize(tmp)
	}
	_ = sts
}

func DiskStatsDeserializeFB() benchutil.Bench {
	bench := benchutil.NewBench("flat.Deserialize")
	bench.Group = DiskStats
	bench.Desc = Flat
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskStatsDeserializeFB))
	return bench
}

func BenchDiskStatsGetJSON(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := sjson.NewProfiler()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = p.Get()
	}
	_ = tmp
}

func DiskGetStatsJSON() benchutil.Bench {
	bench := benchutil.NewBench("json.Get")
	bench.Group = DiskStats
	bench.Desc = JSON
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskStatsGetJSON))
	return bench
}

func BenchDiskStatsSerializeJSON(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := stats.NewProfiler()
	sts, _ := p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = sjson.Serialize(sts)
	}
	_ = tmp
}

func DiskStatsSerializeJSON() benchutil.Bench {
	bench := benchutil.NewBench("json.Serialize")
	bench.Group = DiskStats
	bench.Desc = JSON
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskStatsSerializeJSON))
	return bench
}

func BenchDiskStatsDeserializeJSON(b *testing.B) {
	var sts *structs.Stats
	b.StopTimer()
	p, _ := sjson.NewProfiler()
	tmp, _ := p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sts, _ = sjson.Deserialize(tmp)
	}
	_ = sts
}

func DiskStatsDeserializeJSON() benchutil.Bench {
	bench := benchutil.NewBench("json.Deserialize")
	bench.Group = DiskStats
	bench.Desc = JSON
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskStatsDeserializeJSON))
	return bench
}

// Usage
func BenchDiskGetUsage(b *testing.B) {
	var u *structs.Usage
	b.StopTimer()
	p, _ := usage.NewProfiler()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		u, _ = p.Get()
	}
	_ = u
}

func DiskGetUsage() benchutil.Bench {
	bench := benchutil.NewBench("Get")
	bench.Group = DiskUsage
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskGetUsage))
	return bench
}

func BenchDiskGetUsageFB(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := sfb.NewProfiler()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = p.Get()
	}
	_ = tmp
}

func DiskGetUsageFB() benchutil.Bench {
	bench := benchutil.NewBench("flat.Get")
	bench.Group = DiskUsage
	bench.Desc = Flat
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskGetUsageFB))
	return bench
}

func BenchDiskUsageSerializeFB(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := usage.NewProfiler()
	u, _ := p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = ufb.Serialize(u)
	}
	_ = tmp
}

func DiskUsageSerializeFB() benchutil.Bench {
	bench := benchutil.NewBench("flat.Serialize")
	bench.Group = DiskUsage
	bench.Desc = Flat
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskUsageSerializeFB))
	return bench
}

func BenchDiskUsageDeserializeFB(b *testing.B) {
	var u *structs.Usage
	b.StopTimer()
	p, _ := ufb.NewProfiler()
	tmp, _ := p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		u = ufb.Deserialize(tmp)
	}
	_ = u
}

func DiskUsageDeserializeFB() benchutil.Bench {
	bench := benchutil.NewBench("flat.Deserialize")
	bench.Group = DiskUsage
	bench.Desc = Flat
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskUsageDeserializeFB))
	return bench
}

func BenchDiskUsageGetJSON(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := ujson.NewProfiler()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = p.Get()
	}
	_ = tmp
}

func DiskGetUsageJSON() benchutil.Bench {
	bench := benchutil.NewBench("json.Get")
	bench.Group = DiskUsage
	bench.Desc = JSON
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskUsageGetJSON))
	return bench
}

func BenchDiskUsageSerializeJSON(b *testing.B) {
	var tmp []byte
	b.StopTimer()
	p, _ := usage.NewProfiler()
	u, _ := p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tmp, _ = ujson.Serialize(u)
	}
	_ = tmp
}

func DiskUsageSerializeJSON() benchutil.Bench {
	bench := benchutil.NewBench("json.Serialize")
	bench.Group = DiskUsage
	bench.Desc = JSON
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskUsageSerializeJSON))
	return bench
}

func BenchDiskUsageDeserializeJSON(b *testing.B) {
	var u *structs.Usage
	b.StopTimer()
	p, _ := ujson.NewProfiler()
	tmp, _ := p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		u, _ = ujson.Deserialize(tmp)
	}
	_ = u
}

func DiskUsageDeserializeJSON() benchutil.Bench {
	bench := benchutil.NewBench("json.Deserialize")
	bench.Group = DiskUsage
	bench.Desc = JSON
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchDiskUsageDeserializeJSON))
	return bench
}
