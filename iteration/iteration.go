package iteration

const defaultRepeatCount = 5

func Repeat(char string, counts ...int) string {
	output := ""
	count := defaultRepeatCount

	if len(counts) > 0 {
		count = counts[0]
	}

	for range count {
		output += char
	}

	return output
}
