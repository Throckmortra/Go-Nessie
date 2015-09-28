package gomodels

type Atm struct {
	id            string
	name          string
	languageList  []string
	geocode       Geocode
	hours         []string
	accessibility bool
	amountLeft    int
}

type Geocode struct {
	lat float64
	lng float64
}
