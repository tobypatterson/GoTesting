package location

type Country struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	TimeZone string `json:"time_zone"`
	// Location       string `json:"geo_information"`
	GeoInformation GeoInformation `json:"geo_information"`
	State          []State        `json:"states"`
}

type GeoInformation struct {
	Location GeoLocation `json:"location"`
}

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type State struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
