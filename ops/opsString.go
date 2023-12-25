package ops

import (
	"regexp"
)

func FindSys(cmdStr string) []string {
	rx := regexp.MustCompile(`(sys\.)(\w+)(?:\(([^)]*)\))?`)
	matches := rx.FindStringSubmatch(cmdStr)
	return matches
}

func FindCmdIgnoreParams(cmdStr string) []string {
	rx := regexp.MustCompile(`cmd\(([^)]*)\)`)
	matches := rx.FindStringSubmatch(cmdStr)
	return matches
}

func FindLoadIgnoreParams(loadStr string) []string {
	rx := regexp.MustCompile(`load\(([^)]*)\)`)
	matches := rx.FindStringSubmatch(loadStr)
	return matches
}

func MakeSummaryIgnoreParams(loadStr string) []string {
	rx := regexp.MustCompile(`summary\(([^)]*)\)`)
	matches := rx.FindStringSubmatch(loadStr)
	return matches
}
