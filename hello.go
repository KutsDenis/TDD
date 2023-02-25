package main

import "fmt"

const kazakh = "Kazakh"
const french = "French"
const greetingPrefixEN = "Hello "
const greetingPrefixKZ = "Ассаламу алейкум "
const greetingPrefixFR = "Bonjour "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case kazakh:
		prefix = greetingPrefixKZ
	case french:
		prefix = greetingPrefixFR
	default:
		prefix = greetingPrefixEN
	}
	return
}

func main() {
	fmt.Print(Hello("Go", ""))
}
