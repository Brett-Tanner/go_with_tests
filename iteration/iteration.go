package iteration

func Repeat(char string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += char
	}
	return repeated
}
