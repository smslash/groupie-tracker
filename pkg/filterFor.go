package pkg

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func FilterFor(r *http.Request) []Artists {
	var list []Artists

	r.ParseForm()

	MinCreationDate, err := strconv.Atoi(r.FormValue("MinCreationDate"))
	if err != nil {
		log.Fatalln("Error on converting string to int:", err)
	}
	MaxCreationDate, err := strconv.Atoi(r.FormValue("MaxCreationDate"))
	if err != nil {
		log.Fatalln("Error on converting string to int:", err)
	}
	MinFirstAlbumDate, err := strconv.Atoi(r.FormValue("MinFirstAlbumDate"))
	if err != nil {
		log.Fatalln("Error on converting string to int:", err)
	}
	MaxFirstAlbumDate, err := strconv.Atoi(r.FormValue("MaxFirstAlbumDate"))
	if err != nil {
		log.Fatalln("Error on converting string to int:", err)
	}
	Members := []int{}
	for i := 0; i < Filter.MaxMembers; i++ {
		if r.FormValue("Member"+strconv.Itoa(i+1)) == "on" {
			Members = append(Members, i+1)
		}
	}
	Locations := []string{}
	for _, loc := range r.Form["locations"] {
		Locations = append(Locations, loc)
	}

	for _, v := range Bands {
		find := false
		if v.CreationDate >= MinCreationDate && v.CreationDate <= MaxCreationDate {
			FirstAlbum, err := strconv.Atoi(v.FirstAlbum[6:])
			if err != nil {
				log.Fatal("Error on converting date string to date int")
			}

			if FirstAlbum >= MinFirstAlbumDate && FirstAlbum <= MaxFirstAlbumDate {
				if len(Members) == 0 {
					if len(Locations) == 0 {
						find = true
					} else {
						count := 0

						for _, loc := range Locations {
							for key := range v.Relation.DatesLocations {
								loc = strings.ToLower(strings.ReplaceAll(loc, ",", ""))
								key = strings.ReplaceAll(key, "_", " ")
								key = strings.ReplaceAll(key, "-", " ")
								if strings.Contains(strings.ToLower(key), strings.ToLower(loc)) {
									count++
									break
								}
							}
						}

						if count == len(Locations) {
							find = true
						}
					}
				} else {
					if len(Locations) == 0 {
						for _, num := range Members {
							if num == len(v.Members) {
								find = true
							}
						}
					} else {
						for _, num := range Members {
							if num == len(v.Members) {
								count := 0

								for _, loc := range Locations {
									for key := range v.Relation.DatesLocations {
										loc = strings.ToLower(strings.ReplaceAll(loc, ",", ""))
										key = strings.ReplaceAll(key, "_", " ")
										key = strings.ReplaceAll(key, "-", " ")
										if strings.Contains(strings.ToLower(key), strings.ToLower(loc)) {
											count++
											break
										}
									}
								}

								if count == len(Locations) {
									find = true
								}
							}
						}
					}
				}
			}
		}

		if find {
			list = append(list, v)
		}
	}

	return list
}
