package main

import (
	"fmt"
	"linkedlist/ll"
	"log"
)

func main() {

	list, err := ll.ArrayToCycledLL([]int{10, 20, 30, 40, 50, 70, 80, 90, 100, 110, 130, 156, 199}, 3)
	if err != nil {
		log.Fatal(err)
	}
	// ll.PrintLL(list)

	isCyclic, err := ll.CheckIfCyclicLL(list)
	if err != nil {
		log.Fatal(err)
	}
	if isCyclic {
		pos, err := ll.FindCyclicPos(list)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("the pos where cycle occurs is ", pos)
	} else {
		fmt.Println("no cycle")
	}

	if isCyclic {
		fmt.Println("attempting to remove cycle ....")
		fixedList, err := ll.RemoveCycleFromLL(list)
		if err != nil {
			fmt.Println("unable to remove cycle due to error ", err)
			return
		}
		fmt.Println("fixed list is ")
		ll.PrintLL(fixedList)
	}

	list = ll.ArrayToLL([]int{10, 3, 13, 17, 20, 30, 40, 50, 70, 80, 90, 100, 110, 130, 156, 199})
	list, err = ll.EvenFirstOddNext(list)
	if err != nil {
		log.Fatal(err)
	}
	ll.PrintLL(list)

}
