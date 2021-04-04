package pintia

import "math"

// Status 状态
type Status int

const (
	// IDLE 空闲状态
	IDLE Status = 0
	// INUSE 使用中状态
	INUSE Status = 1
	// UNHEALTHY 不健康
	UNHEALTHY Status = 2
)

// Elevator 电梯
type Elevator struct {
	status Status
	// 电梯当前楼层
	currentFloor int
	// 电梯目标楼层
	destinationFloor int
	// 电梯ID
	elevatorID int
	// 是否就绪
	ready bool
}

// ElevatorsManager 电梯组
type ElevatorsManager struct {
	elevators []Elevator
	// 被呼楼层
	calledFloor int
}

// BetterElevator 更好的电梯
func BetterElevator(firEle Elevator, secEle Elevator, destinationFloor int) (chosenElevator Elevator) {
	firEleDistance := math.Abs(float64(firEle.currentFloor - destinationFloor))
	secEleDistance := math.Abs(float64(secEle.currentFloor - destinationFloor))
	chosenElevator = firEle
	if secEleDistance < firEleDistance {
		chosenElevator = secEle
	}
	return
}

// ChooseElevator 选择电梯
func (elevatorsManager ElevatorsManager) ChooseElevator() (chosenElevator Elevator) {
	for _, elevator := range elevatorsManager.elevators {
		// 不选择不健康/正在使用中/未就绪的电梯
		if elevator.status == UNHEALTHY || elevator.status == INUSE || !elevator.ready {
			continue
		}
		// 如果还没选电梯并且当前电梯没有目标楼层就先选当前电梯
		if !chosenElevator.ready && elevator.destinationFloor == 0 {
			chosenElevator = elevator
		}
		// 根据被呼楼层在已选电梯和当前电梯间选择一个更适合的电梯
		chosenElevator = BetterElevator(chosenElevator, elevator, elevatorsManager.calledFloor)
	}
	return
}
