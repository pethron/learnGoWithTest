package hello

const spanish = "Spanish"
const french = "French"

const spanishGreeting = "Hola"
const englishGreeting = "Hello"
const frenchGreeting = "Bonjour"

func Hello(word, lang string) string {
	greeting := greetingFromLanguage(lang)
	if word == "" {
		word = "World"
	}
	return greeting + ", " + word
}

func greetingFromLanguage(language string) (greeting string) {

	switch language {
	case spanish:
		greeting = spanishGreeting
	case french:
		greeting = frenchGreeting
	default:
		greeting = englishGreeting
	}
	return
}
