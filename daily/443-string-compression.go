import "strconv"

func compress(chars []byte) int {
	lenChars := len(chars)

	write, read := 0, 0

	// read读遍历；write表示压缩后字符串遍历
	for read < lenChars {
		count := 0
		chars[write] = chars[read]

		for read < lenChars && chars[write] == chars[read] {
			count++
			// 这里实现read的向后移动
			read++
		}

		if count > 1 {
			bytesCount := []byte(strconv.Itoa(count))
			for j := 0; j < len(bytesCount); j++ {
				write++
				chars[write] = bytesCount[j]
			}
		}
		// write一定<=read；所以不会超过index
		write++
	}

	return write
}