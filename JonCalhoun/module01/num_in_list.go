package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	if list == nil || len(list) == 0 {
		return false
	}

	for i := 0; i < len(list); i++ {
		if list[i] == num {
			return true
		}
	}

	return false
}

func NumInList2(list []int, num int) bool {
	if list == nil || len(list) == 0 {
		return false
	}
	// using a map; set as map of number to its index
	var mapList map[int]int
	for x, y := range list {
		mapList[y] = x
	}

	_, exist := mapList[num]

	return exist
}
