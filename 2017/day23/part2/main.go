// b = 65
// c = 65
// Jump by 2
// ---
// b = 6500
// b = 106500
// c = b
// c = c + 17000
// f = 1
// d = 2
// e = 2
// g = d
// g = g * e
// g = g - b
// g(0), Jump to 17 ((2*e == 106500))
// f = 0
// e = e + 1
// g = e
// g = g - b
// g(0), jump to 12 (e == b)
// d = d + 1
// g = d
// g = g - b
// g(0), jump to 11
// f(0), jump to 27
// h = h + 1
// g = b
// g = g - c
// g(0), jump to 30
// End
// b = b + 17
// jump to 9

// After 20
// a = 1
// b = 106500
// c = 123500
// d = 2
// e = 106500
// f = 0

// After 24
// d = 106500
//

package main

import "fmt"

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func countCompositeBetween(a, b, inc int) (count int) {
	for i := a; i <= b; i += inc {
		if !isPrime(i) {
			count++
		}
	}
	return
}

func main() {
	// b := 0
	// c := 0
	// d := 0
	// e := 0
	// f := 0
	// h := 0

	// b = 106500
	// c = 123500

	// for {
	// 	f = 1
	// 	d = 2
	// 	for {
	// 		e = 2
	// 		for {
	// 			if d*e == b {
	// 				f = 0
	// 			}
	// 			e++
	// 			if e == b {
	// 				break
	// 			}
	// 		}
	// 		d++
	// 		if d == b {
	// 			break
	// 		}
	// 	}
	// 	if f == 0 {
	// 		h = h + 1
	// 	}
	// 	if c == b {
	// 		break
	// 	}
	// 	b = b + 17
	// }
	// fmt.Printf("Answer: %d\n", h)

	fmt.Printf("Answer: %d", countCompositeBetween(106500, 123500, 17))

}
