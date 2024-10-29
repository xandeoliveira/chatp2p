package cipher

func cipher(message string, k byte) string {
	byte_message := []byte(message)
	buffer := make([]byte, len(byte_message))

	for index, char_byte := range byte_message {
		buffer[index] = char_byte + k
	}

	return string(buffer)
}

func Encoding(message string, k byte) string {
	return cipher(message, k)
}

func Decoding(message string, k byte) string {
	return cipher(message, -k)
}
