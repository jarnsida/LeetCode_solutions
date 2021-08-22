func convert(s string, numRows int) string {

	var i, j int = 0, 0
	step := 2*numRows - 2
	var answer strings.Builder

	answer.Grow(len(s))

	// answer.WriteByte(s[0])
	if numRows == 1 {
		return s
	}

	for {
		if i == 0 || i == numRows-1 {
			for j = 0; (j*step + i) < len(s); j++ {

				answer.WriteByte(s[j*step+i])
				fmt.Println("j=", j, "i=", i, "ans=", answer.String())
			}
		}

		if answer.Len() >= len(s) {
			break
		}

		if i != 0 {
			for j = 0; (j*step - i) < len(s); j++ {
				//   fmt.Println("second case")
				if j == 0 {

					answer.WriteByte(s[i])
					fmt.Println("add j", "j=", j, "i=", i, "ans=", answer.String())
				} else if (j*step + i) < len(s) {

					answer.WriteByte(s[j*step-i])
					answer.WriteByte(s[j*step+i])
					fmt.Println("add 2j", "j=", j, "i=", i, "ans=", answer.String())

				} else if (j*step-i) < len(s) && (j*step+i) >= len(s) {

					answer.WriteByte(s[j*step-i])
					fmt.Println("j=", j, "i=", i, "ans=", answer.String())
				}
			}

		}

		i++
	}
	return answer.String()
}
