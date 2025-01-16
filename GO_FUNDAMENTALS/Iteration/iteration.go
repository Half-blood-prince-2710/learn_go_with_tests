package iteration




func Repeat(char string, count int) string {
	if count<0 {
		return "Repeat count should be greater than equal to 0"
	}
	var repeated string
	for i:=0; i<count; i++ {
		repeated += char
	}
	return repeated
}

