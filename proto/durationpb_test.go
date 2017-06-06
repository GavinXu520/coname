// Code generated by protoc-gen-gogo.
// source: duration.proto
// DO NOT EDIT!

package proto

import testing "testing"
import math_rand "math/rand"
import time "time"
import github_com_maditya_protobuf_proto "github.com/maditya/protobuf/proto"
import github_com_maditya_protobuf_jsonpb "github.com/maditya/protobuf/jsonpb"
import fmt "fmt"
import go_parser "go/parser"
import proto1 "github.com/maditya/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestDurationProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedDuration(popr, false)
	data, err := github_com_maditya_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Duration{}
	if err := github_com_maditya_protobuf_proto.Unmarshal(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(data))
	copy(littlefuzz, data)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_maditya_protobuf_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestDurationMarshalTo(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedDuration(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Duration{}
	if err := github_com_maditya_protobuf_proto.Unmarshal(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func BenchmarkDurationProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*Duration, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedDuration(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_maditya_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkDurationProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_maditya_protobuf_proto.Marshal(NewPopulatedDuration(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &Duration{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_maditya_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestDurationJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedDuration(popr, true)
	marshaler := github_com_maditya_protobuf_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Duration{}
	err = github_com_maditya_protobuf_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestDurationProtoText(t *testing.T) {
t.Skip()
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedDuration(popr, true)
	data := github_com_maditya_protobuf_proto.MarshalTextString(p)
	msg := &Duration{}
	if err := github_com_maditya_protobuf_proto.UnmarshalText(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestDurationProtoCompactText(t *testing.T) {
t.Skip()
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedDuration(popr, true)
	data := github_com_maditya_protobuf_proto.CompactTextString(p)
	msg := &Duration{}
	if err := github_com_maditya_protobuf_proto.UnmarshalText(data, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("seed = %d, %#v !VerboseProto %#v, since %v", seed, msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestDurationVerboseEqual(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedDuration(popr, false)
	data, err := github_com_maditya_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Duration{}
	if err := github_com_maditya_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestDurationGoString(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedDuration(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestDurationSize(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedDuration(popr, true)
	size2 := github_com_maditya_protobuf_proto.Size(p)
	data, err := github_com_maditya_protobuf_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	size := p.Size()
	if len(data) != size {
		t.Errorf("seed = %d, size %v != marshalled size %v", seed, size, len(data))
	}
	if size2 != size {
		t.Errorf("seed = %d, size %v != before marshal proto.Size %v", seed, size, size2)
	}
	size3 := github_com_maditya_protobuf_proto.Size(p)
	if size3 != size {
		t.Errorf("seed = %d, size %v != after marshal proto.Size %v", seed, size, size3)
	}
}

func BenchmarkDurationSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*Duration, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedDuration(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestDurationStringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedDuration(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}

//These tests are generated by github.com/maditya/protobuf/plugin/testgen
