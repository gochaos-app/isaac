package ops

import (
	"regexp"
)

func FindSys(cmdStr string) []string {
	rx := regexp.MustCompile(`(sys\.)(\w+)`)
	matches := rx.FindStringSubmatch(cmdStr)
	return matches
}
