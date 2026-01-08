# Benchmark Summary

Performance comparison of Protobuf vs MessagePack across Golang and Java.

## Test Configuration

- **Records**: 1,000,000 (1M)
- **Iterations**: 10 per algorithm
- **Data Types**: 16 types per record (primitives, strings, dates, complex types)
- **Record Size**: ~306-512 bytes depending on format

## Performance Overview

### Golang Results (p50)

| Metric | Protobuf | MessagePack | Winner |
|--------|----------|-------------|--------|
| **Serialization** | 1627ms | 1571ms | MessagePack (3.5% faster) |
| **Deserialization** | 1765ms | 2047ms | Protobuf (16% faster) |
| **Size** | 292 MB | 489 MB | Protobuf (40% smaller) |
| **Size per Record** | 306 bytes | 512 bytes | Protobuf (1.68x smaller) |

**Throughput (Golang):**
- Protobuf: 614K rec/sec (ser), 567K rec/sec (deser)
- MessagePack: 636K rec/sec (ser), 489K rec/sec (deser)

### Java Results (p50)

| Metric | Protobuf | MessagePack | Winner |
|--------|----------|-------------|--------|
| **Serialization** | 505ms | 1480ms | Protobuf (2.93x faster) |
| **Deserialization** | 1017ms | 2527ms | Protobuf (2.49x faster) |
| **Size** | 292 MB | 469 MB | Protobuf (38% smaller) |
| **Size per Record** | 306 bytes | 492 bytes | Protobuf (1.60x smaller) |

**Throughput (Java):**
- Protobuf: 1.98M rec/sec (ser), 983K rec/sec (deser)
- MessagePack: 676K rec/sec (ser), 396K rec/sec (deser)

## Cross-Language Comparison

### Protobuf Performance

| Operation | Golang | Java | Faster |
|-----------|--------|------|--------|
| **Serialization** | 1627ms | 505ms | Java (3.22x faster) |
| **Deserialization** | 1765ms | 1017ms | Java (1.74x faster) |
| **Size** | 292 MB | 292 MB | Equal |

### MessagePack Performance

| Operation | Golang | Java | Faster |
|-----------|--------|------|--------|
| **Serialization** | 1571ms | 1480ms | Java (6% faster) |
| **Deserialization** | 2047ms | 2527ms | Golang (23% faster) |
| **Size** | 489 MB | 469 MB | Java (4% smaller) |

## Key Insights

### Algorithm Comparison

**Protobuf wins on:**
- Deserialization speed (both languages)
- Size efficiency (both languages, 40% smaller)
- Java performance (significantly faster than MessagePack)

**MessagePack wins on:**
- Golang serialization (slightly faster, 3.5%)
- Simplicity (schema-less format)
