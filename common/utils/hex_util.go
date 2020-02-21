package utils

import (
	"encoding/hex"
	"fmt"
)

func HexEncode(src string) string {
	return hex.EncodeToString([]byte(src))
}

func HexEncodeWithByte(src []byte) string {
	return hex.EncodeToString(src)
}

func HexDecode(encodeSrc string) ([]byte, error) {
	b, err := hex.DecodeString(encodeSrc)
	if err != nil {
		fmt.Printf("hex decode error %v \n", err.Error())
	}

	return b, err
}
