package com.benchmark.config;

import com.google.gson.Gson;
import com.google.gson.annotations.SerializedName;
import java.io.FileReader;
import java.io.IOException;

public class BenchmarkConfig {
    @SerializedName("num_records")
    private int numRecords;

    @SerializedName("num_iterations")
    private int numIterations;

    @SerializedName("data")
    private DataConfig data;

    public static class DataConfig {
        @SerializedName("string_length")
        private int stringLength;

        @SerializedName("binary_size")
        private int binarySize;

        @SerializedName("array_min_size")
        private int arrayMinSize;

        @SerializedName("array_max_size")
        private int arrayMaxSize;

        @SerializedName("map_min_size")
        private int mapMinSize;

        @SerializedName("map_max_size")
        private int mapMaxSize;

        @SerializedName("struct_min_size")
        private int structMinSize;

        @SerializedName("struct_max_size")
        private int structMaxSize;

        public int getStringLength() { return stringLength; }
        public int getBinarySize() { return binarySize; }
        public int getArrayMinSize() { return arrayMinSize; }
        public int getArrayMaxSize() { return arrayMaxSize; }
        public int getMapMinSize() { return mapMinSize; }
        public int getMapMaxSize() { return mapMaxSize; }
        public int getStructMinSize() { return structMinSize; }
        public int getStructMaxSize() { return structMaxSize; }
    }

    public int getNumRecords() { return numRecords; }
    public int getNumIterations() { return numIterations; }
    public DataConfig getData() { return data; }

    public static BenchmarkConfig loadConfig(String path) throws IOException {
        Gson gson = new Gson();
        try (FileReader reader = new FileReader(path)) {
            return gson.fromJson(reader, BenchmarkConfig.class);
        }
    }

    public static BenchmarkConfig getDefault() {
        BenchmarkConfig config = new BenchmarkConfig();
        config.numRecords = 1_000_000;
        config.numIterations = 10;
        config.data = new DataConfig();
        config.data.stringLength = 20;
        config.data.binarySize = 32;
        config.data.arrayMinSize = 1;
        config.data.arrayMaxSize = 5;
        config.data.mapMinSize = 1;
        config.data.mapMaxSize = 3;
        config.data.structMinSize = 1;
        config.data.structMaxSize = 3;
        return config;
    }
}
