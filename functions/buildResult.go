package functions

import (
	"os"
	"path/filepath"
	"strings"
)

func BuildResult(text, style, color string) (strings.Builder, error) {
	var result strings.Builder

	bannerPath := filepath.Join("banners", style+".txt")
	if _, err := os.Stat(bannerPath); os.IsNotExist(err) {
		// Return an error instead of calling ServeError
		return result, err
	}

	bannerData := ReadFile(bannerPath)

	lines := strings.Split(string(bannerData), "\n")
	letterAscii := FillTable(lines)

	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range line {
				if asciiLines, ok := letterAscii[char]; ok {
					result.WriteString(asciiLines[i])
				}
			}
			result.WriteString("\n")
		}
	}

	return result, nil
}
