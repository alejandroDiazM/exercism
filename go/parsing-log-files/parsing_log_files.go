package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	result := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR)\]`)

	return result.MatchString(text)
}

func SplitLogLine(text string) []string {
	result := regexp.MustCompile(`<[~*=-]*>`)

	return result.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	result := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0

	for _, line := range lines {
		if result.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	result := regexp.MustCompile(`end-of-line\d+`)

	return result.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)\s`)
	var result []string
	for _, line := range lines {
		items := re.FindStringSubmatch(line)
		if len(items) > 0 && len(items[1]) > 0 {
			line = `[USR] ` + items[1] + " " + line
		}
		result = append(result, line)
	}
	return result
}
