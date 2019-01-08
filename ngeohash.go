package geocoder

import (
	"strings"
)

var BASE32_CODES = "0123456789bcdefghjkmnpqrstuvwxyz"

func encode(latitude float64, longitude float64) string {
	chars := []string{}
	bits := 0
	bitsTotal := 0
	hash_value := 0
	var maxLat float64 = 90
	var minLat float64 = -90
	var maxLon float64 = 180
	var minLon float64 = -180
	var mid float64
	for len(chars) < 6 {
		if bitsTotal%2 == 0 {
			mid = (maxLon + minLon) / 2
			if longitude > mid {
				hash_value = (hash_value << 1) + 1
				minLon = mid
			} else {
				hash_value = (hash_value << 1) + 0
				maxLon = mid
			}
		} else {
			mid = (maxLat + minLat) / 2
			if latitude > mid {
				hash_value = (hash_value << 1) + 1
				minLat = mid
			} else {
				hash_value = (hash_value << 1) + 0
				maxLat = mid
			}
		}

		bits++
		bitsTotal++
		if bits == 5 {
			var code = BASE32_CODES[hash_value]
			chars = append(chars, string(code))
			bits = 0
			hash_value = 0
		}
	}
	return strings.Join(chars, "")
}
