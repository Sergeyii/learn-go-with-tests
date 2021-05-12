package iteration

func Repeat(name string) string {
	result := ""
	for i := 0; i < 5; i++ {
		result += name
	}

	return result
}
