package geocoder

import (
	"container/heap"
	"errors"
	"math"
)

type cityHeapNode struct {
	distance float64
	city     loc
}
type cityHeap []cityHeapNode

func (h cityHeap) Len() int           { return len(h) }
func (h cityHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h cityHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *cityHeap) Push(x interface{}) {
	*h = append(*h, x.(cityHeapNode))
}
func (h *cityHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findNearestZOrder(k int, lat float64, lon float64) []cityHeapNode {
	var latRad = (lat * 0.017453292519943295)
	var lonRad = (lon * 0.017453292519943295)
	var coslat = math.Cos(latRad)
	var sinlat = math.Sin(latRad)
	var coslon = math.Cos(lonRad)
	var sinlon = math.Sin(lonRad)

	heapq := &cityHeap{}
	heap.Init(heapq)

	var pos = int(math.Abs(float64(bsearch(objs, encode(lat, lon)))))
	var count = len(objs[pos].l)
	var maxOffset = int(math.Min(float64(pos), float64(len(objs)-pos)))
	var offset = 4
	for ; offset < maxOffset; offset++ {
		count += (len(objs[pos-offset].l) + len(objs[pos+offset].l))
		if count >= k {
			break
		}
	}
	offset = int(math.Ceil(float64(offset) * 1.25)) // fudge
	var elements = objs[int(math.Max(0, float64(pos-offset))):int(math.Min(float64(len(objs)), float64((pos+offset+1))))]
	for ii, _ := range elements {
		for jj, _ := range elements[ii].l {
			var city = elements[ii].l[jj]
			var val = (coslat*city.coslat*(city.coslon*coslon+city.sinlon*sinlon) + sinlat*city.sinlat)
			if heapq.Len() == k {
				if (*heapq)[0].distance < val {
					heap.Pop(heapq)
					heap.Push(heapq, cityHeapNode{
						distance: val,
						city:     city,
					})
				}
			} else {
				heap.Push(heapq, cityHeapNode{
					distance: val,
					city:     city,
				})
			}
		}
	}

	var resultArray = []cityHeapNode{}
	var entry interface{}
	for heapq.Len() > 0 {
		entry = heap.Pop(heapq)
		resultArray = append([]cityHeapNode{entry.(cityHeapNode)}, resultArray...)
	}

	return resultArray
}

type Result struct {
	country      string
	country_code string
	region       string
	region_code  string
	city         string
	latitude     float64
	longitude    float64
	distance     float64
}

func GetNearestCities(lat float64, lon float64, k int, units string) ([]Result, error) {
	if k < 1 {
		k = 1
	}
	if units == "" {
		units = "km"
	}
	if !("km" == units || "mi" == units) {
		return nil, errors.New("Distance units must be defined in kilometers (km) or miles (mi)")
	}
	var locations = findNearestZOrder(k, lat, lon)
	var results = []Result{}
	for _, location := range locations {
		var distance float64
		if "km" == units {
			distance = math.Acos(location.distance) * 6371.3
		} else {
			distance = math.Acos(location.distance) * 3959.0
		}
		results = append(results, Result{
			country:      location.city.src.country,
			country_code: location.city.src.country_code,
			region:       location.city.src.region,
			region_code:  location.city.src.region_code,
			city:         location.city.src.name,
			latitude:     location.city.src.lat,
			longitude:    location.city.src.lon,
			distance:     distance,
		})
	}
	return results, nil
}
