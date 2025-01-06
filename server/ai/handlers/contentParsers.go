package aihandlers

import (
	"strings"
)

func ParseGeneratedContent(generatedContent string) (string, string) {
	lines := strings.Split(generatedContent, "\n")
	var title, content string

	for i, line := range lines {
		if strings.HasPrefix(line, "Title (up to 10 words):") {
			if i+1 < len(lines) {
				title = strings.TrimSpace(lines[i+1])
			}
		} else if strings.HasPrefix(line, "Content (up to 100 words):") {
			if i+1 < len(lines) {
				content = strings.TrimSpace(lines[i+1])
			}
		}
	}

	return title, content
}