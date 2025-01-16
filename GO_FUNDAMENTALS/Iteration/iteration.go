package iteration



func Repeat(r string) string {
	var repeated string
	for i:=1; i<6; i++ {
		repeated= repeated + r
	}
	return repeated
}