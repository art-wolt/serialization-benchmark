# Java Benchmarks

This document covers the Java implementation of the serialization benchmarks.

## Prerequisites

- Java 17+
- Bazel 8.4.2+
- Protocol Buffers compiler (`protoc`) - only needed for regenerating proto files

## Running the Benchmark

```bash
# Run the benchmark (tests 1M records by default)
bazel run //java:benchmark
```

## Project Structure

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

## Sample Output

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

## Regenerating Protobuf Code

If you modify the protobuf schema:

```bash
protoc --java_out=java/src/main/java --proto_path=proto proto/feature.proto
```

The generated Java protobuf code will be placed in `java/src/main/java/com/benchmark/proto/Feature.java`.

## Next Steps

- See [Benchmark Scenarios](benchmark-scenarios.md) for methodology details
- See [Benchmark Summary](benchmark-summary.md)
