package twentyfourPoints

import "fmt"

func calculateNumbers(n1 float64, n2 float64, op string) float64 {
	switch op {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	default:
		return n1 + n2
	}
}

func equals(a, b float64) bool {
	return a == b || (a-b < 1e-9 && a-b > -1e-9)
}

func Calculate24(n1, n2, n3, n4 int) bool {
	n := [4]int{n1, n2, n3, n4}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			for k := 0; k < 4; k++ {
				if k == i || k == j {
					continue
				}
				for l := 0; l < 4; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for _, op1 := range []string{"+", "-", "*", "/"} {
						for _, op2 := range []string{"+", "-", "*", "/"} {
							for _, op3 := range []string{"+", "-", "*", "/"} {

								if equals(24, calculateNumbers(calculateNumbers(float64(n[i]), float64(n[j]), op1), calculateNumbers(float64(n[k]), float64(n[l]), op2), op3)) {
									fmt.Println("(", n[i], op1, n[j], ")", op3, "(", n[k], op2, n[l], ")", "=", 24)
									return true
								}

								if equals(24, calculateNumbers(calculateNumbers(calculateNumbers(float64(n[i]), float64(n[j]), op1), float64(n[k]), op2), float64(n[l]), op3)) {
									fmt.Println("(", "(", n[i], op1, n[j], ")", op2, n[k], ")", op3, n[l], "=", 24)
									return true
								}

								if equals(24, calculateNumbers(float64(n[l]), calculateNumbers(calculateNumbers(float64(n[i]), float64(n[j]), op1), float64(n[k]), op2), op3)) {
									fmt.Println(n[l], op3, "(", "(", n[i], op1, n[j], ")", op2, n[k], ")", "=", 24)
									return true
								}

								if equals(24, calculateNumbers(calculateNumbers(float64(n[k]), calculateNumbers(float64(n[i]), float64(n[j]), op1), op2), float64(n[l]), op3)) {
									fmt.Println("(", n[k], op2, "(", n[i], op1, n[j], ")", ")", op3, n[l], "=", 24)
									return true
								}

								if equals(24, calculateNumbers(float64(n[l]), calculateNumbers(float64(n[k]), calculateNumbers(float64(n[i]), float64(n[j]), op1), op2), op3)) {
									fmt.Println(n[l], op3, "(", n[k], op2, "(", n[i], op1, n[j], ")", ")", "=", 24)
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}
