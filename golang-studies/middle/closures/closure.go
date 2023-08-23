package closures

func Count() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
