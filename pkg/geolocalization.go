package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func Geolocalization(v Artists) []Coordinates {
	for location := range v.Relation.DatesLocations {
		resp, err := http.Get("https://maps.googleapis.com/maps/api/geocode/json?address=" + url.QueryEscape(location) + "&key=" + apiKey)
		if err != nil {
			log.Fatalf("Failed to send request to Google Maps API: %v", err)
		}
		defer resp.Body.Close()

		var data GoogleGeoResponse

		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			log.Fatalf("Failed to decode API response: %v", err)
		}

		if len(data.Results) > 0 {
			coord := Coordinates{
				Latitude:  data.Results[0].Geometry.Location.Lat,
				Longitude: data.Results[0].Geometry.Location.Lng,
			}
			v.Coordinate = append(v.Coordinate, coord)
		}
	}

	return v.Coordinate
}
