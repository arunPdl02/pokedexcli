package pokeapi

type RespShallowLocations struct {
	Count     int     `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
