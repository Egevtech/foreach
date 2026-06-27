package foreach

func ForEach[T any](arr []T, hd func(int, T)) {
	for index, data := range arr {
		hd(index, data)
	}
}
