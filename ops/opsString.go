package ops

import (
	"regexp"
	"strings"
)

func FindInst(cmdStr string) []string {
	split := strings.Split(cmdStr, ":")
	return split
}

func ExtractCodeBlocks(input string) []string {
	rx := regexp.MustCompile("```(.+?)```")
	matches := rx.FindAllStringSubmatch(input, -1)

	var results []string
	for _, match := range matches {
		if len(match) > 1 {
			results = append(results, match[1])
		}
	}

	return results
}

func CleanCmd(cmdStr string) string {
	rx := regexp.MustCompile(`[^a-zA-Z0-9\s\-_./|]`)
	clean := rx.ReplaceAllString(cmdStr, " ")
	return clean
}

func FindSys(cmdStr string) []string {
	rx := regexp.MustCompile(`(sys\.)(\w+)(?:\(([^)]*)\))?`)
	matches := rx.FindStringSubmatch(cmdStr)
	return matches
}
