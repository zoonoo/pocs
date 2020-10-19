package popcount

// pc represents the number of 1-bits in the particular number 0-256 represented by a byte.
var pc [256]byte

// init initializes the pc array.
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Count returns the number of int.
func Count(x int64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

// CountLoop runs in loop, using pc.
func CountLoop(x int64) int {
	var res int
	for i := range [8]int{} {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

// CountBruteForce iterates through all 64 bits one by one, as suggested by Problem 2.4
func CountBruteForce(x int64) int {
	var res int
	for i := range [64]int{} {
		shifted := x >> i
		res += int(int64(1) & shifted)
	}
	return res
}
