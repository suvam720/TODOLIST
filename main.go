package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter the keyword as \n ADD:for adding task,\n REMOVE for deleting: ")
	list := []string{}
	for i := 0; i < 10; i++ {
		var keyword string
		fmt.Scan(&keyword)
		var no int
		switch keyword {
		case "ADD":
			fmt.Println("Enter total no of tasks: ")
			fmt.Scan(&no)
			reader := bufio.NewReader(os.Stdin)
			for no > 0 {
				fmt.Println("Enter the task: ")
				val, _ := reader.ReadString('\n')
				list = append(list, val)
				no--
			}
		case "REMOVE":
			var ind int
			fmt.Println("Enter the index: ")
			fmt.Scan(&ind)
			list = RemoveIndex(list, ind)
		}
		fmt.Println(list[:])

	}
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// // type add struct {
// // 	vill string
// // 	ps   string
// // 	pin  int
// // 	city string
// // }

// // func main() {
// // 	address := add{
// // 		vill: "officetila",
// // 		ps:   "Teliamura",
// // 		pin:  799205,
// // 		city: "Teliamura",
// // 	}
// // 	fmt.Println(address)
// // 	fmt.Println(address.vill)
// // 	fmt.Println(add{"Maiganga", "Teliamura", 700205, "Teliamura"})
// // 	address.city = "Agartala"
// // 	fmt.Println(address)
// // }

// // func main() {
// // 	marksheet := map[string]int{"science": 5, "history": 3, "politics": 9}
// // 	laptop := make(map[int]string)
// // 	laptop[1] = "dell"
// // 	laptop[2] = "hp"
// // 	fmt.Println(len(laptop))
// // 	fmt.Println(marksheet["history"])
// // 	delete(marksheet, "history")
// // 	fmt.Println(marksheet)
// // 	for key, element := range laptop {
// // 		fmt.Println("Key:", key, "=>", "Element:", element)
// // 	}
// // }

// // func main() {
// // 	var juice = []string{"lemon", "mango", "watermelon", "orange"}
// // 	score := []int{}
// // 	fmt.Println(juice[:])
// // 	score = append(score, 34, 5, 56, 567, 67)
// // 	score = remidx(score, 4)
// //  sort.Ints(score)
// // 	fmt.Println(score)
// // }
// // func remidx(data []int, index int) []int {
// // 	return append(data[:index], data[index+1:]...)
// // }

// // func main() {
// // 	var arr = [...]int{5, 4, 3, 2, 1}
// // 	var srr = [3]string{"ok", "not", "ok"}
// // 	vrr := &srr
// // 	vrr[1] = "ok"
// // 	//var array = make(int,)
// // 	// fmt.Println(reflect.ValueOf(arr).Kind())
// // 	// fmt.Println(reflect.ValueOf(srr).Kind())

// // 	arr[0] = 1
// // 	arr[1] = 2
// // 	arr[2] = 3
// // 	arr[3] = 4
// // 	arr[4] = 5

// // 	for i := range arr {
// // 		fmt.Println(arr[i])
// // 	}

// // 	for _, value := range vrr {
// // 		fmt.Println(value)
// // 	}
// // 	fmt.Println(check(srr, "lom"))

// // 	fmt.Printf("All elements: %v\n", arr[0:])
// // 	fmt.Printf("All elements: %v\n", arr[0:3])
// // 	fmt.Printf("All elements: %v\n", arr[len(arr):])

// // }
// // func check(array interface{}, values interface{}) bool {
// // 	arre := reflect.ValueOf(array)
// // 	for i := 0; i < arre.Len(); i++ {
// // 		if arre.Index(i).Interface() == values {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }
