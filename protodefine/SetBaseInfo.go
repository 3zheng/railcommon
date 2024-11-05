package protodefine

import (
	"fmt"
	"reflect"
)

// CMDKindId_IDKindNetTCP大类
func setTcpBase(input interface{}) (bool, *BaseInfo) {
	switch data := input.(type) {
	case *TCPTransferMsg:
		fmt.Println("报文类型为*TCPTransferMsg")
		if data.Base == nil {
			fmt.Println("分配了new( BaseInfo)")
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindNetTCP)
		data.Base.SubId = uint32(CMDID_Tcp_IDTCPTransferMsg)
		return true, data.Base
	case *TCPSessionCome:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindNetTCP)
		data.Base.SubId = uint32(CMDID_Tcp_IDTCPSessionCome)
		return true, data.Base
	case *TCPSessionClose:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindNetTCP)
		data.Base.SubId = uint32(CMDID_Tcp_IDTCPSessionClose)
		return true, data.Base
	case *TCPSessionKick:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindNetTCP)
		data.Base.SubId = uint32(CMDID_Tcp_IDTCPSessionKick)
		return true, data.Base
	default:
		return false, nil
	}
}

func setGateBase(input interface{}) (bool, *BaseInfo) {
	switch data := input.(type) {
	case *PulseReq:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindGate)
		data.Base.SubId = uint32(CMDID_Gate_IDPulseReq)
		return true, data.Base
	case *PulseRsp:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindGate)
		data.Base.SubId = uint32(CMDID_Gate_IDPulseRsp)
		return true, data.Base
	case *GateTransferData:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindGate)
		data.Base.SubId = uint32(CMDID_Gate_IDTransferData)
		return true, data.Base
	default:
		return false, nil
	}
}

func setRouterBase(input interface{}) (bool, *BaseInfo) {
	switch data := input.(type) {
	case *RouterTransferData:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindRouter)
		data.Base.SubId = uint32(CMDID_Router_IDTransferDataRt)
		return true, data.Base
	case *RegisterAppReq:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindRouter)
		data.Base.SubId = uint32(CMDID_Router_IDRegisterAppReq)
		return true, data.Base
	case *RegisterAppRsp:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindRouter)
		data.Base.SubId = uint32(CMDID_Router_IDRegisterAppRsp)
		return true, data.Base
	default:
		return false, nil
	}
}

func setClientBase(input interface{}) (bool, *BaseInfo) {
	switch data := input.(type) {
	case *LoginReq:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDLoginReq)
		return true, data.Base
	case *LoginRsp:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDLoginRsp)
		if data.UserBaseInfo == nil { //顺便设置一下复合数据类型
			data.UserBaseInfo = new(BaseUserInfo)
		}
		if data.UserSesionInfo == nil {
			data.UserSesionInfo = new(UserSessionInfo)
		}
		return true, data.Base
	case *LogoutReq:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDLogoutReq)
		return true, data.Base
	case *LogoutRsp:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDLogoutRsp)
		return true, data.Base
	case *QueryFundReq:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDQueryFundReq)
		return true, data.Base
	case *QueryFundRsp:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDQueryFundRsp)
		return true, data.Base
	case *GetOnlineUserReq:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDGetOnlineUserReq)
		return true, data.Base
	case *GetOnlineUserRsp:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDGetOnlineUserRsp)
		return true, data.Base
	case *KickUserReq:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDKickUserReq)
		return true, data.Base
	case *KickUserRsp:
		if data.Base == nil {
			data.Base = new(BaseInfo)
		}
		data.Base.KindId = uint32(CMDKindId_IDKindClient)
		data.Base.SubId = uint32(CMDID_Client_IDKickUserRsp)
		return true, data.Base
	default:
		return false, nil
	}
}

// 设置input的baseinfo值，如果返回false说明这个类型找不到
func SetBaseKindAndSubId(input interface{}) (bool, *BaseInfo) {
	if input == nil {
		return false, nil
	}
	switch reflect.TypeOf(input).String() {
	//以下报文属于tcp.proto，CMDKindId_IDKindNetTCP大类
	case "* TCPTransferMsg":
		fallthrough
	case "* TCPSessionCome":
		fallthrough
	case "* TCPSessionClose":
		fallthrough
	case "* TCPSessionKick":
		return setTcpBase(input)
	//以下报文属于gate.proto,CMDKindId_IDKindGate大类
	case "* PulseReq":
		fallthrough
	case "* PulseRsp":
		fallthrough
	case "* GateTransferData":
		return setGateBase(input)
	//以下报文属于router.proto,CMDKindId_IDKindRouter大类
	case "* RouterTransferData":
		fallthrough
	case "* RegisterAppReq":
		fallthrough
	case "* RegisterAppRsp":
		return setRouterBase(input)
	//以下报文属于client.proto,CMDKindId_IDKindClient大类
	case "* LoginReq":
		fallthrough
	case "* LoginRsp":
		fallthrough
	case "* LogoutReq":
		fallthrough
	case "* LogoutRsp":
		fallthrough
	case "* QueryFundReq":
		fallthrough
	case "* QueryFundRsp":
		fallthrough
	case "* GetOnlineUserReq":
		fallthrough
	case "* GetOnlineUserRsp":
		fallthrough
	case "* KickUserReq":
		fallthrough
	case "* KickUserRsp":
		return setClientBase(input)
	default:
		fmt.Println("input为不识别的类型")
		return false, nil
	}
	return false, nil
}

// 复制除了kindid和subid以外的值
func CopyBaseExceptKindAndSubId(dst *BaseInfo, src *BaseInfo) {
	if dst == nil || src == nil {
		fmt.Println("CopyBaseExceptKindAndSubId传入的参数是空,dst=", dst, ",src=", src)
		return
	}
	dst.ConnId = src.ConnId
	dst.GateConnId = src.GateConnId
	dst.RemoteAdd = src.RemoteAdd
	dst.AttApptype = src.AttApptype
	dst.AttAppid = src.AttAppid
}

func SetCommonMsgBaseByRouterTransferData(dst *BaseInfo, srcRouter *RouterTransferData) {
	if dst == nil || srcRouter == nil {
		fmt.Println("SetCommonMsgBaseByRouterTransferData传入的参数是空,dst=", dst, ",src=", srcRouter)
		return
	}
	CopyBaseExceptKindAndSubId(dst, srcRouter.Base)
	//和客户端相关的都要重新赋值，取的是srcRouter的值，而不是srcRouter.Base的值
	dst.AttAppid = srcRouter.SrcAppid
	dst.AttApptype = srcRouter.SrcApptype
	dst.RemoteAdd = srcRouter.ClientRemoteAddress
	dst.GateConnId = srcRouter.AttGateconnid
}
