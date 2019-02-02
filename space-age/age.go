package space

type Planet string
var Earth float64 = 31557600
 

var PlanetSeconds = map[Planet]float64{
	"Earth":   Earth,
	"Mercury": Earth * 0.2408467,
	"Venus":   Earth * 0.61519726,
	"Mars":    Earth * 1.8808158,
	"Jupiter": Earth * 11.862615,
	"Saturn":  Earth * 29.447498,
	"Uranus":  Earth * 84.016846,
	"Neptune": Earth * 164.79132,
}


func Age(ageSeconds float64, planetName Planet) float64 {
	toEarthYears := ageSeconds/PlanetSeconds[planetName]

	return toEarthYears
}