package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	var buf [50]rune
	j := 0
	for _, r := range "x = " {
		buf[j] = r
		j++
	}
	n := points.x
	if n < 0 {
		buf[j] = '-'
		j++
		n = -n
	}
	if n == 0 {
		buf[j] = '0'
		j++
	} else {
		mag := 1
		t := n
		for t >= 10 {
			t /= 10
			mag *= 10
		}
		for mag > 0 {
			digit := (n / mag) % 10
			r := '0'
			for i := 0; i < digit; i++ {
				r++
			}
			buf[j] = r
			j++
			mag /= 10
		}
	}
	for _, r := range ", y = " {
		buf[j] = r
		j++
	}
	n = points.y
	if n < 0 {
		buf[j] = '-'
		j++
		n = -n
	}
	if n == 0 {
		buf[j] = '0'
		j++
	} else {
		mag := 1
		t := n
		for t >= 10 {
			t /= 10
			mag *= 10
		}
		for mag > 0 {
			digit := (n / mag) % 10
			r := '0'
			for i := 0; i < digit; i++ {
				r++
			}
			buf[j] = r
			j++
			mag /= 10
		}
	}
	buf[j] = '\n'
	j++
	for i := 0; i < j; i++ {
		z01.PrintRune(buf[i])
	}
}
