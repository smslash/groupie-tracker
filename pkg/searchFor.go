package pkg

import (
	"strconv"
	"strings"
)

func SearchFor(value string) []Artists {
	var list []Artists

	for i, v := range Bands {
		find := false
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(value)) ||
			strings.Contains(strconv.Itoa(v.CreationDate), strings.ToLower(value)) ||
			strings.Contains(strings.ToLower(v.FirstAlbum), strings.ToLower(value)) {
			find = true
		}

		if !find {
			for _, j := range v.Members {
				if strings.Contains(strings.ToLower(j), strings.ToLower(value)) {
					find = true
					break
				}
			}
		}

		if !find {
			for j, k := range v.Relation.DatesLocations {
				if strings.Contains(strings.ToLower(j), strings.ToLower(value)) {
					find = true
					break
				}

				flag := false
				for x := 0; x < len(k); x++ {
					if strings.Contains(strings.ToLower(k[x]), strings.ToLower(value)) {
						find = true
						flag = true
						break
					}
				}

				if flag {
					break
				}
			}
		}

		if find {
			list = append(list, Bands[i])
		}
	}
	return list
}
