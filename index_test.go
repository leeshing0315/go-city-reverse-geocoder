package geocoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeocoder(t *testing.T) {
	results, _ := GetNearestCities(22.570627, 113.941378, 3, "mi")
	println(results[0].city, results[0].distance, results[0].latitude, results[0].longitude)
	println(results[1].city, results[1].distance, results[1].latitude, results[1].longitude)
	println(results[2].city, results[2].distance, results[2].latitude, results[2].longitude)
	assert.Equal(t, "Shenzhen", results[0].city)
}
