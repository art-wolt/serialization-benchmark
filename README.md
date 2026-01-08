# Serialization Performance Benchmark

A performance testing repository for comparing different serialization algorithms across multiple programming languages.

## Overview

This repository contains performance benchmarks for serialization algorithms:
- **Protobuf** (Protocol Buffers)
- **MessagePack** (msgpack)

Supported languages:
- Golang ✅
- Java ✅

## Data Model

Each `FeatureRecord` represents a **table row** with all 16 data types populated in every record:

- **Primitive numeric types**: Boolean, Byte, Short, Integer, Long
- **Floating point types**: Float, Double, Decimal
- **String and binary types**: String (20 chars), Binary (32 bytes)
- **Date and time types**: Date, Timestamp, Timestamp_NTZ
- **Complex types**: Array (1-5 string elements), Map (1-3 entries), Struct (1-3 fields)

Each record contains ~306-512 bytes of data depending on the serialization format, making it representative of real-world database rows with mixed data types.

## Configuration

The benchmark behavior is controlled by `benchmark-config.json` in the repository root. This configuration is shared across all language implementations.

```json
{
  "num_records": 1000000,
  "num_iterations": 10,
  "data": {
    "string_length": 20,
    "binary_size": 32,
    "array_min_size": 1,
    "array_max_size": 5,
    "map_min_size": 1,
    "map_max_size": 3,
    "struct_min_size": 1,
    "struct_max_size": 3
  }
}
```

**Configuration options:**
- `num_records`: Number of records to generate for benchmarking (default: 1,000,000)
- `num_iterations`: Number of iterations to run for calculating p50/p99 (default: 10)
- `data`: Data generation settings for string length, binary size, and complex type sizes

## Golang Implementation

### Prerequisites

- Go 1.19+
- Bazel 8.4.2+ (recommended)
- Protocol Buffers compiler (`protoc`) - only needed for regenerating proto files

### Run with Bazel

```bash
# Run the benchmark (tests 1M records by default)
bazel run //golang/cmd/benchmark:benchmark

# View sample generated data (generates and displays 3 sample records)
bazel run //golang/cmd/sample-viewer:sample-viewer
```

### Project Structure

```
golang/
├── cmd/
│   ├── benchmark/        # Main benchmark executable
│   │   └── main.go
│   └── sample-viewer/    # Sample data viewer utility
│       └── main.go
├── internal/
│   └── model/           # Data models and utilities
│       ├── model.go     # Go structs (all 16 data types)
│       ├── generator.go # Random data generator
│       └── converter.go # Protobuf/MessagePack converter
├── proto/               # Generated protobuf code
│   └── feature.pb.go
└── go.mod

proto/                   # Protobuf schema definitions
└── feature.proto
```

### Inspecting Generated Data

To understand what data is being generated for the benchmark, use the sample viewer:

```bash
bazel run //golang/cmd/sample-viewer:sample-viewer
```

This generates 3 sample records and displays **all fields** for each record:
- Record ID and timestamp
- All primitive numeric types (bool, byte, short, int, long)
- All floating point types (float, double, decimal)
- String and binary values
- Date and time values
- Complex types (array, map, struct) with their contents
- Serialized sizes in MessagePack, Protobuf, and JSON formats

Example output (showing one complete record):
```
--- Record 1 ---
ID: record_0
Timestamp: 2026-01-08 12:37:55.618027 +0100 CET

Primitive Numeric Types:
  bool_value:  true
  byte_value:  -56
  short_value: -30930
  int_value:   386352766
  long_value:  3756403615765010272

Floating Point Types:
  float_value:   8709.61
  double_value:  1240.70
  decimal_value: 1143.93

String and Binary Types:
  string_value: "0z4MvEFmgrMQoGQHT1To"
  binary_value: [32 bytes] 13eb2779ac0a5ebb...

Date and Time Types:
  date_value:          1171 (days since epoch)
  timestamp_value:     1767872275618085 (microseconds)
  timestamp_ntz_value: 1767872275618085 (microseconds)

Complex Types:
  array_value: [4 elements]
    [0]: "ymM1iHmoQg"
    [1]: "z4xqSPOzKV"
    ...
  map_value: {2 entries}
    "xqc2sF9h": "bBW7TNUkPUJc"
    "rvCY9cBC": "t0uqGGn0i02x"
  struct_value: {3 fields}
    rK8iqBCO: "MMbXoiH9WNnb"
    ...

Serialized Sizes:
  MessagePack: 541 bytes
  Protobuf: 346 bytes
  JSON (compact): 698 bytes
  Protobuf is 1.56x smaller than MessagePack
```

This shows how each record is a complete table row with all 16 data types populated.

### Regenerating Protobuf Code

If you modify the protobuf schema:

```bash
cd golang
PATH=$PATH:~/go/bin protoc --go_out=. --go_opt=paths=source_relative \
  --proto_path=../proto ../proto/feature.proto
mv feature.pb.go proto/
```

## Benchmark Methodology

For each serialization algorithm, the benchmark:

1. **Generates 1M random records** with diverse data types
2. **Runs 10 iterations** for each algorithm
3. **Measures serialization time** - time to convert all records to bytes
4. **Measures deserialization time** - time to convert bytes back to objects
5. **Calculates percentiles** - p50 (median) and p99 for statistical accuracy
6. **Calculates total size** - compressed size of serialized data
7. **Reports comparative metrics** - speed and size comparisons

