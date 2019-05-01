package stack

import (
	"fmt"
	"strconv"
)

func openLock(deadends []string, target string) int {
	//start := "0000"
	dict := make(map[string]bool)
	for _,v := range deadends {
		//i, err := strconv.Atoi(v)
		//if err != nil {
		//	panic("Invalid input")
		//}
		dict[v] = true
	}
	//t, err := strconv.Atoi(target)
	//if err != nil {
	//	panic("Invalid input")
	//}


	fmt.Println(dict)
	res,err := compute(0,"0000", target, deadends, 0)
	if err!=nil {
		panic("Invalid input")
	}
	return res
}


func compute(start int, s, t string, deadends []string, count int) (int, error) {
	if s == t {
		return count,nil
	}
	//tmp := 0

	if start >= len(s) {
		return -1,nil
	}


	for i:=start; i < len(s); i++ {
		for n := 1; n < 10; n++ {
			num := int(s[i]) - 48
			num += n
			if num > 9 {
				break
			}
			//fmt.Printf("%s\n", s[:start]+strconv.Itoa(n)+s[start+1:])
			newStr := s[:i] + strconv.Itoa(num) + s[i+1:]
			//fmt.Println(newStr)
			c,_ := compute(i + 1, newStr, t, deadends, count+1)
			count += c
		}
	}

	return count,nil
}