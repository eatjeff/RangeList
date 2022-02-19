package main

import (
	"errors"
	"fmt"
)

type RangeList struct {
	//TODO: implement
	data [][]int
	dict map[int]int
}

/** Initialize RangeList */
func Constructor() RangeList {
	return RangeList{
		data: [][]int{},
		dict: make(map[int]int),
	}
}

var minEle int
var maxEle int

func (rangeList *RangeList) Add(rangeElement [2]int) error {
	//TODO: implement

	if len(rangeList.data) == 0 {
		rangeList.data = append(rangeList.data, rangeElement[:2])
		minEle = rangeElement[0]
		maxEle = rangeElement[1]
	} else {
		//判断左右边界值
		for _, ele := range rangeList.data {
			if ele[1] > maxEle {
				maxEle = ele[1]
			}

			if ele[0] < minEle {
				minEle = ele[0]

			}
		}
		//fmt.Printf("new add element is : %v\n", rangeElement)
		//fmt.Printf("current add element min : %v\n", minEle)
		//fmt.Printf("current add element min : %v\n", maxEle)
		if rangeElement[0] > maxEle && rangeElement[1] > maxEle {
			rangeList.data = append(rangeList.data, rangeElement[:2])
			fmt.Println("add function: ", rangeList.data)
		} else {
			//新元素最小值和当前最大值相等
			if rangeElement[0] == maxEle {
				if rangeElement[1] == maxEle {
					return nil
				} else {
					rangeList.data[len(rangeList.data)-1][1] = rangeElement[1]
				}
			}
			//新元素最小值比当前最小值大，比当前最大值小
			if rangeElement[0] > minEle && rangeElement[1] < maxEle {
				//比较每个区间最大值
				for _, ele := range rangeList.data {
					if rangeElement[0] > ele[0] && rangeElement[1] < ele[1] {
						return nil
					}
					if rangeElement[0] > ele[0] && rangeElement[1] > ele[1] {
						ele[1] = rangeElement[1]
					}
				}
			}
		}
	}
	return nil
}

func (rangeList *RangeList) Remove(rangeElement [2]int) error {
	//TODO: implement
	if len(rangeList.data) == 0 {
		return errors.New("current rangeList data does not contain any data")
	}

	if rangeElement[0] > minEle && rangeElement[1] < maxEle {
		for _, ele := range rangeList.data {
			if rangeElement[0] > ele[1] {
				continue
			}
			if rangeElement[0] == ele[0] && rangeElement[1] < ele[1] {
				ele[0] = rangeElement[1]
				//continue
				break
			}
			if rangeElement[0] > ele[0] && rangeElement[1] < ele[1] {
				rangeList.data = rangeList.data[0 : len(rangeList.data)-1]

				firstSection := []int{ele[0], rangeElement[0]}
				secondSection := []int{rangeElement[1], ele[1]}
				rangeList.data = append(rangeList.data, firstSection)
				rangeList.data = append(rangeList.data, secondSection)
				break
			}
			if rangeElement[0] < ele[1] && rangeElement[1] < maxEle {
				rangeList.data = rangeList.data[0:1]
				rangeList.data[0][1] = rangeElement[0]

				secondSection := []int{rangeElement[1], maxEle}
				rangeList.data = append(rangeList.data, secondSection)
				break

			}
		}
	}

	return nil
}

func (rangeList *RangeList) Print() error {
	//TODO: implement
	fmt.Println(rangeList.data)
	return nil
}

func main() {

	//fmt.Println("hello world2")
	rl := RangeList{}
	rl.Add([2]int{1, 5})
	rl.Print()
	//    Should display: [1, 5)
	rl.Add([2]int{10, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 21})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{2, 4})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{3, 8})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 10})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 11})
	rl.Print()
	//Should display: [1, 8) [11, 21)
	rl.Remove([2]int{15, 17})
	rl.Print()
	// Should display: [1, 8) [11, 15) [17, 21)
	rl.Remove([2]int{3, 19})
	rl.Print()
	// Should display: [1, 3) [19, 21)
}
