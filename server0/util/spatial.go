package util

import "math"

/*

Haversine
formula:	a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
c = 2 ⋅ atan2( √a, √(1−a) )
d = R ⋅ c
where	φ is latitude, λ is longitude, R is earth’s radius (mean radius = 6,371km);
note that angles need to be in radians to pass to trig functions!
JavaScript:
var R = 6371e3; // metres
var φ1 = lat1.toRadians();
var φ2 = lat2.toRadians();
var Δφ = (lat2-lat1).toRadians();
var Δλ = (lon2-lon1).toRadians();

var a = Math.sin(Δφ/2) * Math.sin(Δφ/2) +
        Math.cos(φ1) * Math.cos(φ2) *
        Math.sin(Δλ/2) * Math.sin(Δλ/2);
var c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1-a));

var d = R * c;

 */

// earth radius in meters
const eR = 6371000.
const meterToMiles = 0.000621371
const coeff = eR * meterToMiles

// calculates distance between 2 coordinates in miles using Haversine formula
func Dist(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	var φ1 = ToRadians(lat1)
	var φ2 = ToRadians(lat2)
	var Δφ = ToRadians(lat2 - lat1)
	var Δλ = ToRadians(lon2 - lon1)
	var a = math.Sin(Δφ/2) * math.Sin(Δφ/2) +
		math.Cos(φ1) * math.Cos(φ2) *
			math.Sin(Δλ/2) * math.Sin(Δλ/2)
	c := 2. * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return MeterToMile(c * eR)
}

func ToRadians(deg float64) float64 {
	return deg * math.Pi / 180
}

func MeterToMile(m float64) float64 {
	return m * meterToMiles
}
