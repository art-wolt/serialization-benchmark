package com.benchmark;

import com.benchmark.config.BenchmarkConfig;
import com.benchmark.model.FeatureRecord;
import com.benchmark.model.ProtoConverter;
import com.benchmark.model.RecordGenerator;
import com.benchmark.proto.Feature;
import org.msgpack.core.MessageBufferPacker;
import org.msgpack.core.MessagePack;
import org.msgpack.core.MessageUnpacker;
import org.msgpack.jackson.dataformat.MessagePackFactory;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Benchmark {
    private static final String CONFIG_PATH = "benchmark-config.json";

    static class BenchmarkResult {
        String algorithm;
        int numRecords;
        int numRuns;
        List<Long> serializationTimes = new ArrayList<>();
        List<Long> deserializationTimes = new ArrayList<>();
        long serializationP50;
        long serializationP99;
        long deserializationP50;
        long deserializationP99;
        long totalSize;
        long avgSizePerRecord;
    }

    public static void main(String[] args) {
        System.out.println("Starting serialization benchmark...");

        // Load configuration
        BenchmarkConfig config;
        try {
            config = BenchmarkConfig.loadConfig(CONFIG_PATH);
            System.out.println("Loaded configuration from " + CONFIG_PATH);
        } catch (IOException e) {
            System.out.println("Warning: Could not load config from " + CONFIG_PATH + ": " + e.getMessage());
            System.out.println("Using default configuration...");
            config = BenchmarkConfig.getDefault();
        }

        System.out.println("\nConfiguration:");
        System.out.println("  Records: " + config.getNumRecords());
        System.out.println("  Iterations: " + config.getNumIterations());
        System.out.println("\nGenerating " + config.getNumRecords() + " records...");

        // Generate test data
        RecordGenerator generator = new RecordGenerator(config.getData());
        List<FeatureRecord> records = generator.generateRecords(config.getNumRecords());
        System.out.println("Generated " + records.size() + " records");
        System.out.println("Running " + config.getNumIterations() + " iterations for each algorithm...\n");

        // Run protobuf benchmark
        System.out.println("Running Protobuf benchmark...");
        BenchmarkResult protobufResult = benchmarkProtobuf(records, config.getNumIterations());
        printResult(protobufResult);

        // Run msgpack benchmark
        System.out.println("\nRunning MessagePack benchmark...");
        BenchmarkResult msgpackResult = benchmarkMsgpack(records, config.getNumIterations());
        printResult(msgpackResult);

        // Print comparison
        System.out.println("\n=== Comparison (p50) ===");
        System.out.printf("Serialization: Protobuf %.2fx %s than MessagePack%n",
            (double) msgpackResult.serializationP50 / protobufResult.serializationP50,
            comparison(protobufResult.serializationP50, msgpackResult.serializationP50));

        System.out.printf("Deserialization: Protobuf %.2fx %s than MessagePack%n",
            (double) msgpackResult.deserializationP50 / protobufResult.deserializationP50,
            comparison(protobufResult.deserializationP50, msgpackResult.deserializationP50));

        System.out.printf("Size: Protobuf %.2fx %s than MessagePack%n",
            (double) msgpackResult.totalSize / protobufResult.totalSize,
            comparison(protobufResult.totalSize, msgpackResult.totalSize));
    }

    private static BenchmarkResult benchmarkProtobuf(List<FeatureRecord> records, int runs) {
        BenchmarkResult result = new BenchmarkResult();
        result.algorithm = "Protobuf";
        result.numRecords = records.size();
        result.numRuns = runs;

        // Convert to protobuf format once
        List<Feature.FeatureRecord> pbRecords = new ArrayList<>(records.size());
        for (FeatureRecord record : records) {
            pbRecords.add(ProtoConverter.toProto(record));
        }

        // Run multiple iterations
        for (int run = 0; run < runs; run++) {
            // Benchmark serialization
            byte[][] serializedData = new byte[pbRecords.size()][];
            long totalSize = 0;
            long startTime = System.nanoTime();

            for (int i = 0; i < pbRecords.size(); i++) {
                serializedData[i] = pbRecords.get(i).toByteArray();
                totalSize += serializedData[i].length;
            }

            result.serializationTimes.add(System.nanoTime() - startTime);

            // Store size from first run
            if (run == 0) {
                result.totalSize = totalSize;
            }

            // Benchmark deserialization
            startTime = System.nanoTime();

            for (byte[] data : serializedData) {
                try {
                    Feature.FeatureRecord.parseFrom(data);
                } catch (Exception e) {
                    throw new RuntimeException("Failed to deserialize protobuf record", e);
                }
            }

            result.deserializationTimes.add(System.nanoTime() - startTime);
            System.out.printf("  Run %d/%d complete%n", run + 1, runs);
        }

        // Calculate percentiles
        result.serializationP50 = calculatePercentile(result.serializationTimes, 50);
        result.serializationP99 = calculatePercentile(result.serializationTimes, 99);
        result.deserializationP50 = calculatePercentile(result.deserializationTimes, 50);
        result.deserializationP99 = calculatePercentile(result.deserializationTimes, 99);
        result.avgSizePerRecord = result.totalSize / records.size();

        return result;
    }

    private static BenchmarkResult benchmarkMsgpack(List<FeatureRecord> records, int runs) {
        BenchmarkResult result = new BenchmarkResult();
        result.algorithm = "MessagePack";
        result.numRecords = records.size();
        result.numRuns = runs;

        ObjectMapper objectMapper = new ObjectMapper(new MessagePackFactory());

        // Run multiple iterations
        for (int run = 0; run < runs; run++) {
            // Benchmark serialization
            byte[][] serializedData = new byte[records.size()][];
            long totalSize = 0;
            long startTime = System.nanoTime();

            for (int i = 0; i < records.size(); i++) {
                try {
                    serializedData[i] = objectMapper.writeValueAsBytes(records.get(i));
                    totalSize += serializedData[i].length;
                } catch (Exception e) {
                    throw new RuntimeException("Failed to serialize msgpack record " + i, e);
                }
            }

            result.serializationTimes.add(System.nanoTime() - startTime);

            // Store size from first run
            if (run == 0) {
                result.totalSize = totalSize;
            }

            // Benchmark deserialization
            startTime = System.nanoTime();

            for (byte[] data : serializedData) {
                try {
                    objectMapper.readValue(data, FeatureRecord.class);
                } catch (Exception e) {
                    throw new RuntimeException("Failed to deserialize msgpack record", e);
                }
            }

            result.deserializationTimes.add(System.nanoTime() - startTime);
            System.out.printf("  Run %d/%d complete%n", run + 1, runs);
        }

        // Calculate percentiles
        result.serializationP50 = calculatePercentile(result.serializationTimes, 50);
        result.serializationP99 = calculatePercentile(result.serializationTimes, 99);
        result.deserializationP50 = calculatePercentile(result.deserializationTimes, 50);
        result.deserializationP99 = calculatePercentile(result.deserializationTimes, 99);
        result.avgSizePerRecord = result.totalSize / records.size();

        return result;
    }

    private static long calculatePercentile(List<Long> values, int percentile) {
        if (values.isEmpty()) return 0;

        List<Long> sorted = new ArrayList<>(values);
        sorted.sort(Long::compareTo);

        int index = (int) Math.round((sorted.size() - 1) * percentile / 100.0);
        index = Math.max(0, Math.min(index, sorted.size() - 1));

        return sorted.get(index);
    }

    private static void printResult(BenchmarkResult result) {
        System.out.println("Algorithm: " + result.algorithm);
        System.out.println("  Records: " + result.numRecords);
        System.out.println("  Runs: " + result.numRuns);
        System.out.println("  Serialization:");
        System.out.printf("    p50: %.2fs (%.2f ns/op)%n",
            result.serializationP50 / 1e9,
            (double) result.serializationP50 / result.numRecords);
        System.out.printf("    p99: %.2fs (%.2f ns/op)%n",
            result.serializationP99 / 1e9,
            (double) result.serializationP99 / result.numRecords);
        System.out.println("  Deserialization:");
        System.out.printf("    p50: %.2fs (%.2f ns/op)%n",
            result.deserializationP50 / 1e9,
            (double) result.deserializationP50 / result.numRecords);
        System.out.printf("    p99: %.2fs (%.2f ns/op)%n",
            result.deserializationP99 / 1e9,
            (double) result.deserializationP99 / result.numRecords);
        System.out.printf("  Total Size: %.2f MB%n", result.totalSize / (1024.0 * 1024.0));
        System.out.printf("  Avg Size per Record: %d bytes%n", result.avgSizePerRecord);
    }

    private static String comparison(long a, long b) {
        return a < b ? "faster" : "slower";
    }
}
