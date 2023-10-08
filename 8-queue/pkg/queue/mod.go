package queue

// Because the default Go modulus is stupid
func mod(a, b int) int {
	return (a%b + b) % b
}
