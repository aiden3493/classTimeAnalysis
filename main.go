package main

import "fmt"

type dataType struct {
	name string
	data []string
}

func MonData() [23]dataType {
	data01 := dataType{
		name: "곽민우",
		data: []string{"△", "O", "X", "/", "O", "null", "null"},
	}

	data02 := dataType{
		name: "길현아",
		data: []string{"O", "O", "O", "null", "O", "null", "null"},
	}

	data03 := dataType{
		name: "김민재",
		data: []string{"O", "O", "O", "/", "O", "O", "null"},
	}

	data04 := dataType{
		name: "김승민",
		data: []string{"null", "O", "X", "△", "null", "null", "null"},
	}

	data05 := dataType{
		name: "김시은",
		data: []string{"null", "O", "O", "null", "O", "null", "null"},
	}

	data06 := dataType{
		name: "김지율",
		data: []string{"O", "O", "null", "null", "null", "null", "null"},
	}

	data07 := dataType{
		name: "김지휘",
		data: []string{"O", "null", "null", "X", "null", "null", "null"},
	}

	data08 := dataType{
		name: "박효은",
		data: []string{"null", "/", "O", "△", "O", "null", "null"},
	}

	data09 := dataType{
		name: "안범진",
		data: []string{"O", "null", "null", "X", "null", "null", "null"},
	}

	data10 := dataType{
		name: "엄하주",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data11 := dataType{
		name: "유도경",
		data: []string{"O", "null", "null", "/", "null", "null", "null"},
	}

	data12 := dataType{
		name: "유재이",
		data: []string{"O", "null", "null", "null", "null", "null", "null"},
	}

	data13 := dataType{
		name: "유지아",
		data: []string{"O", "O", "O", "/", "null", "O", "null"},
	}

	data14 := dataType{
		name: "이민호",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data15 := dataType{
		name: "이서연",
		data: []string{"null", "null", "null", "/", "null", "null", "null"},
	}

	data16 := dataType{
		name: "이승아",
		data: []string{"null", "null", "X", "/", "null", "□", "null"},
	}

	data17 := dataType{
		name: "이진영",
		data: []string{"null", "O", "null", "/", "O", "null", "null"},
	}

	data18 := dataType{
		name: "임설아",
		data: []string{"null", "null", "null", "/", "null", "null", "null"},
	}

	data19 := dataType{
		name: "장윤우",
		data: []string{"null", "O", "X", "null", "null", "null", "null"},
	}

	data20 := dataType{
		name: "정세희",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data21 := dataType{
		name: "최창민",
		data: []string{"null", "null", "null", "△", "O", "null", "null"},
	}

	data22 := dataType{
		name: "허정현",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data23 := dataType{
		name: "황우찬",
		data: []string{"null", "null", "null", "/", "null", "null", "null"},
	}

	generalizationData := [23]dataType{
		data01, data02, data03, data04, data05, data06, data07, data08, data09, data10,
		data11, data12, data13, data14, data15, data16, data17, data18, data19, data20,
		data21, data22, data23,
	}

	return generalizationData
}

func TueData() [23]dataType {

	data01 := dataType{
		name: "곽민우",
		data: []string{"△", "X", "△", "O", "O", "O", "null"},
	}

	data02 := dataType{
		name: "길현아",
		data: []string{"O", "X", "O", "O", "O", "O", "null"},
	}

	data03 := dataType{
		name: "김민재",
		data: []string{"/", "X", "O", "O", "O", "O", "null"},
	}

	data04 := dataType{
		name: "김승민",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data05 := dataType{
		name: "김시은",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data06 := dataType{
		name: "김지율",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data07 := dataType{
		name: "김지휘",
		data: []string{"X", "O", "O", "O", "O", "O", "null"},
	}

	data08 := dataType{
		name: "박효은",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data09 := dataType{
		name: "안범진",
		data: []string{"X", "△", "O", "O", "O", "O", "null"},
	}

	data10 := dataType{
		name: "엄하주",
		data: []string{"△", "O", "/", "O", "O", "O", "null"},
	}

	data11 := dataType{
		name: "유도경",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data12 := dataType{
		name: "유재이",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data13 := dataType{
		name: "유지아",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data14 := dataType{
		name: "이민호",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data15 := dataType{
		name: "이서연",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data16 := dataType{
		name: "이승아",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data17 := dataType{
		name: "이진영",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data18 := dataType{
		name: "임설아",
		data: []string{"O", "O", "X", "O", "O", "null", "null"},
	}

	data19 := dataType{
		name: "장윤우",
		data: []string{"X", "O", "/", "O", "O", "O", "null"},
	}

	data20 := dataType{
		name: "정세희",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data21 := dataType{
		name: "최창민",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data22 := dataType{
		name: "허정현",
		data: []string{"O", "O", "O", "O", "O", "O", "null"},
	}

	data23 := dataType{
		name: "황우찬",
		data: []string{"O", "O", "/", "O", "O", "O", "null"},
	}

	generalizationData := [23]dataType{
		data01, data02, data03, data04, data05, data06, data07, data08, data09, data10,
		data11, data12, data13, data14, data15, data16, data17, data18, data19, data20,
		data21, data22, data23,
	}

	return generalizationData
}

func ThuData() [23]dataType {
	data01 := dataType{
		name: "곽민우",
		data: []string{"O", "O", "null", "null", "O", "null", "△"},
	}

	data02 := dataType{
		name: "길현아",
		data: []string{"O", "O", "null", "null", "/", "null", "null"},
	}

	data03 := dataType{
		name: "김민재",
		data: []string{"O", "O", "X", "null", "△", "null", "null"},
	}

	data04 := dataType{
		name: "김승민",
		data: []string{"O", "O", "X", "O", "null", "null", "X"},
	}

	data05 := dataType{
		name: "김시은",
		data: []string{"O", "O", "X", "O", "null", "X", "X"},
	}

	data06 := dataType{
		name: "김지율",
		data: []string{"O", "O", "null", "null", "/", "X", "null"},
	}

	data07 := dataType{
		name: "김지휘",
		data: []string{"O", "O", "△", "null", "△", "X", "null"},
	}

	data08 := dataType{
		name: "박효은",
		data: []string{"O", "O", "X", "O", "null", "X", "null"},
	}

	data09 := dataType{
		name: "안범진",
		data: []string{"O", "O", "null", "null", "null", "X", "null"},
	}

	data10 := dataType{
		name: "엄하주",
		data: []string{"O", "O", "null", "null", "/", "null", "null"},
	}

	data11 := dataType{
		name: "유도경",
		data: []string{"O", "O", "null", "null", "/", "△", "△"},
	}

	data12 := dataType{
		name: "유재이",
		data: []string{"O", "O", "null", "O", "/", "null", "null"},
	}

	data13 := dataType{
		name: "유지아",
		data: []string{"O", "O", "/", "O", "null", "null", "null"},
	}

	data14 := dataType{
		name: "이민호",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data15 := dataType{
		name: "이서연",
		data: []string{"O", "O", "/", "null", "null", "null", "null"},
	}

	data16 := dataType{
		name: "이승아",
		data: []string{"O", "O", "null", "null", "null", "null", "null"},
	}

	data17 := dataType{
		name: "이진영",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data18 := dataType{
		name: "임설아",
		data: []string{"O", "O", "null", "null", "O", "/", "null"},
	}

	data19 := dataType{
		name: "장윤우",
		data: []string{"O", "O", "null", "O", "null", "△", "null"},
	}

	data20 := dataType{
		name: "정세희",
		data: []string{"O", "O", "null", "null", "null", "null", "null"},
	}

	data21 := dataType{
		name: "최창민",
		data: []string{"O", "O", "/", "null", "null", "null", "null"},
	}

	data22 := dataType{
		name: "허정현",
		data: []string{"O", "O", "null", "null", "null", "null", "null"},
	}

	data23 := dataType{
		name: "황우찬",
		data: []string{"O", "O", "null", "/", "null", "null", "null"},
	}

	generalizationData := [23]dataType{
		data01, data02, data03, data04, data05, data06, data07, data08, data09, data10,
		data11, data12, data13, data14, data15, data16, data17, data18, data19, data20,
		data21, data22, data23,
	}

	return generalizationData
}

func FriData() [23]dataType {
	data01 := dataType{
		name: "곽민우",
		data: []string{"X", "O", "O", "null", "null", "null", "null"},
	}

	data02 := dataType{
		name: "길현아",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data03 := dataType{
		name: "김민재",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data04 := dataType{
		name: "김승민",
		data: []string{"X", "O", "O", "null", "null", "null", "null"},
	}

	data05 := dataType{
		name: "김시은",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data06 := dataType{
		name: "김지율",
		data: []string{"O", "/", "O", "null", "null", "null", "null"},
	}

	data07 := dataType{
		name: "김지휘",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data08 := dataType{
		name: "박효은",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data09 := dataType{
		name: "안범진",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data10 := dataType{
		name: "엄하주",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data11 := dataType{
		name: "유도경",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data12 := dataType{
		name: "유재이",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data13 := dataType{
		name: "유지아",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data14 := dataType{
		name: "이민호",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data15 := dataType{
		name: "이서연",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data16 := dataType{
		name: "이승아",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data17 := dataType{
		name: "이진영",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data18 := dataType{
		name: "임설아",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data19 := dataType{
		name: "장윤우",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data20 := dataType{
		name: "정세희",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data21 := dataType{
		name: "최창민",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data22 := dataType{
		name: "허정현",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	data23 := dataType{
		name: "황우찬",
		data: []string{"O", "O", "O", "null", "null", "null", "null"},
	}

	generalizationData := [23]dataType{
		data01, data02, data03, data04, data05, data06, data07, data08, data09, data10,
		data11, data12, data13, data14, data15, data16, data17, data18, data19, data20,
		data21, data22, data23,
	}

	return generalizationData
}

func NextMonData() [23]dataType {
	data01 := dataType{
		name: "곽민우",
		data: []string{"O", "O", "O", "X", "O", "X", "X"},
	}

	data02 := dataType{
		name: "길현아",
		data: []string{"O", "O", "O", "X", "O", "X", "X"},
	}

	data03 := dataType{
		name: "김민재",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data04 := dataType{
		name: "김승민",
		data: []string{"O", "X", "O", "X", "O", "X", "O"},
	}

	data05 := dataType{
		name: "김시은",
		data: []string{"O", "O", "O", "X", "O", "X", "X"},
	}

	data06 := dataType{
		name: "김지율",
		data: []string{"X", "O", "O", "X", "X", "X", "O"},
	}

	data07 := dataType{
		name: "김지휘",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data08 := dataType{
		name: "박효은",
		data: []string{"O", "O", "O", "X", "O", "X", "X"},
	}

	data09 := dataType{
		name: "안범진",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data10 := dataType{
		name: "엄하주",
		data: []string{"O", "null", "null", "null", "null", "X", "X"},
	}

	data11 := dataType{
		name: "유도경",
		data: []string{"O", "X", "O", "X", "X", "null", "X"},
	}

	data12 := dataType{
		name: "유재이",
		data: []string{"O", "O", "X", "X", "X", "O", "O"},
	}

	data13 := dataType{
		name: "유지아",
		data: []string{"O", "O", "O", "O", "O", "X", "O"},
	}

	data14 := dataType{
		name: "이민호",
		data: []string{"X", "X", "null", "null", "null", "null", "null"},
	}

	data15 := dataType{
		name: "이서연",
		data: []string{"X", "O", "X", "X", "X", "X", "X"},
	}

	data16 := dataType{
		name: "이승아",
		data: []string{"null", "null", "null", "null", "null", "null", "null"},
	}

	data17 := dataType{
		name: "이진영",
		data: []string{"X", "null", "null", "null", "null", "null", "null"},
	}

	data18 := dataType{
		name: "임설아",
		data: []string{"□", "O", "O", "X", "X", "X", "X"},
	}

	data19 := dataType{
		name: "장윤우",
		data: []string{"△", "X", "null", "X", "O", "O", "O"},
	}

	data20 := dataType{
		name: "정세희",
		data: []string{"X", "X", "null", "null", "null", "X", "O"},
	}

	data21 := dataType{
		name: "최창민",
		data: []string{"O", "O", "null", "X", "X", "X", "△"},
	}

	data22 := dataType{
		name: "허정현",
		data: []string{"O", "O", "O", "X", "X", "X", "X"},
	}

	data23 := dataType{
		name: "황우찬",
		data: []string{"O", "O", "X", "X", "X", "X", "O"},
	}

	generalizationData := [23]dataType{
		data01, data02, data03, data04, data05, data06, data07, data08, data09, data10,
		data11, data12, data13, data14, data15, data16, data17, data18, data19, data20,
		data21, data22, data23,
	}

	return generalizationData
}

func main() {
	late := 0
	talk := 0
	sleep := 0
	bruise := 0
	question := 0

	MonData := MonData()
	TueData := TueData()
	ThuData := ThuData()
	FriData := FriData()
	NextMonData := NextMonData()

	for _, e := range MonData {
		for _, e2 := range e.data {
			switch e2 {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1
			}
		}
	}
	
	for _, e := range TueData {
		for _, e2 := range e.data {
			switch e2 {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1
			}
		}
	}

	for _, e := range ThuData {
		for _, e2 := range e.data {
			switch e2 {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1
			}
		}
	}

	for _, e := range FriData {
		for _, e2 := range e.data {
			switch e2 {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1
			}
		}
	}

	for _, e := range NextMonData {
		for _, e2 := range e.data {
			switch e2 {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1
			}
		}
	}

	fmt.Println("\n--------------------\n행동별 비율\n\n지각 :", late, "\n떠들기 :", talk, "\n눕기 :", sleep, "\n멍 때리기 :", bruise, "\n질문 :", question, "\n\n총합 :", late+talk+sleep+bruise+question, "\n--------------------")

	first := []int{}
	second := []int{}
	third := []int{}
	fourth := []int{}
	fifth := []int{}
	sixth := []int{}
	seventh := []int{}

	for i := 0; i <= 6; i++ {

		late := 0
		sleep := 0
		bruise := 0
		question := 0
		talk := 0

		for _, e := range MonData {
			switch e.data[i] {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1 
			}
		}

		for _, e := range TueData {
			switch e.data[i] {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1 
			}
		}

		for _, e := range ThuData {
			switch e.data[i] {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1 
			}
		}

		for _, e := range FriData {
			switch e.data[i] {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1 
			}
		}

		for _, e := range NextMonData {
			switch e.data[i] {
				case "□" :
					late += 1
				case "△" :
					sleep += 1
				case "/" :
					bruise += 1
				case "O" :
					question += 1
				case "X" :
					talk += 1 
			}
		}

		switch i {
			case 0:
				first = []int{late, sleep, bruise, question, talk}
			case 1:
				second = []int{late, sleep, bruise, question, talk}
			case 2:
				third = []int{late, sleep, bruise, question, talk}
			case 3:
				fourth = []int{late, sleep, bruise, question, talk}
			case 4: 
				fifth = []int{late, sleep, bruise, question, talk}
			case 5:
				sixth = []int{late, sleep, bruise, question, talk}
			case 6:
				seventh = []int{late, sleep, bruise, question, talk}
		}
	}

	fmt.Println("시간별 행동 비율")

	fmt.Println("firstClass")
	fmt.Println("\tlate :", first[0])
	fmt.Println("\tsleep :", first[1])
	fmt.Println("\tbruise :", first[2])
	fmt.Println("\tquestion :", first[3])
	fmt.Println("\ttalk :", first[4])

	fmt.Println("--------------------")

	fmt.Println("secondClass")
	fmt.Println("\tlate :", second[0])
	fmt.Println("\tsleep :", second[1])
	fmt.Println("\tbruise :", second[2])
	fmt.Println("\tquestion :", second[3])
	fmt.Println("\ttalk :", second[4])

	fmt.Println("--------------------")

	fmt.Println("thirdClass")
	fmt.Println("\tlate :", third[0])
	fmt.Println("\tsleep :", third[1])
	fmt.Println("\tbruise :", third[2])
	fmt.Println("\tquestion :", third[3])
	fmt.Println("\ttalk :", third[4])

	fmt.Println("--------------------")

	fmt.Println("fourthClass")
	fmt.Println("\tlate :", fourth[0])
	fmt.Println("\tsleep :", fourth[1])
	fmt.Println("\tbruise :", fourth[2])
	fmt.Println("\tquestion :", fourth[3])
	fmt.Println("\ttalk :", fourth[4])

	fmt.Println("--------------------")

	fmt.Println("fifthClass")
	fmt.Println("\tlate :", fifth[0])
	fmt.Println("\tsleep :", fifth[1])
	fmt.Println("\tbruise :", fifth[2])
	fmt.Println("\tquestion :", fifth[3])
	fmt.Println("\ttalk :", fifth[4])

	fmt.Println("--------------------")

	fmt.Println("sixthClass")
	fmt.Println("\tlate :", sixth[0])
	fmt.Println("\tsleep :", sixth[1])
	fmt.Println("\tbruise :", sixth[2])
	fmt.Println("\tquestion :", sixth[3])
	fmt.Println("\ttalk :", sixth[4])

	fmt.Println("--------------------")

	fmt.Println("seventhClass")
	fmt.Println("\tlate :", seventh[0])
	fmt.Println("\tsleep :", seventh[1])
	fmt.Println("\tbruise :", seventh[2])
	fmt.Println("\tquestion :", seventh[3])
	fmt.Println("\ttalk :", seventh[4])

	fmt.Println("--------------------")
}