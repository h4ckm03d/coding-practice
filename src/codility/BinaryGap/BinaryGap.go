package solution

func Solution(N int) int {
	maxGap := 0
    currentGap := 0
    
    var max = func (a, b int) int {
        if a < b {
            return b
        }
        return a
    }

	// Skip the tailing zero(s)
	for N > 0 && N%2 == 0 {
		N = N / 2
	}

	for N > 0 {
		remainder := N % 2
		if remainder == 0 {
			// Inside a gap
			currentGap = currentGap + 1
		} else {
			// Gap ends
			if currentGap != 0 {
				maxGap = max(currentGap, maxGap)
				currentGap = 0
			}
		}
		N = N / 2
	}
	return maxGap
}
