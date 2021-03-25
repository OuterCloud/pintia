package pintia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeNumbersSort(t *testing.T) {
	// 准备待测数据
	var numbersStr = "-1 -4 -2 9 -999 8 0 6 4 5 3 2 1 7 8 9 9 100"
	var expectStr = "-999->-4->-2->-1->0->1->2->3->4->5->6->7->8->8->9->9->9->100"
	// 执行待测方法
	sortedNumberStr := NumbersSort(numbersStr)
	t.Log(sortedNumberStr)
	// 断言
	assert := assert.New(t)
	assert.Equal(expectStr, sortedNumberStr, "排序错误")
}
