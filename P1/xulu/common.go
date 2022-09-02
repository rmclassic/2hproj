package xulu

var xuluLetters = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
}

var xuluVerbs = map[string]func(int, int) int{
	"abcd": func(a, b int) int {
		return a + b
	},
	"bcde": func(a, b int) int {
		return a - b
	},
	"dede": func(a, b int) int {
		return a * b
	},
}
