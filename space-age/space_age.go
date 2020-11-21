// Package space does math with the space-time continuum
package space

// Planet is the name of a planet in our solar system
type Planet string

const secondsInYear = 31557600

var relativePeriod = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the age of someone in Earth Years, given their time on another planet
func Age(seconds float64, planet Planet) (age float64) {
	// Convert seconds to years
	years := seconds / secondsInYear

	return years / relativePeriod[planet]
}
