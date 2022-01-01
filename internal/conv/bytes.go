package conv

func BytesToStringWithoutDelim(bytes []byte) string {
	return string(BytesWithoutDelim(bytes))
}

func BytesWithoutDelim(bytes []byte) []byte {
	return bytes[:len(bytes)-1]
}
