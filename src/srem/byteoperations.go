package main

import (
"encoding/binary"
"bytes"
)

func int64ToBytes(value uint64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, value)
	return buf.Bytes()
}

func bytesToInt64(b []byte) uint64 {
    return binary.BigEndian.Uint64(b)
}

func getDataFromInstruction(instruction []byte) []byte {
	data := []byte{instruction[2], instruction[3], instruction[4],  instruction[5],  instruction[6],  instruction[7],  instruction[8],  instruction[9]}
	return data
}