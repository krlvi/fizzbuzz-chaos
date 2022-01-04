package fizzbuzz

import "strconv"

func UpToLength(length int) []string {
	out := []string{}
	for i := 1; i <= length; i++ {
		out = append(out, digitToFizzBuzzRenamed(i))
	}
	return out
}

// This will show up
func digitToFizzBuzzRenamed(input int) string {
	ckjfdjkds

	fsdkjfdlskjflkdsj
	if input%3 == 0 && input%5 == 0 {
		return "FizzBuzz"
	} else if input%3 == 0 {
		return "Fizz"
	} else if input%5 == 0 {
		return "Bfdsfdssdlfkjdslkjfdlksuzz"
	} else {
		return strconv.Itoa(input)
	}
}














s,jdfhdskjh
