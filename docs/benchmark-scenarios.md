# Benchmark Scenarios

This document explains how the serialization benchmarks work and what data they test.

## Data Model

Each `FeatureRecord` represents a **table row** with all 16 data types populated in every record:

### Primitive Numeric Types
- Boolean
- Byte
- Short
- Integer
- Long

### Floating Point Types
- Float
- Double
- Decimal

### String and Binary Types
- String (20 characters)
- Binary (32 bytes)

### Date and Time Types
- Date
- Timestamp
- Timestamp_NTZ (no timezone)

### Complex Types
- Array (1-5 string elements)
- Map (1-3 key-value entries)
- Struct (1-3 fields)

Each record contains **~306-512 bytes** of data depending on the serialization format, making it representative of real-world database rows with mixed data types.

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

## What Gets Measured

Each benchmark reports:

- **p50 (median)**: The typical performance you can expect
- **p99**: The 99th percentile, showing worst-case scenarios
- **Total Size**: Size of all serialized records in MB
- **Avg Size per Record**: Average bytes per record
- **Throughput**: Records processed per second
- **Comparative Metrics**: Speed and size ratios between algorithms

## Example Record

Here's what a single generated record looks like:

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
