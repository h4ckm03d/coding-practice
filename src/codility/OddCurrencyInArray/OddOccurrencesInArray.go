package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	len := len(A)
	//fmt.Println(len/2)
	for i := 0; i <= len-2; i++ {
		//fmt.Println("#",i, i+2, len)

		if i+2 >= len {
			//fmt.Println("=>|", A[i])
			return A[i]
		} else if A[i] != A[i+2] {
			//fmt.Println("=>", A[i], A[i+2])
			return A[i]
		}

		//fmt.Println(A[i], A[i+2])
		if i > 0 && i%2 == 0 {
			i += 2
		}
	}

	return 0
}
