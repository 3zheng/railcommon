package railcommon

import (
	"fmt"
	"testing"
	protodefine "github.com/3zheng/railproto"
	proto "google.golang.org/protobuf/proto"
)

func TestAgent(t *testing.T) {
	//创建连接GATE的客户端
	loginReq := new(protodefine.LoginReq)
	loginReq.LoginAccount = "yourname"
	loginReq.LoginPassword = "E10ADC3949BA59ABBE56E057F20F883E" //123456的MD5加密
	ch := make(chan proto.Message, 1000)
	session := CreateClient("127.0.0.1:4101", ch)
	if session == nil {
		fmt.Println("连接gate失败")
		return
	}
	go ReceiveMsg(ch)
	//发送登录请求
}

// 循环阻塞读取
func ReceiveMsg(ch chan proto.Message) {
	fmt.Println("ReceiveMsg")
}
