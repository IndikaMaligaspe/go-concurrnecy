package movies

import (
	"fmt"
)

type Movie struct {
	ID           int
	Name         string
	Director     string
	YearReleased int
}

func (m Movie) String() string {
	return fmt.Sprintf(
		"Name: \t\t%q\n"+
			"Director: \t%q\n"+
			"Released:\t%d\n", m.Name, m.Director, m.YearReleased)
}

var Movies = []Movie{
	Movie{
		ID:           1,
		Name:         "Dances with wolves",
		Director:     "Kevin Costner",
		YearReleased: 1979,
	},
	Movie{
		ID:           2,
		Name:         "Taken",
		Director:     "Liam Nieson",
		YearReleased: 2012,
	},
	Movie{
		ID:           3,
		Name:         "Bad Boys",
		Director:     "Will Smith",
		YearReleased: 1994,
	},
	Movie{
		ID:           4,
		Name:         "The Loard of the Rings",
		Director:     "Fruedo Baggins",
		YearReleased: 2012,
	},
	Movie{
		ID:           5,
		Name:         "Harry Poter and the Socerrers Stone",
		Director:     "Harry Poter",
		YearReleased: 2012,
	},
	Movie{
		ID:           6,
		Name:         "Joker",
		Director:     "Todd Philips",
		YearReleased: 2019,
	},
	Movie{
		ID:           7,
		Name:         "Shawshank Redemption",
		Director:     "Frank Darabont",
		YearReleased: 1994,
	},
	Movie{
		ID:           8,
		Name:         "Inglorious Bastards",
		Director:     "Quentin Tarantino",
		YearReleased: 2009,
	}, Movie{
		ID:           9,
		Name:         "No Country for old men",
		Director:     "Joel Coen",
		YearReleased: 2007,
	}, Movie{
		ID:           10,
		Name:         "Jaws",
		Director:     "Steven Spielberg",
		YearReleased: 1975,
	},
}
