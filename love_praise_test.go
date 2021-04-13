package pintia

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Girl 女孩
type Girl struct {
	// Name 名字
	Name string
}

// AcceptConfess 接收表白
func (girl *Girl) AcceptConfess(boy Boy) {
	log.Printf("%+v接收了%+v的表白", girl.Name, boy.Name)
	boy.LovePraiseChan <- LovePraise{From: girl.Name, Words: "同意"}
}

// RefuseConfess 拒绝表白
func (girl *Girl) RefuseConfess(boy Boy) {
	log.Printf("%+v拒绝了%+v的表白", girl.Name, boy.Name)
	boy.LovePraiseChan <- LovePraise{From: girl.Name, Words: "不同意"}
}

// Boy 男孩
type Boy struct {
	// Name 名字
	Name string
	// Status 状态(单身或非单身)
	Status string
	// LoveChan 爱的表白通道
	LovePraiseChan chan LovePraise
}

// LovePraise 爱的表白
type LovePraise struct {
	From  string
	To    string
	Words string
}

// WaitForLoveCallBack 等待爱的回复
func (me *Boy) WaitForLoveCallBack() LovePraise {
	// 永葆单身
	for {
		select {
		case rec := <-me.LovePraiseChan:
			if rec.Words == "同意" {
				return LovePraise{
					From:  me.Name,
					To:    rec.From,
					Words: "我群发的，有人先同意了，下次",
				}
			} else if rec.Words == "不同意" {
				return LovePraise{
					From:  me.Name,
					To:    rec.From,
					Words: "我朋友拿我手机群发的",
				}
			}
		}
	}
}

// 测试爱的表达
func TestLovePraise(t *testing.T) {
	// 开始测试
	lucy := Girl{"Lucy"}
	lily := Girl{"Lily"}
	girls := []Girl{lucy, lily}
	me := Boy{Name: "我", LovePraiseChan: make(chan LovePraise), Status: "单身"}
	for i := 0; i < len(girls); i++ {
		go func() {
			// 异步等待爱的回复——不耽误正常吃饭睡觉
			callBack := me.WaitForLoveCallBack()
			log.Printf("%+v回复%+v道:%+v", callBack.From, callBack.To, callBack.Words)
		}()
	}
	lucy.AcceptConfess(me)
	lily.RefuseConfess(me)
	// 断言
	assert.Equal(t, "单身", me.Status)
}
