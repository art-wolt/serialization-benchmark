package main

import (
	"fmt"
	"log"
	"serialization-benchmark/internal/config"
	"serialization-benchmark/internal/model"
	pb "serialization-benchmark/proto"
	"sort"
	"time"

	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/proto"
)

const (
	configPath = "benchmark-config.json"
)

type BenchmarkResult struct {
	Algorithm            string
	NumRecords           int
	NumRuns              int
	SerializationTimes   []time.Duration
	DeserializationTimes []time.Duration
	SerializationP50     time.Duration
	SerializationP99     time.Duration
	DeserializationP50   time.Duration
	DeserializationP99   time.Duration
	TotalSize            int64
	AvgSizePerRecord     int64
}

func main() {
	fmt.Println("Starting serialization benchmark...")

	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Printf("Warning: Could not load config from %s: %v", configPath, err)
		log.Println("Using default configuration...")
		cfg = config.DefaultConfig()
	}

	fmt.Printf("Configuration:\n")
	fmt.Printf("  Records: %d\n", cfg.NumRecords)
	fmt.Printf("  Iterations: %d\n", cfg.NumIterations)
	fmt.Printf("\nGenerating %d records...\n", cfg.NumRecords)

	// Generate test data once with configuration
	gen := model.NewGenerator(
		cfg.DataGeneration.StringLength,
		cfg.DataGeneration.BinarySize,
		cfg.DataGeneration.ArrayMinSize,
		cfg.DataGeneration.ArrayMaxSize,
		cfg.DataGeneration.MapMinSize,
		cfg.DataGeneration.MapMaxSize,
		cfg.DataGeneration.StructMinSize,
		cfg.DataGeneration.StructMaxSize,
	)
	records := gen.GenerateRecords(cfg.NumRecords)
	fmt.Printf("Generated %d records\n", len(records))
	fmt.Printf("Running %d iterations for each algorithm...\n\n", cfg.NumIterations)

	// Run protobuf benchmark
	fmt.Println("Running Protobuf benchmark...")
	protobufResult := benchmarkProtobuf(records, cfg.NumIterations)
	printResult(protobufResult)

	// Run msgpack benchmark
	fmt.Println("\nRunning MessagePack benchmark...")
	msgpackResult := benchmarkMsgpack(records, cfg.NumIterations)
	printResult(msgpackResult)

	// Print comparison
	fmt.Println("\n=== Comparison (p50) ===")
	fmt.Printf("Serialization: Protobuf %.2fx %s than MessagePack\n",
		float64(msgpackResult.SerializationP50)/float64(protobufResult.SerializationP50),
		comparison(protobufResult.SerializationP50, msgpackResult.SerializationP50))

	fmt.Printf("Deserialization: Protobuf %.2fx %s than MessagePack\n",
		float64(msgpackResult.DeserializationP50)/float64(protobufResult.DeserializationP50),
		comparison(protobufResult.DeserializationP50, msgpackResult.DeserializationP50))

	fmt.Printf("Size: Protobuf %.2fx %s than MessagePack\n",
		float64(msgpackResult.TotalSize)/float64(protobufResult.TotalSize),
		comparison(protobufResult.TotalSize, msgpackResult.TotalSize))
}

func benchmarkProtobuf(records []*model.FeatureRecord, runs int) BenchmarkResult {
	result := BenchmarkResult{
		Algorithm:            "Protobuf",
		NumRecords:           len(records),
		NumRuns:              runs,
		SerializationTimes:   make([]time.Duration, runs),
		DeserializationTimes: make([]time.Duration, runs),
	}

	// Convert to protobuf format once
	pbRecords := make([]*pb.FeatureRecord, len(records))
	for i, record := range records {
		pbRecords[i] = model.ConvertToProto(record)
	}

	// Run multiple iterations
	var serializedData [][]byte
	for run := 0; run < runs; run++ {
		// Benchmark serialization
		serializedData = make([][]byte, len(pbRecords))
		var totalSize int64
		start := time.Now()
		for i, record := range pbRecords {
			data, err := proto.Marshal(record)
			if err != nil {
				log.Fatalf("Failed to serialize protobuf record %d: %v", i, err)
			}
			serializedData[i] = data
			totalSize += int64(len(data))
		}
		result.SerializationTimes[run] = time.Since(start)

		// Store size from first run
		if run == 0 {
			result.TotalSize = totalSize
		}

		// Benchmark deserialization
		start = time.Now()
		for i, data := range serializedData {
			record := &pb.FeatureRecord{}
			if err := proto.Unmarshal(data, record); err != nil {
				log.Fatalf("Failed to deserialize protobuf record %d: %v", i, err)
			}
		}
		result.DeserializationTimes[run] = time.Since(start)

		fmt.Printf("  Run %d/%d complete\n", run+1, runs)
	}

	// Calculate percentiles
	result.SerializationP50 = calculatePercentile(result.SerializationTimes, 50)
	result.SerializationP99 = calculatePercentile(result.SerializationTimes, 99)
	result.DeserializationP50 = calculatePercentile(result.DeserializationTimes, 50)
	result.DeserializationP99 = calculatePercentile(result.DeserializationTimes, 99)
	result.AvgSizePerRecord = result.TotalSize / int64(len(records))

	return result
}

