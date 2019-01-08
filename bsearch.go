package geocoder

func bsearch(arr []objMetaData, searchElement string) int {
	var minIndex = 0
	var maxIndex = len(arr) - 1
	var currentIndex = 0
	var currentElement objMetaData

	for minIndex <= maxIndex {
		currentIndex = (minIndex+maxIndex)/2 | 0
		currentElement = arr[currentIndex]

		if currentElement.h < searchElement {
			minIndex = currentIndex + 1
		} else if currentElement.h > searchElement {
			maxIndex = currentIndex - 1
		} else {
			return currentIndex
		}
	}
	return maxIndex
}
