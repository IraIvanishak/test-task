package main

import (
	"github.com/IraIvanishak/test-task-contentquo/breeds"
	"github.com/IraIvanishak/test-task-contentquo/dataTransfer"
)

func main() {
	// Retrieve cat breed data from API.
	breedsData := dataTransfer.GetDataFromAPI().Data

	// Logic
	mapBreeds := breeds.GroupBreedsNamesByCountry(breedsData)
	breeds.OrderByLength(mapBreeds)

	// Convert map[string][]strimg to FormattedBreeds for json keys; order changes
	var formattedBreeds breeds.FormattedBreeds
	for country, breedsList := range mapBreeds {
		formattedBreeds = append(formattedBreeds, breeds.CountryBreeds{Country: country, Breeds: breedsList})
	}

	// Write the formatted data to a JSON file.
	dataTransfer.WriteResultsToFile(formattedBreeds)
}
