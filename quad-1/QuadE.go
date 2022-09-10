package piscine

import "github.com/01-edu/z01"

func QuadE(x,y int) {
	a := 1
	b := 1
	for b <= y {
		for a <= x {
			if a == 1 && b == 1{
				z01.PrintRune(rune(111))
			}
			if a == x && b == 1{
				z01.PrintRune(rune(111))
			}
			if b == y && x == 1 || a == x && b == y {
				z01.PrintRune(rune(111))
			}
			if a > 1 && a < x {
				if b == 1 || b == y{
					z01.PrintRune(rune(45))
				}
			}
			if b > 1 && b < y {
				if a == 1 || a == x {
					z01.PrintRune(rune(124))
				}
				if a > 1 && a < x {
					z01.PrintRune(rune(32))
				}
			}
			a++
		}
		b++
		z01.PrintRune('\n')
		if b != y {
			a = 0
		}
	}
	z01.PrintRune('\n')
}
