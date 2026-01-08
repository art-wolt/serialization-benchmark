package com.benchmark.model;

import com.benchmark.config.BenchmarkConfig.DataConfig;
import java.security.SecureRandom;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class RecordGenerator {
    private static final String CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    private final SecureRandom random;
    private final DataConfig config;

    public RecordGenerator(DataConfig config) {
        this.random = new SecureRandom();
        this.config = config;
    }

    public List<FeatureRecord> generateRecords(int count) {
        List<FeatureRecord> records = new ArrayList<>(count);
        for (int i = 0; i < count; i++) {
            records.add(generateRecord("record_" + i));
        }
        return records;
    }

    public FeatureRecord generateRecord(String id) {
        FeatureRecord record = new FeatureRecord();
        long now = System.currentTimeMillis() * 1000; // Convert to microseconds

        record.setId(id);
        record.setTimestamp(now);

        // Primitive numeric types
        record.setBoolValue(random.nextBoolean());
        record.setByteValue((byte) random.nextInt(256));
        record.setShortValue((short) random.nextInt(65536));
        record.setIntValue(random.nextInt());
        record.setLongValue(random.nextLong());

        // Floating point types
        record.setFloatValue(random.nextFloat() * 10000);
        record.setDoubleValue(random.nextDouble() * 10000);
        record.setDecimalValue(String.format("%.2f", random.nextDouble() * 10000));

        // String and binary types
        record.setStringValue(randomString(config.getStringLength()));
        record.setBinaryValue(randomBytes(config.getBinarySize()));

        // Date and time types
        record.setDateValue(random.nextInt(20000)); // Days since epoch
        record.setTimestampValue(now);
        record.setTimestampNtzValue(now);

        // Complex types
        record.setArrayValue(generateArrayValue());
        record.setMapValue(generateMapValue());
        record.setStructValue(generateStructValue());

        return record;
    }

    private FeatureRecord.ArrayValue generateArrayValue() {
        int size = random.nextInt(config.getArrayMaxSize() - config.getArrayMinSize() + 1)
                   + config.getArrayMinSize();
        List<String> elements = new ArrayList<>(size);
        for (int i = 0; i < size; i++) {
            elements.add(randomString(10));
        }
        return new FeatureRecord.ArrayValue(elements);
    }

    private FeatureRecord.MapValue generateMapValue() {
        int size = random.nextInt(config.getMapMaxSize() - config.getMapMinSize() + 1)
                   + config.getMapMinSize();
        Map<String, String> entries = new HashMap<>(size);
        for (int i = 0; i < size; i++) {
            entries.put(randomString(8), randomString(12));
        }
        return new FeatureRecord.MapValue(entries);
    }

    private FeatureRecord.StructValue generateStructValue() {
        int size = random.nextInt(config.getStructMaxSize() - config.getStructMinSize() + 1)
                   + config.getStructMinSize();
        Map<String, String> fields = new HashMap<>(size);
        for (int i = 0; i < size; i++) {
            fields.put(randomString(8), randomString(12));
        }
        return new FeatureRecord.StructValue(fields);
    }

    private String randomString(int length) {
        StringBuilder sb = new StringBuilder(length);
        for (int i = 0; i < length; i++) {
            sb.append(CHARSET.charAt(random.nextInt(CHARSET.length())));
        }
        return sb.toString();
    }

    private byte[] randomBytes(int length) {
        byte[] bytes = new byte[length];
        random.nextBytes(bytes);
        return bytes;
    }
}
