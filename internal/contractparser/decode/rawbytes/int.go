package rawbytes

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

func decodeInt(hex string) (string, int, error) {
	var ret, buffer string
	var i int
	var err error

	for 2*i < len(hex) {
		start := 2 * i
		end := start + 2
		part := hex[start:end]
		buffer += part
		i++

		val, err := strconv.ParseInt(part, 16, 32)
		if err != nil {
			return ret, 0, err
		}
		if val < 127 {
			break
		}
	}

	ret, err = decodeSignedHex(buffer)
	if err != nil {
		return ret, 0, err
	}

	return ret, i * 2, err
}

func decodeSigned(source []byte) (string, error) {
	if len(source) == 0 {
		return "", fmt.Errorf("expected non-empty byte array")
	}

	// Split input into 8-bit bitstrings
	segments := make([]string, len(source))
	for i, curByte := range source {
		segments[i] = fmt.Sprintf("%08b", curByte)
	}

	// Trim off leading continuation bit from each segment
	for i, segment := range segments {
		segments[i] = segment[1:]
	}

	// Trim off the sign flag from the first segment
	firstSegment := []rune(segments[0])
	isNegative := firstSegment[0] == '1'
	segments[0] = string(firstSegment[1:])

	// Reverse the order of the segments.
	// Source: https://github.com/golang/go/wiki/SliceTricks#reversing
	for i := len(segments)/2 - 1; i >= 0; i-- {
		opp := len(segments) - 1 - i
		segments[i], segments[opp] = segments[opp], segments[i]
	}

	// Concat all the bits
	bitStringBuf := bytes.Buffer{}
	for _, segment := range segments {
		bitStringBuf.WriteString(segment)
	}
	bitString := bitStringBuf.String()

	// Add sign flag
	if isNegative {
		bitString = "-" + bitString
	}

	// Convert from base 2 to base 10
	ret := new(big.Int)
	_, success := ret.SetString(bitString, 2)
	if !success {
		return "", fmt.Errorf("failed to parse bit string %s to big.Int", bitString)
	}

	return fmt.Sprintf("%v", ret), nil
}

// decodeSignedHex decodes a zarith encoded signed integer from the entire input hex string.
// Assumes the input contains no extra trailing bytes.
func decodeSignedHex(source string) (string, error) {
	bytes, err := hex.DecodeString(source)
	if err != nil {
		return "", err
	}

	return decodeSigned(bytes)
}