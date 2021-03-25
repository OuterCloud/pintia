package pintia

import (
	"strconv"
	"strings"
)

// SortedNumber 具有排序属性的数字结构
type SortedNumber struct {
	Self  int
	Left  *SortedNumber
	Right *SortedNumber
}

var sortedNumbers []string

// NumbersSort 将输入的任意n个整数从小到大输出
// 输入格式:输入在一行中给出n个整数，其间以空格分隔
// 输出格式:在一行中将n个整数从小到大输出，其间以“->”相连
func NumbersSort(threeNumsStr string) (sortedNumbersStr string) {
	nums := strings.Split(threeNumsStr, " ")
	intNumber, _ := strconv.Atoi(nums[0])
	sortedNumber := SortedNumber{Self: intNumber}
	for i := 1; i < len(nums); i++ {
		intNumber, _ := strconv.Atoi(nums[i])
		InsertNumber(&sortedNumber, intNumber)
	}
	PrintSortedNumber(sortedNumber)
	sortedNumbersStr = strings.Join(sortedNumbers, "->")
	return
}

// PrintSortedNumber 打印已排序数字
func PrintSortedNumber(sortedNumber SortedNumber) {
	if sortedNumber.Left != nil {
		PrintSortedNumber(*sortedNumber.Left)
	}
	sortedNumbers = append(sortedNumbers, strconv.Itoa(sortedNumber.Self))
	if sortedNumber.Right != nil {
		PrintSortedNumber(*sortedNumber.Right)
	}
}

// InsertNumber 插入数字
func InsertNumber(sortedNumber *SortedNumber, intNumber int) {
	if intNumber >= sortedNumber.Self {
		// fmt.Printf("%+v大于%+v\n", intNumber, sortedNumber.Self)
		if sortedNumber.Right == nil {
			sortedNumber.Right = &SortedNumber{
				Self: intNumber,
			}
		} else {
			InsertNumber(sortedNumber.Right, intNumber)
		}
	} else {
		// fmt.Printf("%+v小于%+v\n", intNumber, sortedNumber.Self)
		if sortedNumber.Left == nil {
			sortedNumber.Left = &SortedNumber{
				Self: intNumber,
			}
		} else {
			InsertNumber(sortedNumber.Left, intNumber)
		}
	}
}