Times are measured separately and percentiles are calculated to provide reliable performance insights across multiple runs.

## Sample Output

```
Starting serialization benchmark...
Generating 1000000 records...
Generated 1000000 records
Running 10 iterations for each algorithm...

Running Protobuf benchmark...
  Run 1/10 complete
  ...
  Run 10/10 complete
Algorithm: Protobuf
  Records: 1000000
  Runs: 10
  Serialization:
    p50: 1.627380042s (1627.38 ns/op)
    p99: 1.692971125s (1692.97 ns/op)
  Deserialization:
    p50: 1.764898792s (1764.90 ns/op)
    p99: 1.808452584s (1808.45 ns/op)
  Total Size: 291.86 MB
  Avg Size per Record: 306 bytes

Running MessagePack benchmark...
  Run 1/10 complete
  ...
  Run 10/10 complete
Algorithm: MessagePack
  Records: 1000000
  Runs: 10
  Serialization:
    p50: 1.571492416s (1571.49 ns/op)
    p99: 1.603961459s (1603.96 ns/op)
  Deserialization:
    p50: 2.04694825s (2046.95 ns/op)
    p99: 2.121021125s (2121.02 ns/op)
  Total Size: 489.01 MB
  Avg Size per Record: 512 bytes

=== Comparison (p50) ===
Serialization: Protobuf 0.97x slower than MessagePack
Deserialization: Protobuf 1.16x faster than MessagePack
Size: Protobuf 1.68x smaller than MessagePack
```

**Key Results (based on p50):**
- **Record Size**: Each record contains all 16 data types (306 bytes in Protobuf, 512 bytes in MessagePack)
- **Serialization**: MessagePack is slightly faster (1571ms vs 1627ms) - 3.5% faster
- **Deserialization**: Protobuf is 1.16x faster than MessagePack (1765ms vs 2047ms) - 16% faster
- **Size Efficiency**: Protobuf produces files 1.68x smaller than MessagePack (292MB vs 489MB) - 40% smaller
- **Consistency**: Both algorithms show excellent performance stability (low variance between p50 and p99)
- **Throughput**: ~614K records/second (Protobuf ser.), ~636K records/second (MessagePack ser.), ~567K records/second (Protobuf deser.), ~489K records/second (MessagePack deser.)

## Java Implementation

### Prerequisites

- Java 17+
- Bazel 8.4.2+
- Protocol Buffers compiler (`protoc`) - only needed for regenerating proto files

### Run with Bazel

```bash
# Run the benchmark (tests 1M records by default)
bazel run //java:benchmark
```

### Project Structure

```
java/
├── src/main/java/com/benchmark/
│   ├── Benchmark.java         # Main benchmark executable
│   ├── config/
│   │   └── BenchmarkConfig.java  # Config reader
│   ├── model/
│   │   ├── FeatureRecord.java    # Java POJOs (all 16 data types)
│   │   ├── RecordGenerator.java  # Random data generator
│   │   └── ProtoConverter.java   # Protobuf converter
│   └── proto/
│       └── Feature.java          # Generated protobuf code
└── BUILD.bazel
```

### Regenerating Protobuf Code

If you modify the protobuf schema:

```bash
protoc --java_out=java/src/main/java --proto_path=proto proto/feature.proto
```

The generated Java protobuf code will be placed in `java/src/main/java/com/benchmark/proto/Feature.java`.

### Sample Output (Java)

```
Starting serialization benchmark...
Loaded configuration from benchmark-config.json

Configuration:
  Records: 1000000
  Iterations: 10

Generating 1000000 records...
Generated 1000000 records
Running 10 iterations for each algorithm...

Running Protobuf benchmark...
Algorithm: Protobuf
  Records: 1000000
  Runs: 10
  Serialization:
    p50: 0.50s (504.92 ns/op)
    p99: 0.97s (971.25 ns/op)
  Deserialization:
    p50: 1.02s (1016.60 ns/op)
    p99: 1.18s (1178.59 ns/op)
  Total Size: 292.44 MB
  Avg Size per Record: 306 bytes

Running MessagePack benchmark...
Algorithm: MessagePack
  Records: 1000000
  Runs: 10
  Serialization:
    p50: 1.48s (1479.54 ns/op)
    p99: 1.66s (1657.71 ns/op)
  Deserialization:
    p50: 2.53s (2527.23 ns/op)
    p99: 2.79s (2785.80 ns/op)
  Total Size: 469.34 MB
  Avg Size per Record: 492 bytes

=== Comparison (p50) ===
Serialization: Protobuf 2.93x faster than MessagePack
Deserialization: Protobuf 2.49x faster than MessagePack
Size: Protobuf 1.60x smaller than MessagePack
```

**Key Results (Java, based on p50):**
- **Serialization**: Protobuf is 2.93x faster than MessagePack (505ms vs 1480ms)
- **Deserialization**: Protobuf is 2.49x faster than MessagePack (1017ms vs 2527ms)
- **Size Efficiency**: Protobuf produces files 1.60x smaller than MessagePack (292MB vs 469MB)
- **Throughput**: ~1.98M records/second (Protobuf ser.), ~676K records/second (MessagePack ser.), ~983K records/second (Protobuf deser.), ~396K records/second (MessagePack deser.)

## License

MIT
