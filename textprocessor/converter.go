package textprocessor

import (
	"fmt"
	"strconv"
)

// Конвертирует HEX → DEC
func ConvertHexToDec(hexStr string) string {
	num, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return hexStr
	}
	return fmt.Sprintf("%d", num)
}

// Конвертирует BIN → DEC
func ConvertBinToDec(binStr string) string {
	num, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return binStr
	}
	return fmt.Sprintf("%d", num)
}
