# Serialization Performance Benchmark

A performance testing repository for comparing different serialization algorithms across multiple programming languages.

## What It Does

This benchmark compares serialization performance of:
- **Protobuf** (Protocol Buffers)
- **MessagePack** (msgpack)

Across multiple languages:
- Golang ✅
- Java ✅

Each test serializes and deserializes 1 million records (default setting) containing all common data types (primitives, strings, dates, complex types like arrays/maps/structs), measuring speed and size efficiency.

## Quick Start

**Golang:**
```bash
bazel run //golang/cmd/benchmark:benchmark
```

**Java:**
```bash
bazel run //java:benchmark
```

## Documentation

📚 See the [documentation](docs/) for detailed information:

- [Benchmark Scenarios & Methodology](docs/benchmark-scenarios.md) - How the benchmarks work
- [Golang Implementation](docs/benchmark-go.md) - Running and understanding Go benchmarks
- [Java Implementation](docs/benchmark-java.md) - Running and understanding Java benchmarks
- [Benchmark Results & Summary](docs/benchmark-summary.md) - Performance comparisons

## License

MIT
