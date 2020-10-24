func reverseBits(num uint32) uint32 {
	var res uint32

	for i := 0; i < 32; i++ {
		res = num&1 + res<<1
		num = num >> 1
	}

	return res
}