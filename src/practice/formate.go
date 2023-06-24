package main

import (
	"encoding/hex"
	"fmt"
)

// 保留两位小数
func setTwoPrecision(num float64) string {
	return fmt.Sprintf("%.2f\n", num)
}

// 字节切片转16进制
func byte2HEX(bytes []byte) string {
	return hex.EncodeToString(bytes)
}
