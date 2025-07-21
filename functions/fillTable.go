package functions

func FillTable(datasStrTable []string) map[rune][]string {
	letterAscii := make(map[rune][]string)
	startRune := rune(32)

	for i := 0; i+8 < len(datasStrTable); {

		if datasStrTable[i] == "" {
			i++
			continue
		}

		value := []string{}
		for j := 0; j < 8; j++ {
			value = append(value, datasStrTable[i+j])
		}
		letterAscii[startRune] = value
		startRune++
		i += 9

	}

	return letterAscii
}
