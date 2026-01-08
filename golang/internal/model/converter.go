package model

import (
	pb "serialization-benchmark/proto"
)

// ConvertToProto converts a msgpack FeatureRecord to protobuf FeatureRecord
func ConvertToProto(record *FeatureRecord) *pb.FeatureRecord {
	pbRecord := &pb.FeatureRecord{
		Id:        record.ID,
		Timestamp: record.Timestamp.UnixMicro(),

		// Primitive numeric types
		BoolValue:  record.BoolValue,
		ByteValue:  int32(record.ByteValue),
		ShortValue: int32(record.ShortValue),
		IntValue:   record.IntValue,
		LongValue:  record.LongValue,

		// Floating point types
		FloatValue:   record.FloatValue,
		DoubleValue:  record.DoubleValue,
		DecimalValue: record.DecimalValue,

		// String and binary types
		StringValue: record.StringValue,
		BinaryValue: record.BinaryValue,

		// Date and time types
		DateValue:         record.DateValue,
		TimestampValue:    record.TimestampValue,
		TimestampNtzValue: record.TimestampNtzValue,
	}

	// Convert array value
	if record.ArrayValue != nil {
		pbRecord.ArrayValue = &pb.ArrayValue{
			Elements: record.ArrayValue.Elements,
		}
	}

	// Convert map value
	if record.MapValue != nil {
		pbRecord.MapValue = &pb.MapValue{
			Entries: record.MapValue.Entries,
		}
	}

	// Convert struct value
	if record.StructValue != nil {
		pbRecord.StructValue = &pb.StructValue{
			Fields: record.StructValue.Fields,
		}
	}

	return pbRecord
}
