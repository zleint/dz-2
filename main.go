package main

import (
	"errors"
	"fmt"
)

type smap = map[string]float64

func main() {
	var x string
	m := map[string]float64{
		"USD": 1.0,
		"EUR": 0.93,
		"RUB": 0.011,
	}
	for {
		frist_val, second_val, sum, err := userInput()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		result := converter(m, frist_val, second_val, sum)
		fmt.Println(result)
		fmt.Println("Желаете попробовать снова? y/n")
		fmt.Scan(&x)
		if x == "y" {
			continue
		}
		break
	}
}

func userInput() (string, string, float64, error) {
	var frist_val string
	var second_val string
	var sum float64
	fmt.Println("Введите валюту: ")
	fmt.Scan(&frist_val)
	fmt.Println("Введите вторую валюту: ")
	fmt.Scan(&second_val)
	fmt.Println("Введите количество валюты: ")
	fmt.Scan(&sum)
	if sum == 0.0 {
		return frist_val, second_val, sum, errors.New("Ошибка! Попробуйте снова")
	}
	return frist_val, second_val, sum, nil

}
func converter(m smap, frist_val string, second_val string, sum float64) float64 {
	course_x := m[frist_val]
	course_y := m[second_val]
	result := (sum * course_y) / course_x
	return result
}
