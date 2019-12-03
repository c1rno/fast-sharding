package sharding

type CalculationAlgorithm func(current uint32, max uint32) uint32

func Sharder(algorithm CalculationAlgorithm, maximum uint32) func(uint32) uint32 {
	return func(current uint32) uint32 {
		return algorithm(current, maximum)
	}
}

func Modulo(current uint32, max uint32) uint32 {
	return current % max
}

// same as here https://github.com/golang/go/commit/46a75870ad5b9b9711e69ffce3738a3ab2057789
func PowerWithShift(current uint32, max uint32) uint32 {
	return uint32(uint64(current) * uint64(max) >> 32)
}
