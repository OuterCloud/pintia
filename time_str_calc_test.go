package pintia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateStrAddDuration(t *testing.T) {
	// 准备测试数据
	inputDateStr := "20210630"
	durationDays := -3
	timeStrConverter := TimeStrConverter{InputDateFormat: "20060102", OutPutDateFormat: "2006-01-02"}
	// 执行待测方法
	outputDateStr, err := timeStrConverter.AddDuration(inputDateStr, durationDays)
	t.Logf("outputDateStr->%+v", outputDateStr)
	// 断言
	assert.Nil(t, err)
	assert.Equal(t, "2021-06-27", outputDateStr)
}
