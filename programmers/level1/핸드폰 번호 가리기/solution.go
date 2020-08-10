package solution

func solution(phone_number string) string {
	number := []rune(phone_number)
	for i := 0; i < len(phone_number)-4; i++ {
		number[i] = '*'
	}
	return string(number)
}
