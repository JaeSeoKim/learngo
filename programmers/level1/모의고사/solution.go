package solution

func solution(answers []int) []int {
	result := [3]int{0}
	for i, answer := range answers {
		checkAnswer(i, answer, &result)
	}
	return checkWinner(result)
}

func checkWinner(result [3]int) []int {
	var returnValue []int
	maxValue := 0
	for _, value := range result {
		if maxValue <= value {
			maxValue = value
		}
	}
	for i, value := range result {
		if maxValue == value {
			returnValue = append(returnValue, i+1)
		}
	}
	return returnValue
}

func checkAnswer(i, answer int, result *[3]int) {
	answerTmp2 := [8]int{2, 1, 2, 3, 2, 4, 2, 5}
	answerTmp3 := [5]int{3, 1, 2, 4, 5}
	if (i%5 + 1) == answer {
		result[0]++
	}
	if (answerTmp2[i%8]) == answer {
		result[1]++
	}
	if (answerTmp3[(i%10)/2]) == answer {
		result[2]++
	}
}
