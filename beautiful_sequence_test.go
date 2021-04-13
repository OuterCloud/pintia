package pintia

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForExam(t *testing.T) {
	input := `3
3
0 1 -1
5
2 3 4 -2 -3
5
5 5 5 5 5`
	lines := strings.Split(input, "\n")
	testCasesCount, _ := strconv.Atoi(lines[0])
	lineIndex := 1
	rets := []string{}
	for i := 1; i <= testCasesCount; i++ {
		_, _ = strconv.Atoi(lines[lineIndex]) // 数组个数暂时没用
		lineIndex++
		ret := judgeCanBeSeqBeatiful(lines[lineIndex])
		if ret {
			rets = append(rets, "Yes")
		} else {
			rets = append(rets, "No")
		}
		lineIndex++
	}
	fmt.Print(strings.Join(rets, "\n"))
}

func TestJudgeCanBeSeqBeatiful(t *testing.T) {
	// 准备测试数据
	tests := []struct {
		caseDesc string
		str      string
		expect   bool
	}{
		{
			"测试不能变成美丽序列",
			"1 1 1 0 1 -1 1",
			false,
		},
		{
			"测试本身就是美丽序列",
			"1 0 -1 0 1 -1 1",
			true,
		},
		{
			"测试能够变成美丽序列",
			"-1 -1 -1 1 1 1",
			true,
		},
		{
			"测试考试测试数据",
			"0 1 -1",
			true,
		},
		{
			"测试考试测试数据",
			"2 3 4 -2 -3",
			true,
		},
		{
			"测试考试测试数据",
			"5 5 5 5 5",
			false,
		},
		{
			"测试自定义数据",
			"5 5 5 5 5 -1 -1 -1 -1 -1",
			true,
		},
		{
			"测试自定义数据",
			"0 0 0 0 0",
			true,
		},
		{
			"测试自定义数据",
			"1 1 1 1 1 1 1",
			false,
		}, {
			"测试自定义数据",
			"-1 -1 -1 -1 -1",
			false,
		}, {
			"测试超级数据",
			"0 0 0 0 0 1 1 1 1 1 1 1 -1 -1 -1 -1 -1",
			true,
		},
	}
	for _, test := range tests {
		// 执行待测方法
		ret := judgeCanBeSeqBeatiful(test.str)
		// 断言
		assert := assert.New(t)
		log := fmt.Sprintf("用例(%+v)失败", test.caseDesc)
		assert.Equal(test.expect, ret, log)
	}
}
