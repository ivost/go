package model

// Caribou Coffee	100 5th St	 	55402			44.978348	-93.268623	5

type POI struct {
	Name      string
	Address1  string
	Address2  string
	Zip       string
	ZipSuffix string
	Phone     string
	Lat       float64
	Lng       float64
	Radius    float64
}

