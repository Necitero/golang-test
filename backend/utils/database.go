package utils

func DBStringInsert(val string, last bool) string {
	var formattedString string = "'" + val + "'"
	if !last {
		formattedString += ","
	}
	return formattedString
}
