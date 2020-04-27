package encode

import (
	"strconv"
)

// RunLengthEncode RLE encoding
func RunLengthEncode(input string) string {
	response := ""
	n := len(input)

	for i := 0; i < n; i++ {
		count := 1
		for i < n-1 && input[i] == input[i+1] {
			count++
			i++
		}
		if count > 1 {
			response = response + strconv.Itoa(count) + string(input[i])
		} else {
			response = response + string(input[i])
		}
	}
	return response
}

// RunLengthDecode RLE decoding
func RunLengthDecode(input string) string {

	response := ""
	currentCount := ""

	for _, s := range input {

		if _, err := strconv.ParseInt(string(s), 10, 64); err == nil {
			currentCount += string(s)
		} else {
			convertedNumber, _ := strconv.Atoi(currentCount)

			if convertedNumber == 0 {
				convertedNumber++
			}

			for i := 0; i < convertedNumber; i++ {
				response = response + string(s)
			}
			currentCount = ""
		}
	}
	return response

}
