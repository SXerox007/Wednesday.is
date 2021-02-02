package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// parse body
func ParseBody(r *http.Request, body interface{}) {
	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		json.NewDecoder(r.Body).Decode(&body)
	}
}

// I, V, X, L, C, D, and M stand respectively for 1, 5, 10, 50, 100, 500, and 1,000
func RomanToIntergerConverter(value string) int {
	// traverse the string
	// match the symbols
	var decimal int
	var lNum int
	//for _, item := range value {
	for i := len(value) - 1; i >= 0; i-- {
		log.Println("Item:", value[i])

		switch value[i] {
		case 'I':
			decimal = processDecimal(1, lNum, decimal)
			lNum = 1
			break
		case 'V':
			decimal = processDecimal(5, lNum, decimal)
			lNum = 5
			break
		case 'X':
			decimal = processDecimal(10, lNum, decimal)
			lNum = 10
			break
		case 'L':
			decimal = processDecimal(50, lNum, decimal)
			lNum = 50
			break
		case 'C':
			decimal = processDecimal(100, lNum, decimal)
			lNum = 100
			break
		case 'D':
			decimal = processDecimal(500, lNum, decimal)
			lNum = 500
			break
		case 'M':
			decimal = processDecimal(1000, lNum, decimal)
			lNum = 1000
			break
		default:
			return -1
		}
	}
	return decimal
}

// process decimal
func processDecimal(decimal, lNum, lDec int) int {
	if lNum > decimal {
		return lNum - decimal
	} else {
		return lNum + decimal
	}
}
