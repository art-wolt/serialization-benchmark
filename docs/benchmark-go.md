# Golang Benchmarks

This document covers the Golang implementation of the serialization benchmarks.

## Prerequisites

- Go 1.19+
- Bazel 8.4.2+ (recommended)
- Protocol Buffers compiler (`protoc`) - only needed for regenerating proto files

## Running the Benchmark

```bash
# Run the benchmark (tests 1M records by default)
bazel run //golang/cmd/benchmark:benchmark
```

## Project Structure

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

## Inspecting Generated Data

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

## Regenerating Protobuf Code

If you modify the protobuf schema:

```bash
cd golang
PATH=$PATH:~/go/bin protoc --go_out=. --go_opt=paths=source_relative \
  --proto_path=../proto ../proto/feature.proto
mv feature.pb.go proto/
```

## Next Steps

- See [Benchmark Scenarios](benchmark-scenarios.md) for methodology details
- See [Benchmark Summary](benchmark-summary.md)
