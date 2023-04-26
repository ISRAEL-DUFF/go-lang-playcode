package main

import "fmt"

func updateMap(arrayMap *map[string]int, item string) {
	v, prs := (*arrayMap)[item]
	if prs {
		(*arrayMap)[item] = v + 1
	} else {
		(*arrayMap)[item] = 1
	}
}

func arrayIntersection(arr1 []string, arr2 []string) []string {
	arrayMap := make(map[string]int)
	n := len(arr1)
	m := len(arr2)

	i := 0
	j := 0

	for {
		if i < n {
			updateMap(&arrayMap, arr1[i])

			i++
		}

		if j < m {
			updateMap(&arrayMap, arr2[j])

			j++
		}

		if m == j && n == i {
			break
		}
	}

	result := []string{}

	for k, v := range arrayMap {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}

func main() {
	fmt.Println(arrayIntersection([]string{"Green", "Orange", "Blue", "Red"}, []string{"Black", "Green", "Red", "Purpo"}))
}
