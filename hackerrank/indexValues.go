package main

import (
	"fmt"
	"strings"
)

/*func GetOddStrings(data []string) []string {
	var result []string
	for _, value := range data {
		var oddData string
		for i, v := range value {
			if i%2 != 0 {
				oddData += string(v)
			}
		}
		result = append(result, oddData)
	}

	return result
}
func GetEvenStrings(data []string) []string {
	var result []string
	for _, value := range data {
		var evenData string
		for i, v := range value {
			if i%2 == 0 {
				evenData += string(v)
			}
		}
		result = append(result, evenData)
	}

	return result
}*/
func EvenOdd(data []string) []string {
	var result []string
	//var m []string
	fmt.Println(len(data))
	//temp := make([]string, 0)
	for _, value := range data {
		var evenOutput string
		var oddOutput string

		for j := 0; j < len(value); j++ {

			if j%2 == 0 {
				evenOutput += string(value[j])
			} else {
				oddOutput += string(value[j])
			}

		}
		//	temp = append(temp, evenOutput, oddOutput)
		//fmt.Println(temp)

		result = append(result, evenOutput, oddOutput)
	}

	return result
}
func main() {
	//var a []string
	a := EvenOdd([]string{"Hacker", "Rank"})
	fmt.Println(strings.NewReplacer("[", "", "]", "").Replace(fmt.Sprintf("%v", a)))
	//fmt.Println(strings.Join(strings.Split(strings.Trim(fmt.Sprintf("%v", a), "[]"), " "), "\n"))

}
