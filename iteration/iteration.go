package iteration

//const count = 5

func Repeat(value string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated = repeated + value
		//repeated +=value
	}
	return repeated
}
