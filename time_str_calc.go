package pintia

import (
	"fmt"
	"time"
)

// TimeStrConverter 时间字符串转换器
type TimeStrConverter struct {
	InputDateFormat  string
	OutPutDateFormat string
}

// AddDuration 时间字符串加减天
// dateStr 是输入的时间字符串
// n 天数（负数为减，正数为加）
func (t *TimeStrConverter) AddDuration(dateStr string, n int) (string, error) {
	// 1. 将dateStr按指定格式转换成时间
	endDate, err := time.Parse(t.InputDateFormat, dateStr)
	if err != nil {
		return "", err
	}
	// 2. 加上n天——24*n小时
	durationStr := fmt.Sprintf("%+vh", 24*n)
	duration, _ := time.ParseDuration(durationStr)
	startDate := endDate.Add(duration)
	// 3. 将计算后的日期转成指定格式的时间字符串
	startDateStr := startDate.Format(t.OutPutDateFormat)
	return startDateStr, nil
}
