package cipher

func cipher(message string, k int) string {
	byte_message := []byte(message)
	buffer := make([]byte, len(byte_message))

	for index, char_byte := range byte_message {
		buffer[index] = byte((int(char_byte) + k + 256) % 256)
	}

	return string(buffer)
}

func Encoding(message string, k int) string {
	return cipher(message, k)
}

func Decoding(message string, k int) string {
	return cipher(message, -k)
}
