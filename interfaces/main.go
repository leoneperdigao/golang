package main

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	s := square{sideLength: 10}
	t := triangle{base: 10, height: 10}

	printGreeting(eb)
	printGreeting(sb)

	printArea(s)
	printArea(t)
}
