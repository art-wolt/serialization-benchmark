package model

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// Generator creates random feature records for testing
type Generator struct {
	rnd            *randomSource
	stringLength   int
	binarySize     int
	arrayMinSize   int
	arrayMaxSize   int
	mapMinSize     int
	mapMaxSize     int
	structMinSize  int
	structMaxSize  int
}

type randomSource struct{}

func (r *randomSource) Int63n(n int64) int64 {
	max := big.NewInt(n)
	result, _ := rand.Int(rand.Reader, max)
	return result.Int64()
}

func (r *randomSource) Intn(n int) int {
	return int(r.Int63n(int64(n)))
}

func (r *randomSource) Float32() float32 {
	max := big.NewInt(1 << 24)
	result, _ := rand.Int(rand.Reader, max)
	return float32(result.Int64()) / float32(1<<24)
}

func (r *randomSource) Float64() float64 {
	max := big.NewInt(1 << 53)
	result, _ := rand.Int(rand.Reader, max)
	return float64(result.Int64()) / float64(1<<53)
}

func NewGenerator(stringLength, binarySize, arrayMinSize, arrayMaxSize, mapMinSize, mapMaxSize, structMinSize, structMaxSize int) *Generator {
	return &Generator{
		rnd:           &randomSource{},
		stringLength:  stringLength,
		binarySize:    binarySize,
		arrayMinSize:  arrayMinSize,
		arrayMaxSize:  arrayMaxSize,
		mapMinSize:    mapMinSize,
		mapMaxSize:    mapMaxSize,
		structMinSize: structMinSize,
		structMaxSize: structMaxSize,
	}
}

// NewGeneratorWithDefaults creates a generator with default settings
func NewGeneratorWithDefaults() *Generator {
	return NewGenerator(20, 32, 1, 5, 1, 3, 1, 3)
}

// GenerateRecord creates a single record with all data types populated
func (g *Generator) GenerateRecord(id string) *FeatureRecord {
	return &FeatureRecord{
		ID:        id,
		Timestamp: time.Now(),

		// Primitive numeric types
		BoolValue:  g.rnd.Intn(2) == 1,
		ByteValue:  int8(g.rnd.Intn(256) - 128),
		ShortValue: int16(g.rnd.Intn(65536) - 32768),
		IntValue:   int32(g.rnd.Int63n(1<<31) - 1<<30),
		LongValue:  g.rnd.Int63n(1 << 62),

		// Floating point types
		FloatValue:   g.rnd.Float32() * 10000,
		DoubleValue:  g.rnd.Float64() * 10000,
		DecimalValue: fmt.Sprintf("%.2f", g.rnd.Float64()*10000),

		// String and binary types
		StringValue: g.randomString(g.stringLength),
		BinaryValue: g.randomBytes(g.binarySize),

		// Date and time types
		DateValue:         int32(g.rnd.Intn(20000)), // Days since epoch
		TimestampValue:    time.Now().UnixMicro(),
		TimestampNtzValue: time.Now().UnixMicro(),

		// Complex types
		ArrayValue:  g.generateArrayValue(),
		MapValue:    g.generateMapValue(),
		StructValue: g.generateStructValue(),
	}
}

func (g *Generator) generateArrayValue() *ArrayValue {
	size := g.rnd.Intn(g.arrayMaxSize-g.arrayMinSize+1) + g.arrayMinSize
	elements := make([]string, size)
	for i := 0; i < size; i++ {
		elements[i] = g.randomString(10)
	}
	return &ArrayValue{Elements: elements}
}

func (g *Generator) generateMapValue() *MapValue {
	size := g.rnd.Intn(g.mapMaxSize-g.mapMinSize+1) + g.mapMinSize
	entries := make(map[string]string, size)
	for i := 0; i < size; i++ {
		entries[g.randomString(8)] = g.randomString(12)
	}
	return &MapValue{Entries: entries}
}

func (g *Generator) generateStructValue() *StructValue {
	size := g.rnd.Intn(g.structMaxSize-g.structMinSize+1) + g.structMinSize
	fields := make(map[string]string, size)
	for i := 0; i < size; i++ {
		fields[g.randomString(8)] = g.randomString(12)
	}
	return &StructValue{Fields: fields}
}

func (g *Generator) randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[g.rnd.Intn(len(charset))]
	}
	return string(b)
}

func (g *Generator) randomBytes(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return b
}

// GenerateRecords generates n records
func (g *Generator) GenerateRecords(n int) []*FeatureRecord {
	records := make([]*FeatureRecord, n)
	for i := 0; i < n; i++ {
		records[i] = g.GenerateRecord(fmt.Sprintf("record_%d", i))
	}
	return records
}
