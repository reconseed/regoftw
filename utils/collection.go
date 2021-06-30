package utils

func ListContainsElement(list []string, element string) int {
	index := -1
	for i, listElement := range list {
		if listElement == element {
			index = i
			break
		}
	}
	return index
}
