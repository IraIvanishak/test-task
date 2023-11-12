package breeds

import "sort"

// GroupBreedsNamesByCountry groups cat breeds by their country.
func GroupBreedsNamesByCountry(breeds []Breed) map[string][]string {
	groupedBreeds := make(map[string][]string)
	for _, breed := range breeds {
		groupedBreeds[breed.Country] = append(groupedBreeds[breed.Country], breed.Breed)
	}
	return groupedBreeds
}

// OrderByLength sorts the breeds within each country by the length of their names.
func OrderByLength(breeds map[string][]string) {
	for _, names := range breeds {
		sort.Slice(names, func(i, j int) bool {
			return len(names[i]) < len(names[j])
		})
	}
}
