package aihandlers

import (
	"strings"
)

func ParseContent(generatedText string) (string, string) {
    lines := strings.Split(generatedText, "\n")

    var title, content string   
	 
    for _, line := range lines {
        if strings.HasPrefix(line, "Title:") {
            title = strings.TrimPrefix(line, "Title: ")
            title = strings.Trim(title, "\"")
        } else if strings.HasPrefix(line, "Content:") {
            content = strings.TrimPrefix(line, "Content: ")
        }
    }

    return title, content
}