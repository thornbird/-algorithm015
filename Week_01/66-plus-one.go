func plusOne(digits []int) []int {
    digitsLen := len(digits)
    for tmpIndex := digitsLen-1; tmpIndex >= 0; tmpIndex-- {
        if digits[tmpIndex] < 9 {
            digits[tmpIndex]++
            return digits
        }

        digits[tmpIndex] = 0
    }
    output := make([]int, digitsLen+1, digitsLen+1)
    output[0] = 1
    return output
}