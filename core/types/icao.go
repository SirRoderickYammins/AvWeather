package types

type Airport struct {
	ICAO      string // For ident...might make ICAO it's own type
	Latitude  string // For future potential map
	Longitude string // For future potential map
	IAP       bool   // Airport has an instrument approach?
}

// Create airport gen function
/* func genAirport(airport *Airport) err error {
} */
