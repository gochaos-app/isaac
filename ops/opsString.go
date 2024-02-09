package ops

import (
	"fmt"
	"regexp"
	"strings"
)

func FindInst(cmdStr string) []string {
	split := strings.Split(cmdStr, ":")
	return split
}

func FindCmd(IsaacStr, cmdStr string) string {
	rx := regexp.MustCompile(IsaacStr + `:(.*)`)
	matches := rx.FindStringSubmatch(cmdStr)
	if len(matches) > 1 {
		fmt.Println(matches[1])
	}
	return matches[1]
}

func CleanCmd(cmdStr string) string {
	rx := regexp.MustCompile(`[^a-zA-Z0-9\s\-_./]`)
	clean := rx.ReplaceAllString(cmdStr, " ")
	return clean
}

func FindSys(cmdStr string) []string {
	rx := regexp.MustCompile(`(sys\.)(\w+)(?:\(([^)]*)\))?`)
	matches := rx.FindStringSubmatch(cmdStr)
	return matches
}
