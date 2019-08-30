package main

import "fmt"

var GMap map[Course]string

type Student struct {
	name    string
	age     int
	courses []Course
}

type Course struct {
	name     string
	position string
}

type Item struct {
	name     string
	position string
}

func main() {
	GMap = make(map[Course]string)
	// 同一struct的object之间的比较
	Chinese := Course{"Chinese", "2-1"}
	English := Course{"Chinese", "2-1"}
	fmt.Println(Chinese == English)

	SA := Student{"A", 13, []Course{Chinese, English}}
	SB := Student{"A", 13, []Course{Chinese, English}}
	fmt.Println(&SA == &SB)
	// fmt.Println(SA==SB) // 非法语句

	//courseMap := make(map[Student]string) // runtime error, 因为Student中有不可比较的字段，所以不能作为map的key，但是如果换成*Student,就不会有错误

	assign()
	fmt.Println(GMap)
}

func assign() {
	Physics := Course{"Physics", "2-1"}
	GMap[Physics] = Physics.name
}
