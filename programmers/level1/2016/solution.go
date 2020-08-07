package solution

func solution(a int, b int) string {
	dayOfWeek := [7]string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	index := 5
	mounth := 1
	day := 1
	for {
		if mounth == a && day == b {
			break
		}
		if mounth == 2 {
			if day < 29 {
				day++
			} else {
				mounth++
				day = 1
			}
		} else {
			if (mounth%2 == 0 && mounth < 8) || (mounth%2 != 0 && mounth >= 8) {
				if day < 30 {
					day++
				} else {
					mounth++
					day = 1
				}
			} else {
				if day < 31 {
					day++
				} else {
					mounth++
					day = 1
				}
			}
		}
		index++
		if index > 6 {
			index = 0
		}
	}
	return dayOfWeek[index]
}
