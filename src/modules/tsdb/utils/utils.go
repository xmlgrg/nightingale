package utils

import (
	"io"
	"os"
	"strconv"
)

// RRDTOOL UTILS
// 监控数据对应的rrd文件名称

const RRDDIRS uint64 = 1000

func QueryRrdFile(seriesID interface{}, dsType string, step int) string {
	switch seriesID.(type) {
	case uint64:
		return strconv.FormatUint(seriesID.(uint64)%RRDDIRS, 10) + "/" +
			strconv.FormatUint(seriesID.(uint64), 10) + "_" + dsType + "_" + strconv.Itoa(step) + ".rrd"
	case string:
		return seriesID.(string)[0:2] + "/" + seriesID.(string) + "_" + dsType + "_" + strconv.Itoa(step) + ".rrd"
	}
	return ""
}

func RrdFileName(baseDir string, seriesID interface{}, dsType string, step int) string {
	switch seriesID.(type) {
	case uint64:
		return baseDir + "/" + strconv.FormatUint(seriesID.(uint64)%RRDDIRS, 10) + "/" +
			strconv.FormatUint(seriesID.(uint64), 10) + "_" + dsType + "_" + strconv.Itoa(step) + ".rrd"
	case string:
		return baseDir + "/" + seriesID.(string)[0:2] + "/" + seriesID.(string) + "_" + dsType + "_" + strconv.Itoa(step) + ".rrd"
	}
	return ""
}

// WriteFile writes data to a file named by filename.
// file must not exist
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

func HashKey(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
