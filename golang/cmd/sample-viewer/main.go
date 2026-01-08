package main

import (
	"encoding/json"
	"fmt"
	"log"
	"serialization-benchmark/internal/config"
	"serialization-benchmark/internal/model"

	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/proto"
)

const (
	configPath = "benchmark-config.json"
)

func main() {
	fmt.Println("=== Sample Data Viewer ===")

	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Printf("Warning: Could not load config from %s: %v\n", configPath, err)
		log.Println("Using default configuration...")
		cfg = config.DefaultConfig()
	}

	// Generate 3 sample records with configuration
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
	records := gen.GenerateRecords(3)

	for i, record := range records {
		fmt.Printf("--- Record %d ---\n", i+1)
		fmt.Printf("ID: %s\n", record.ID)
		fmt.Printf("Timestamp: %v\n\n", record.Timestamp)

		// Primitive numeric types
		fmt.Println("Primitive Numeric Types:")
		fmt.Printf("  bool_value:  %v\n", record.BoolValue)
		fmt.Printf("  byte_value:  %d\n", record.ByteValue)
		fmt.Printf("  short_value: %d\n", record.ShortValue)
		fmt.Printf("  int_value:   %d\n", record.IntValue)
		fmt.Printf("  long_value:  %d\n", record.LongValue)

		// Floating point types
		fmt.Println("\nFloating Point Types:")
		fmt.Printf("  float_value:   %.2f\n", record.FloatValue)
		fmt.Printf("  double_value:  %.2f\n", record.DoubleValue)
		fmt.Printf("  decimal_value: %s\n", record.DecimalValue)

		// String and binary types
		fmt.Println("\nString and Binary Types:")
		fmt.Printf("  string_value: \"%s\"\n", record.StringValue)
		fmt.Printf("  binary_value: [%d bytes] %x...\n", len(record.BinaryValue), record.BinaryValue[:min(8, len(record.BinaryValue))])

		// Date and time types
		fmt.Println("\nDate and Time Types:")
		fmt.Printf("  date_value:          %d (days since epoch)\n", record.DateValue)
		fmt.Printf("  timestamp_value:     %d (microseconds)\n", record.TimestampValue)
		fmt.Printf("  timestamp_ntz_value: %d (microseconds)\n", record.TimestampNtzValue)

		// Complex types
		fmt.Println("\nComplex Types:")
		if record.ArrayValue != nil {
			fmt.Printf("  array_value: [%d elements]\n", len(record.ArrayValue.Elements))
			for j, elem := range record.ArrayValue.Elements {
				fmt.Printf("    [%d]: \"%s\"\n", j, elem)
			}
		}

		if record.MapValue != nil {
			fmt.Printf("  map_value: {%d entries}\n", len(record.MapValue.Entries))
			for k, v := range record.MapValue.Entries {
				fmt.Printf("    \"%s\": \"%s\"\n", k, v)
			}
		}

		if record.StructValue != nil {
			fmt.Printf("  struct_value: {%d fields}\n", len(record.StructValue.Fields))
			for k, v := range record.StructValue.Fields {
				fmt.Printf("    %s: \"%s\"\n", k, v)
			}
		}

		// Show serialized sizes
		fmt.Println("\nSerialized Sizes:")
		showSizes(record)

		fmt.Println()
	}
}

func showSizes(record *model.FeatureRecord) {
	// Serialize with msgpack
	msgpackData, err := msgpack.Marshal(record)
	if err != nil {
		fmt.Printf("  MessagePack: error - %v\n", err)
	} else {
		fmt.Printf("  MessagePack: %d bytes\n", len(msgpackData))
	}

	// Serialize with protobuf
	pbRecord := model.ConvertToProto(record)
	protoData, err := proto.Marshal(pbRecord)
	if err != nil {
		fmt.Printf("  Protobuf: error - %v\n", err)
	} else {
		fmt.Printf("  Protobuf: %d bytes\n", len(protoData))
	}

	// JSON for comparison (compact)
	jsonData, err := json.Marshal(record)
	if err != nil {
		fmt.Printf("  JSON (compact): error - %v\n", err)
	} else {
		fmt.Printf("  JSON (compact): %d bytes\n", len(jsonData))
	}

	// Calculate compression ratios
	if len(msgpackData) > 0 && len(protoData) > 0 {
		ratio := float64(len(msgpackData)) / float64(len(protoData))
		fmt.Printf("  Protobuf is %.2fx smaller than MessagePack\n", ratio)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
