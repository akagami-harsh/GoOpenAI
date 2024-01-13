package prompts

import "os"

func InputPrompt() string {
	input := make([]byte, 0, 1024)

	for {
		b := make([]byte, 1)
		_, err := os.Stdin.Read(b)
		if err != nil {
			return string(input)
		}

		if b[0] == '\n' {
			return string(input)
		}

		input = append(input, b[0])
	}

}
