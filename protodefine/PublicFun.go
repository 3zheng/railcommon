package protodefine

import (
	"fmt"
	"runtime"
)

func GetAppTypeName(appType uint32) string {
	var name string
	switch appType {
	case uint32(EnumAppType_Gate):
		name = "Gate"
	case uint32(EnumAppType_Router):
		name = "Router"
	case uint32(EnumAppType_Login):
		name = "Login"
	case uint32(EnumAppType_Online):
		name = "Online"
	case uint32(EnumAppType_Fund):
		name = "Fund"
	case uint32(EnumAppType_List):
		name = "List"
	case uint32(EnumAppType_FreeMatch):
		name = "FreeMatch"
	case uint32(EnumAppType_Match):
		name = "Match"
	case uint32(EnumAppType_TableLogic):
		name = "TableLogic"
	case uint32(EnumAppType_MatchPhase):
		name = "MatchPhase"
	case uint32(EnumAppType_RankList):
		name = "RankList"
	case uint32(EnumAppType_MatchDB):
		name = "MatchDB"
	}
	return name
}

func OutputMyLog(a ...interface{}) {
	funcName, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Println("Func Name=", runtime.FuncForPC(funcName).Name())
		fmt.Printf("file: %s    line=%d\n", file, line)
		fmt.Println(a...)
	}
}
