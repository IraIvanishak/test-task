package breeds

import (
	"reflect"
	"testing"
)

func TestGroupBreedsNamesByCountry(t *testing.T) {
	breeds := []Breed{
		{"Siamese", "Thailand", "", "", ""},
		{"Maine Coon", "USA", "", "", ""},
		{"Siamese", "Thailand", "", "", ""},
	}

	grouped := GroupBreedsNamesByCountry(breeds)

	expected := map[string][]string{
		"Thailand": {"Siamese", "Siamese"},
		"USA":      {"Maine Coon"},
	}

	if !reflect.DeepEqual(grouped, expected) {
		t.Errorf("Expected %v, got %v", expected, grouped)
	}
}

func TestOrderByLength(t *testing.T) {
	grouped := map[string][]string{
		"Thailand": {"fjsdkfdksfsdkgk", "dkfkd"},
		"USA":      {"fddf", "llft"},
	}

	OrderByLength(grouped)

	expected := map[string][]string{
		"Thailand": {"dkfkd", "fjsdkfdksfsdkgk"},
		"USA":      {"fddf", "llft"},
	}

	if !reflect.DeepEqual(grouped, expected) {
		t.Errorf("Expected %v, got %v", expected, grouped)
	}
}
