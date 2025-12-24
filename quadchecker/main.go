package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil || len(input) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	s := string(input)

	if s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}

	lines := splitLines(s)

	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	width := len(lines[0])
	height := len(lines)

	for _, line := range lines {
		if len(line) != width {
			fmt.Println("Not a quad function")
			return
		}
	}

	results := []string{}

	if isQuadA(lines, width, height) {
		results = append(results, fmt.Sprintf("[quadA] [%d] [%d]", width, height))
	}
	if isQuadB(lines, width, height) {
		results = append(results, fmt.Sprintf("[quadB] [%d] [%d]", width, height))
	}
	if isQuadC(lines, width, height) {
		results = append(results, fmt.Sprintf("[quadC] [%d] [%d]", width, height))
	}
	if isQuadD(lines, width, height) {
		results = append(results, fmt.Sprintf("[quadD] [%d] [%d]", width, height))
	}
	if isQuadE(lines, width, height) {
		results = append(results, fmt.Sprintf("[quadE] [%d] [%d]", width, height))
	}

	if len(results) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	for i, r := range results {
		if i > 0 {
			fmt.Print(" || ")
		}
		fmt.Print(r)
	}
	fmt.Println()
}

func splitLines(s string) []string {
	lines := []string{}
	current := ""

	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, current)
			current = ""
		} else if s[i] == '\r' {
			continue
		} else {
			current += string(s[i])
		}
	}

	if current != "" {
		lines = append(lines, current)
	}

	return lines
}

func checkQuad(l []string, w, h int, tl, tr, bl, br, edge byte) bool {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := l[y][x]

			switch {
			case y == 0 && x == 0:
				if c != tl {
					return false
				}
			case y == 0 && x == w-1:
				if c != tr {
					return false
				}
			case y == h-1 && x == 0:
				if c != bl {
					return false
				}
			case y == h-1 && x == w-1:
				if c != br {
					return false
				}
			case y == 0 || y == h-1:
				if c != edge {
					return false
				}
			case x == 0 || x == w-1:
				if c != '|' && edge != 'B' {
					return false
				}
				if edge == 'B' && c != 'B' {
					return false
				}
			default:
				if c != ' ' {
					return false
				}
			}
		}
	}
	return true
}

func isQuadA(l []string, w, h int) bool {
	if w == 1 && h == 1 {
		return l[0][0] == 'o'
	}
	return checkQuad(l, w, h, 'o', 'o', 'o', 'o', '-')
}

func isQuadB(l []string, w, h int) bool {
	if w == 1 && h == 1 {
		return l[0][0] == '/'
	}
	return checkQuad(l, w, h, '/', '\\', '\\', '/', '*')
}

func isQuadC(l []string, w, h int) bool {
	if w == 1 && h == 1 {
		return l[0][0] == 'A'
	}
	return checkQuad(l, w, h, 'A', 'A', 'C', 'C', 'B')
}

func isQuadD(l []string, w, h int) bool {
	if w == 1 && h == 1 {
		return l[0][0] == 'A'
	}
	return checkQuad(l, w, h, 'A', 'C', 'A', 'C', 'B')
}

func isQuadE(l []string, w, h int) bool {
	if w == 1 && h == 1 {
		return l[0][0] == 'A'
	}
	return checkQuad(l, w, h, 'A', 'C', 'C', 'A', 'B')
}
