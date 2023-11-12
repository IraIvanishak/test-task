package breeds

// Breed represents the structure of a single cat breed
type Breed struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}

// CountryBreeds groups breeds by country for JSON serialization.
type CountryBreeds struct {
	Country string   `json:"country"`
	Breeds  []string `json:"breeds"`
}

// FormattedBreeds is a slice of CountryBreeds to hold grouped breed data.
type FormattedBreeds []CountryBreeds
