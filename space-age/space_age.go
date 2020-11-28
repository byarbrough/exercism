// Package space does math with the space-time continuum
package space

// Planet is the name of a planet in our solar system
type Planet string

// earthYearSecs is how many seconds are in one Earth year
const earthYearSecs = 31557600.0

// PlanetYearSecs maps number seconds in a planet's year
var PlanetYearSecs = map[Planet]float64{
	"Mercury": earthYearSecs * 0.2408467,
	"Venus":   earthYearSecs * 0.61519726,
	"Earth":   earthYearSecs * 1.0,
	"Mars":    earthYearSecs * 1.8808158,
	"Jupiter": earthYearSecs * 11.862615,
	"Saturn":  earthYearSecs * 29.447498,
	"Uranus":  earthYearSecs * 84.016846,
	"Neptune": earthYearSecs * 164.79132,
}

// Age returns the age of someone in Earth Years, given their time on another planet
func Age(seconds float64, planet Planet) (age float64) {
	return seconds / PlanetYearSecs[planet]
}