func benchmarkMsgpack(records []*model.FeatureRecord, runs int) BenchmarkResult {
	result := BenchmarkResult{
		Algorithm:            "MessagePack",
		NumRecords:           len(records),
		NumRuns:              runs,
		SerializationTimes:   make([]time.Duration, runs),
		DeserializationTimes: make([]time.Duration, runs),
	}

	// Run multiple iterations
	var serializedData [][]byte
	for run := 0; run < runs; run++ {
		// Benchmark serialization
		serializedData = make([][]byte, len(records))
		var totalSize int64
		start := time.Now()
		for i, record := range records {
			data, err := msgpack.Marshal(record)
			if err != nil {
				log.Fatalf("Failed to serialize msgpack record %d: %v", i, err)
			}
			serializedData[i] = data
			totalSize += int64(len(data))
		}
		result.SerializationTimes[run] = time.Since(start)

		// Store size from first run
		if run == 0 {
			result.TotalSize = totalSize
		}

		// Benchmark deserialization
		start = time.Now()
		for i, data := range serializedData {
			record := &model.FeatureRecord{}
			if err := msgpack.Unmarshal(data, record); err != nil {
				log.Fatalf("Failed to deserialize msgpack record %d: %v", i, err)
			}
		}
		result.DeserializationTimes[run] = time.Since(start)

		fmt.Printf("  Run %d/%d complete\n", run+1, runs)
	}

	// Calculate percentiles
	result.SerializationP50 = calculatePercentile(result.SerializationTimes, 50)
	result.SerializationP99 = calculatePercentile(result.SerializationTimes, 99)
	result.DeserializationP50 = calculatePercentile(result.DeserializationTimes, 50)
	result.DeserializationP99 = calculatePercentile(result.DeserializationTimes, 99)
	result.AvgSizePerRecord = result.TotalSize / int64(len(records))

	return result
}

func calculatePercentile(durations []time.Duration, percentile int) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	// Create a copy and sort
	sorted := make([]time.Duration, len(durations))
	copy(sorted, durations)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	// Calculate index for percentile
	index := int(float64(len(sorted)-1) * float64(percentile) / 100.0)
	if index < 0 {
		index = 0
	}
	if index >= len(sorted) {
		index = len(sorted) - 1
	}

	return sorted[index]
}

func printResult(result BenchmarkResult) {
	fmt.Printf("Algorithm: %s\n", result.Algorithm)
	fmt.Printf("  Records: %d\n", result.NumRecords)
	fmt.Printf("  Runs: %d\n", result.NumRuns)
	fmt.Printf("  Serialization:\n")
	fmt.Printf("    p50: %v (%.2f ns/op)\n",
		result.SerializationP50,
		float64(result.SerializationP50.Nanoseconds())/float64(result.NumRecords))
	fmt.Printf("    p99: %v (%.2f ns/op)\n",
		result.SerializationP99,
		float64(result.SerializationP99.Nanoseconds())/float64(result.NumRecords))
	fmt.Printf("  Deserialization:\n")
	fmt.Printf("    p50: %v (%.2f ns/op)\n",
		result.DeserializationP50,
		float64(result.DeserializationP50.Nanoseconds())/float64(result.NumRecords))
	fmt.Printf("    p99: %v (%.2f ns/op)\n",
		result.DeserializationP99,
		float64(result.DeserializationP99.Nanoseconds())/float64(result.NumRecords))
	fmt.Printf("  Total Size: %.2f MB\n", float64(result.TotalSize)/(1024*1024))
	fmt.Printf("  Avg Size per Record: %d bytes\n", result.AvgSizePerRecord)
}

func comparison(a, b interface{}) string {
	switch v := a.(type) {
	case time.Duration:
		if v < b.(time.Duration) {
			return "faster"
		}
		return "slower"
	case int64:
		if v < b.(int64) {
			return "smaller"
		}
		return "larger"
	}
	return "unknown"
}
