package com.benchmark.model;

import com.benchmark.proto.Feature;
import com.google.protobuf.ByteString;

public class ProtoConverter {
    public static Feature.FeatureRecord toProto(FeatureRecord record) {
        Feature.FeatureRecord.Builder builder = Feature.FeatureRecord.newBuilder()
            .setId(record.getId())
            .setTimestamp(record.getTimestamp())
            .setBoolValue(record.isBoolValue())
            .setByteValue(record.getByteValue())
            .setShortValue(record.getShortValue())
            .setIntValue(record.getIntValue())
            .setLongValue(record.getLongValue())
            .setFloatValue(record.getFloatValue())
            .setDoubleValue(record.getDoubleValue())
            .setDecimalValue(record.getDecimalValue())
            .setStringValue(record.getStringValue())
            .setBinaryValue(ByteString.copyFrom(record.getBinaryValue()))
            .setDateValue(record.getDateValue())
            .setTimestampValue(record.getTimestampValue())
            .setTimestampNtzValue(record.getTimestampNtzValue());

        // Array value
        if (record.getArrayValue() != null) {
            Feature.ArrayValue.Builder arrayBuilder = Feature.ArrayValue.newBuilder();
            arrayBuilder.addAllElements(record.getArrayValue().getElements());
            builder.setArrayValue(arrayBuilder.build());
        }

        // Map value
        if (record.getMapValue() != null) {
            Feature.MapValue.Builder mapBuilder = Feature.MapValue.newBuilder();
            mapBuilder.putAllEntries(record.getMapValue().getEntries());
            builder.setMapValue(mapBuilder.build());
        }

        // Struct value
        if (record.getStructValue() != null) {
            Feature.StructValue.Builder structBuilder = Feature.StructValue.newBuilder();
            structBuilder.putAllFields(record.getStructValue().getFields());
            builder.setStructValue(structBuilder.build());
        }

        return builder.build();
    }

    public static FeatureRecord fromProto(Feature.FeatureRecord proto) {
        FeatureRecord record = new FeatureRecord();
        record.setId(proto.getId());
        record.setTimestamp(proto.getTimestamp());
        record.setBoolValue(proto.getBoolValue());
        record.setByteValue((byte) proto.getByteValue());
        record.setShortValue((short) proto.getShortValue());
        record.setIntValue(proto.getIntValue());
        record.setLongValue(proto.getLongValue());
        record.setFloatValue(proto.getFloatValue());
        record.setDoubleValue(proto.getDoubleValue());
        record.setDecimalValue(proto.getDecimalValue());
        record.setStringValue(proto.getStringValue());
        record.setBinaryValue(proto.getBinaryValue().toByteArray());
        record.setDateValue(proto.getDateValue());
        record.setTimestampValue(proto.getTimestampValue());
        record.setTimestampNtzValue(proto.getTimestampNtzValue());

        // Array value
        if (proto.hasArrayValue()) {
            record.setArrayValue(new FeatureRecord.ArrayValue(
                proto.getArrayValue().getElementsList()
            ));
        }

        // Map value
        if (proto.hasMapValue()) {
            record.setMapValue(new FeatureRecord.MapValue(
                proto.getMapValue().getEntriesMap()
            ));
        }

        // Struct value
        if (proto.hasStructValue()) {
            record.setStructValue(new FeatureRecord.StructValue(
                proto.getStructValue().getFieldsMap()
            ));
        }

        return record;
    }
}
