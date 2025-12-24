package main

func CamelToSnakeCase(s string) string {
	snake := []byte{}

	for i := 0; i < len(s); i++ {
		snake = append(snake, s[i])
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {

			if s[i] >= 'A' && s[i] <= 'Z' && i+1 < len(s) && s[i+1] >= 'A' && s[i+1] <= 'Z' {
				return s
			}
			if s[i] >= 'a' && s[i] <= 'z' && i+1 < len(s) && s[i+1] >= 'A' && s[i+1] <= 'Z' {
				snake = append(snake, '_')
			}
			if s[i] >= 'A' && s[i] <= 'Z' && i == len(s)-1 {
				return s
			}
		} else {
			return s
		}
	}
	return string(snake)
}

// func main() {
// 	fmt.Println(CamelToSnakeCase("HelloWorlD"))
// 	fmt.Println(CamelToSnakeCase("helloWorld"))
// 	fmt.Println(CamelToSnakeCase("camelCase"))
// 	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
// 	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
// 	fmt.Println(CamelToSnakeCase("heY2"))
// }
