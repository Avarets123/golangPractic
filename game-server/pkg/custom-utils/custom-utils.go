package custom_utils

func Map[T any, T2 any](list []T, mapFn func(T) T2) []T2 {
	newList := []T2{}

	for _, data := range list {

		newList = append(newList, mapFn(data))

	}

	return newList
}

func Filter[T any](list []T, filterFn func(T) bool) []T {

	filteredList := []T{}

	for _, data := range list {

		if filterFn(data) {
			filteredList = append(filteredList, data)
		}

	}

	return filteredList

}

func Reduce[T any, T2 any](list []T, initAndAcc T2, reduceFn func(T2, T) T2) T2 {

	for _, data := range list {

		initAndAcc = reduceFn(initAndAcc, data)

	}
	return initAndAcc
}
