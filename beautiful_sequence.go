package pintia

import (
	"fmt"
	"strconv"
	"strings"
)

// 判断整数数组是美丽序列
func isBeautifulSequence(seq []int) bool {
	for i := 0; i < len(seq); i++ {
		if i+1 >= len(seq) {
			break
		}
		if seq[i]*seq[i+1] > 0 {
			return false
		}
	}
	return true
}

// 数字附加属性结构体
type numPlus struct {
	num       int
	isDeleted bool
}

// 初始化数组附加结构
func initSeq(seq []int) (seqPlus []numPlus) {
	for i := 0; i < len(seq); i++ {
		seqPlus = append(seqPlus, numPlus{
			num:       seq[i],
			isDeleted: false,
		})
	}
	return
}

// 按美丽序列规则重新排列数组
func reorderBeatifulSeq(seqPlus []numPlus) (reorderedSeq []int) {
	// 按正数负数0的顺序排序争取最后用0才是最优解因为0是擦边球
	var positiveNums []numPlus
	var zeros []numPlus
	var negativeNums []numPlus
	for i := 0; i < len(seqPlus); i++ {
		if seqPlus[i].num > 0 {
			positiveNums = append(positiveNums, seqPlus[i])
		} else if seqPlus[i].num == 0 {
			zeros = append(zeros, seqPlus[i])
		} else {
			negativeNums = append(negativeNums, seqPlus[i])
		}
	}
	var index = 0
	posIndex, negaIndex, zeroIndex := 0, 0, 0
	for true {
		if posIndex >= len(positiveNums) && negaIndex >= len(negativeNums) && zeroIndex >= len(zeros) {
			break
		}
		if posIndex < len(positiveNums) {
			reorderedSeq = append(reorderedSeq, positiveNums[index].num)
			posIndex++
		}
		if negaIndex < len(negativeNums) {
			reorderedSeq = append(reorderedSeq, negativeNums[index].num)
			negaIndex++
		} else if zeroIndex < len(zeros) {
			reorderedSeq = append(reorderedSeq, zeros[index].num)
			zeroIndex++
		}
	}
	return
}

// 将字符串转成整型序列
func parseStrToIntSeq(str string) (intArray []int) {
	strArray := strings.Split(str, " ")
	for i := 0; i < len(strArray); i++ {
		intNum, _ := strconv.Atoi(strArray[i])
		intArray = append(intArray, intNum)
	}
	return
}

// 判断字符串是否能够变成美丽序列
func judgeCanBeSeqBeatiful(str string) bool {
	intSeq := parseStrToIntSeq(str)
	if isBeautifulSequence(intSeq) {
		log := fmt.Sprintf("本身就是美丽序列:%+v", intSeq)
		fmt.Println(log)
		return true
	}
	intSeqPlus := initSeq(intSeq)
	reorderedSeq := reorderBeatifulSeq(intSeqPlus)
	if len(reorderedSeq) < len(intSeq) {
		log := fmt.Sprintf("%+v无法变成美丽序列", str)
		fmt.Println(log)
		return false
	}
	if isBeautifulSequence(reorderedSeq) {
		log := fmt.Sprintf("%+v变成了美丽序列:%+v", str, reorderedSeq)
		fmt.Println(log)
		return true
	}
	return false
}
