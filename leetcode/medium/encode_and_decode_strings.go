package medium

import (
	"fmt"
	"math"
)

// https://neetcode.io/problems/string-encode-and-decode?list=neetcode150

type EncodeAndDecodeStringsSolution struct{}

func (s *EncodeAndDecodeStringsSolution) Encode(strs []string) string {
	result := ""
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		result += fmt.Sprintf("%d#%s", len(str), str)
	}
	return result
}

func (s *EncodeAndDecodeStringsSolution) Decode(str string) []string {
	result := make([]string, 0)

	for i := 0; i < len(str); i++ {
		mult := 1
		length := 0

		for j := i; j < len(str); j++ {
			if str[j] >= '0' && str[j] <= '9' {
				length += mult * int(str[j]-'0')
				mult *= 10
			} else {
				break
			}
		}

		cut := int(math.Log10(float64(mult)))
		result = append(result, str[i+cut+1:i+cut+1+length])
		i += length + 1
	}

	return result
}
