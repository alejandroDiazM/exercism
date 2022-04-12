package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, rune := range log {
		switch rune {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	result := ""
	for _, rune := range log {
		if rune == oldRune {
			result += string(newRune)
		} else {
			result += string(rune)
		}
	}
	return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	length := len([]rune(log))
	if length <= limit {
		return true
	} else {
		return false
	}
}
