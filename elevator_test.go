package pintia

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChooseElevator(t *testing.T) {
	// 准备测试数据
	tests := []struct {
		caseID             int
		elevatorsManager   ElevatorsManager
		expectedElevatorID int
	}{
		{
			1,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				// 1到10号电梯都在-2层并且都处于空闲状态
				for i := 1; i <= 10; i++ {
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           IDLE,
						ready:            true,
						currentFloor:     -2,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期1号电梯被选择
			1,
		},
		{
			2,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					var status Status
					// 4号电梯使用中的情况
					if i == 4 {
						status = INUSE
					} else {
						status = IDLE
					}
					// 1到10号电梯分别在1到10层
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           status,
						ready:            true,
						currentFloor:     i,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期5号电梯被选择
			5,
		},
		{
			3,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					var status Status
					// 4号电梯和5号电梯使用中的情况
					if i == 4 || i == 5 {
						status = INUSE
					} else {
						status = IDLE
					}
					// 1到10号电梯分别在1到10层
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           status,
						ready:            true,
						currentFloor:     i,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期6号电梯被选择
			6,
		},
		{
			4,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					var status Status
					// 5号电梯使用中的情况
					if i == 5 {
						status = INUSE
					} else {
						status = IDLE
					}
					// 1到10号电梯分别在1到10层
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           status,
						ready:            true,
						currentFloor:     i,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期4号电梯被选择
			4,
		},
		{
			5,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					var status Status
					// 1-10号电梯都在使用中的情况
					status = INUSE
					// 1到10号电梯分别在1到10层
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           status,
						ready:            true,
						currentFloor:     i,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期没有电梯被选择
			0,
		},
		{
			6,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					var status Status
					// 1-9号电梯都在使用中的情况
					if 1 <= i && i <= 9 {
						status = INUSE
					}
					// 1到10号电梯分别在1到10层
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           status,
						ready:            true,
						currentFloor:     i,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期10号电梯被选择
			10,
		},
		{
			7,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					var status Status
					// 3-6号电梯都在使用中的情况
					if 3 <= i && i <= 6 {
						status = INUSE
					}
					// 1到10号电梯分别在1到10层
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           status,
						ready:            true,
						currentFloor:     i,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期7号电梯被选择
			7,
		},
		{
			8,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					// 全部空闲
					var status Status
					status = IDLE
					var currentFloor int
					// 1-5号电梯在5层
					if 1 <= i && i <= 5 {
						currentFloor = 5
					}
					// 6-10号电梯在4层
					if 6 <= i && i <= 10 {
						currentFloor = 4
					}
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           status,
						ready:            true,
						currentFloor:     currentFloor,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期1号电梯被选择
			1,
		},
		{
			9,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					// 全部空闲
					var status Status
					status = IDLE
					var currentFloor int
					// 1-5号电梯在-100层
					if 1 <= i && i <= 5 {
						currentFloor = -100
					}
					// 6-10号电梯在100层
					if 6 <= i && i <= 10 {
						currentFloor = 100
					}
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           status,
						ready:            true,
						currentFloor:     currentFloor,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期6号电梯被选择
			6,
		},
		{
			10,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					// 全部不健康
					var status Status
					status = UNHEALTHY
					var currentFloor int
					// 1-5号电梯在-100层
					if 1 <= i && i <= 5 {
						currentFloor = -100
					}
					// 6-10号电梯在100层
					if 6 <= i && i <= 10 {
						currentFloor = 100
					}
					elevators = append(elevators, Elevator{
						elevatorID:       i,
						status:           status,
						ready:            true,
						currentFloor:     currentFloor,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期0号电梯被选择
			0,
		},
		{
			11,
			func() (elevatorsManager ElevatorsManager) {
				elevators := []Elevator{}
				for i := 1; i <= 10; i++ {
					// 全部空闲
					var status Status
					status = IDLE
					var currentFloor int
					// 1-5号电梯在-100层
					if 1 <= i && i <= 5 {
						currentFloor = -100
					}
					// 6-10号电梯在100层
					if 6 <= i && i <= 10 {
						currentFloor = 100
					}
					elevators = append(elevators, Elevator{
						elevatorID: i,
						status:     status,
						// 电梯未就绪
						ready:            false,
						currentFloor:     currentFloor,
						destinationFloor: 0,
					})
				}
				// 被呼楼层5
				elevatorsManager = ElevatorsManager{
					elevators:   elevators,
					calledFloor: 5,
				}
				return
			}(),
			// 预期0号电梯被选择
			0,
		},
	}
	for _, test := range tests {
		// 执行待测方法
		chosenElevator := test.elevatorsManager.ChooseElevator()
		// 断言
		assert := assert.New(t)
		testPassed := assert.Equal(test.expectedElevatorID, chosenElevator.elevatorID)
		if !testPassed {
			t.Log(fmt.Sprintf("用例%+v未通过", test.caseID))
		}
	}
}
