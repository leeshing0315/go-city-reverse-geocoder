package geocoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeocoder(t *testing.T) {
	results, _ := GetNearestCities(22.570627, 113.941378, 3, "mi")
	println(results[0].City, results[0].Distance, results[0].Latitude, results[0].Longitude)
	println(results[1].City, results[1].Distance, results[1].Latitude, results[1].Longitude)
	println(results[2].City, results[2].Distance, results[2].Latitude, results[2].Longitude)
	assert.Equal(t, "Shenzhen", results[0].City)
}
