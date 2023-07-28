package pkg

var (
	Bands  []Artists
	Filter Filters
)

type Artists struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Coordinate   []Coordinates
	Relation     DatesLocations
	EncodeLoc    string
}

type Relation struct {
	Index []DatesLocations
}

type DatesLocations struct {
	DatesLocations map[string][]string
}

type GoogleGeoResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64
				Lng float64
			}
		}
	}
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Filters struct {
	MaxCreationDate   int
	MinCreationDate   int
	MaxFirstAlbumDate int
	MinFirstAlbumDate int
	MinMembers        int
	MaxMembers        int
}
