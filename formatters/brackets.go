package formatters

// MakeSquareBrackets Берет строку  в квадратные скобики
func MakeSquareBrackets(str string) string {
	return "[" + str + "]"
}

// MakeParentheses Берет строку в круглые скобки
func MakeParentheses(str string) string {
	return "(" + str + ")"
}

// MakeSurround Окружает входящую строку подстрокой
func MakeSurround(input string, environment string) string {
	return environment + input + environment
}
