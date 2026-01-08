package model

import "time"

type ArrayValue struct {
	Elements []string `msgpack:"elements"`
}

type MapValue struct {
	Entries map[string]string `msgpack:"entries"`
}

type StructValue struct {
	Fields map[string]string `msgpack:"fields"`
}

// FeatureRecord represents a table row with all data types
type FeatureRecord struct {
	ID        string    `msgpack:"id"`
	Timestamp time.Time `msgpack:"timestamp"`

	// Primitive numeric types
	BoolValue  bool  `msgpack:"bool_value"`
	ByteValue  int8  `msgpack:"byte_value"`
	ShortValue int16 `msgpack:"short_value"`
	IntValue   int32 `msgpack:"int_value"`
	LongValue  int64 `msgpack:"long_value"`

	// Floating point types
	FloatValue   float32 `msgpack:"float_value"`
	DoubleValue  float64 `msgpack:"double_value"`
	DecimalValue string  `msgpack:"decimal_value"`

	// String and binary types
	StringValue string `msgpack:"string_value"`
	BinaryValue []byte `msgpack:"binary_value"`

	// Date and time types
	DateValue         int32 `msgpack:"date_value"`
	TimestampValue    int64 `msgpack:"timestamp_value"`
	TimestampNtzValue int64 `msgpack:"timestamp_ntz_value"`

	// Complex types (containers)
	ArrayValue  *ArrayValue  `msgpack:"array_value"`
	MapValue    *MapValue    `msgpack:"map_value"`
	StructValue *StructValue `msgpack:"struct_value"`
}
