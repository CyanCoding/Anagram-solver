package main

import gosort "github.com/cyancoding/go-sort"

func CreateHash(dictionary []string) map[string][]string {
	var hashMap = make(map[string][]string)

	for _, element := range dictionary {
		arr := hashMap[gosort.SortString(element)]
		hashMap[gosort.SortString(element)] = append(arr, element)
	}

	return hashMap

}
