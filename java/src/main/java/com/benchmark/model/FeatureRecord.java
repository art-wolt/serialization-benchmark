package com.benchmark.model;

import java.util.List;
import java.util.Map;

public class FeatureRecord {
    private String id;
    private long timestamp;

    // Primitive numeric types
    private boolean boolValue;
    private byte byteValue;
    private short shortValue;
    private int intValue;
    private long longValue;

    // Floating point types
    private float floatValue;
    private double doubleValue;
    private String decimalValue;  // Using String for arbitrary precision

    // String and binary types
    private String stringValue;
    private byte[] binaryValue;

    // Date and time types
    private int dateValue;          // Days since epoch
    private long timestampValue;    // Microseconds since epoch
    private long timestampNtzValue; // Microseconds since epoch (no timezone)

    // Complex types
    private ArrayValue arrayValue;
    private MapValue mapValue;
    private StructValue structValue;

    public static class ArrayValue {
        private List<String> elements;

        public ArrayValue() {}
        public ArrayValue(List<String> elements) {
            this.elements = elements;
        }

        public List<String> getElements() { return elements; }
        public void setElements(List<String> elements) { this.elements = elements; }
    }

    public static class MapValue {
        private Map<String, String> entries;

        public MapValue() {}
        public MapValue(Map<String, String> entries) {
            this.entries = entries;
        }

        public Map<String, String> getEntries() { return entries; }
        public void setEntries(Map<String, String> entries) { this.entries = entries; }
    }

    public static class StructValue {
        private Map<String, String> fields;

        public StructValue() {}
        public StructValue(Map<String, String> fields) {
            this.fields = fields;
        }

        public Map<String, String> getFields() { return fields; }
        public void setFields(Map<String, String> fields) { this.fields = fields; }
    }

    // Getters and setters
    public String getId() { return id; }
    public void setId(String id) { this.id = id; }

    public long getTimestamp() { return timestamp; }
    public void setTimestamp(long timestamp) { this.timestamp = timestamp; }

    public boolean isBoolValue() { return boolValue; }
    public void setBoolValue(boolean boolValue) { this.boolValue = boolValue; }

    public byte getByteValue() { return byteValue; }
    public void setByteValue(byte byteValue) { this.byteValue = byteValue; }

    public short getShortValue() { return shortValue; }
    public void setShortValue(short shortValue) { this.shortValue = shortValue; }

    public int getIntValue() { return intValue; }
    public void setIntValue(int intValue) { this.intValue = intValue; }

    public long getLongValue() { return longValue; }
    public void setLongValue(long longValue) { this.longValue = longValue; }

    public float getFloatValue() { return floatValue; }
    public void setFloatValue(float floatValue) { this.floatValue = floatValue; }

    public double getDoubleValue() { return doubleValue; }
    public void setDoubleValue(double doubleValue) { this.doubleValue = doubleValue; }

    public String getDecimalValue() { return decimalValue; }
    public void setDecimalValue(String decimalValue) { this.decimalValue = decimalValue; }

    public String getStringValue() { return stringValue; }
    public void setStringValue(String stringValue) { this.stringValue = stringValue; }

    public byte[] getBinaryValue() { return binaryValue; }
    public void setBinaryValue(byte[] binaryValue) { this.binaryValue = binaryValue; }

    public int getDateValue() { return dateValue; }
    public void setDateValue(int dateValue) { this.dateValue = dateValue; }

    public long getTimestampValue() { return timestampValue; }
    public void setTimestampValue(long timestampValue) { this.timestampValue = timestampValue; }

    public long getTimestampNtzValue() { return timestampNtzValue; }
    public void setTimestampNtzValue(long timestampNtzValue) { this.timestampNtzValue = timestampNtzValue; }

    public ArrayValue getArrayValue() { return arrayValue; }
    public void setArrayValue(ArrayValue arrayValue) { this.arrayValue = arrayValue; }

    public MapValue getMapValue() { return mapValue; }
    public void setMapValue(MapValue mapValue) { this.mapValue = mapValue; }

    public StructValue getStructValue() { return structValue; }
    public void setStructValue(StructValue structValue) { this.structValue = structValue; }
}
