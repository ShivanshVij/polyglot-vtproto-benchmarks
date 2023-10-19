/*
	Copyright 2023 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package demo

import (
	polyglotBenchmark "benchmark/polyglot/benchmark"
	vtBenchmark "benchmark/vtproto/benchmark"
	"bytes"
	"crypto/rand"
	"github.com/loopholelabs/polyglot"
	"github.com/loov/hrtime/hrtesting"
	"math"
	"runtime"
	"testing"
)

func BenchmarkEncodeBytes(b *testing.B) {
	randData := make([]byte, 512)
	_, _ = rand.Read(randData)

	polyglotData := polyglotBenchmark.BytesData{
		Bytes: randData,
	}
	polyglotBuf := polyglot.NewBuffer()

	vtData := vtBenchmark.BytesData{
		Bytes: randData,
	}
	vtBuf := make([]byte, 0, 1024)

	b.ResetTimer()

	b.Run("polyglot", func(b *testing.B) {
		b.SetBytes(512)
		bench := hrtesting.NewBenchmark(b)
		for bench.Next() {
			polyglotData.Encode(polyglotBuf)
			polyglotBuf.Reset()
		}
		runtime.KeepAlive(polyglotData)
	})

	b.Run("vtproto", func(b *testing.B) {
		var err error
		b.SetBytes(512)
		bench := hrtesting.NewBenchmark(b)
		for bench.Next() {
			_, err = vtData.MarshalToVT(vtBuf)
			if err != nil {
				b.Fatal(err)
			}
			vtBuf = vtBuf[:0]
		}
		runtime.KeepAlive(vtData)
	})
}

func BenchmarkEncodeBytesParallel(b *testing.B) {
	randData := make([]byte, 512)
	_, _ = rand.Read(randData)

	polyglotData := polyglotBenchmark.BytesData{
		Bytes: randData,
	}

	b.Run("polyglot", func(b *testing.B) {
		b.SetBytes(512)
		b.RunParallel(func(pb *testing.PB) {
			polyglotBuf := polyglot.NewBuffer()
			for pb.Next() {
				polyglotData.Encode(polyglotBuf)
				polyglotBuf.Reset()
			}
			runtime.KeepAlive(polyglotData)
		})
	})

	vtData := vtBenchmark.BytesData{
		Bytes: randData,
	}

	b.Run("vtproto", func(b *testing.B) {
		b.SetBytes(512)
		b.RunParallel(func(pb *testing.PB) {
			var err error
			vtBuf := make([]byte, 0, 1024)
			for pb.Next() {
				_, err = vtData.MarshalToVT(vtBuf)
				if err != nil {
					b.Fatal(err)
				}
				vtBuf = vtBuf[:0]
			}
			runtime.KeepAlive(vtData)
		})
	})
}

func BenchmarkDecodeBytes(b *testing.B) {
	randData := make([]byte, 512)
	_, _ = rand.Read(randData)

	polyglotData := polyglotBenchmark.BytesData{
		Bytes: randData,
	}
	polyglotBuf := polyglot.NewBuffer()
	polyglotData.Encode(polyglotBuf)
	polyglotBytes := polyglotBuf.Bytes()

	vtData := vtBenchmark.BytesData{
		Bytes: randData,
	}

	vtBuf := make([]byte, 1024)
	_, _ = vtData.MarshalToVT(vtBuf)
	vtBuf = vtBuf[:vtData.SizeVT()]

	b.ResetTimer()

	b.Run("polyglot", func(b *testing.B) {
		var err error
		b.SetBytes(512)
		bench := hrtesting.NewBenchmark(b)
		for bench.Next() {
			polyglotData.Bytes = nil
			err = polyglotData.Decode(polyglotBytes)
			if err != nil {
				b.Fatal(err)
			}
			if !bytes.Equal(polyglotData.Bytes, randData) {
				b.Fail()
			}
		}
		runtime.KeepAlive(polyglotData)
	})

	b.Run("vtproto", func(b *testing.B) {
		var err error
		b.SetBytes(512)
		bench := hrtesting.NewBenchmark(b)
		for bench.Next() {
			vtData.Reset()
			err = vtData.UnmarshalVT(vtBuf)
			if err != nil {
				b.Fatal(err)
			}
			if !bytes.Equal(vtData.Bytes, randData) {
				b.Fail()
			}
		}
		runtime.KeepAlive(vtData)
	})
}

func BenchmarkDecodeBytesParallel(b *testing.B) {
	randData := make([]byte, 512)
	_, _ = rand.Read(randData)

	polyglotData := polyglotBenchmark.BytesData{
		Bytes: randData,
	}
	polyglotBuf := polyglot.NewBuffer()
	polyglotData.Encode(polyglotBuf)
	polyglotBytes := polyglotBuf.Bytes()

	vtData := vtBenchmark.BytesData{
		Bytes: randData,
	}

	vtBuf := make([]byte, 1024)
	_, _ = vtData.MarshalToVT(vtBuf)
	vtBuf = vtBuf[:vtData.SizeVT()]

	b.ResetTimer()

	b.Run("polyglot", func(b *testing.B) {
		b.SetBytes(512)
		b.RunParallel(func(pb *testing.PB) {
			var err error
			polyglotData := new(polyglotBenchmark.BytesData)
			for pb.Next() {
				polyglotData.Bytes = nil
				err = polyglotData.Decode(polyglotBytes)
				if err != nil {
					b.Fatal(err)
				}
				if !bytes.Equal(polyglotData.Bytes, randData) {
					b.Fail()
				}
			}
			runtime.KeepAlive(polyglotData)
		})
	})

	b.Run("vtproto", func(b *testing.B) {
		b.SetBytes(512)
		b.RunParallel(func(pb *testing.PB) {
			var err error
			vtData := new(vtBenchmark.BytesData)
			for pb.Next() {
				vtData.Bytes = nil
				err = vtData.UnmarshalVT(vtBuf)
				if err != nil {
					b.Fatal(err)
				}
				if !bytes.Equal(vtData.Bytes, randData) {
					b.Fail()
				}
			}
			runtime.KeepAlive(vtData)
		})
	})
}

func BenchmarkEncodeInt32(b *testing.B) {
	polyglotData := polyglotBenchmark.IntegerData{
		Integer: math.MaxInt32,
	}
	polyglotBuf := polyglot.NewBuffer()

	vtData := vtBenchmark.IntegerData{
		Integer: math.MaxInt32,
	}
	vtBuf := make([]byte, 0, 1024)

	b.ResetTimer()

	b.Run("polyglot", func(b *testing.B) {
		bench := hrtesting.NewBenchmark(b)
		for bench.Next() {
			polyglotData.Encode(polyglotBuf)
			polyglotBuf.Reset()
		}
		runtime.KeepAlive(polyglotData)
	})

	b.Run("vtproto", func(b *testing.B) {
		var err error
		bench := hrtesting.NewBenchmark(b)
		for bench.Next() {
			_, err = vtData.MarshalToVT(vtBuf)
			if err != nil {
				b.Fatal(err)
			}
			vtBuf = vtBuf[:0]
		}
		runtime.KeepAlive(vtData)
	})
}

func BenchmarkDecodeInt32(b *testing.B) {
	polyglotData := polyglotBenchmark.IntegerData{
		Integer: math.MaxInt32,
	}
	polyglotBuf := polyglot.NewBuffer()
	polyglotData.Encode(polyglotBuf)
	polyglotBytes := polyglotBuf.Bytes()

	vtData := vtBenchmark.IntegerData{
		Integer: math.MaxInt32,
	}
	vtBuf := make([]byte, 1024)
	_, _ = vtData.MarshalToVT(vtBuf)
	vtBuf = vtBuf[:vtData.SizeVT()]

	b.ResetTimer()

	b.Run("polyglot", func(b *testing.B) {
		var err error
		bench := hrtesting.NewBenchmark(b)
		for bench.Next() {
			err = polyglotData.Decode(polyglotBytes)
			if err != nil {
				b.Fatal(err)
			}
			if polyglotData.Integer != math.MaxInt32 {
				b.Fail()
			}
		}
		runtime.KeepAlive(polyglotData)
	})

	b.Run("vtproto", func(b *testing.B) {
		var err error
		bench := hrtesting.NewBenchmark(b)
		for bench.Next() {
			err = vtData.UnmarshalVT(vtBuf)
			if err != nil {
				b.Fatal(err)
			}
			if vtData.Integer != math.MaxInt32 {
				b.Fail()
			}
		}
		runtime.KeepAlive(vtData)
	})
}
