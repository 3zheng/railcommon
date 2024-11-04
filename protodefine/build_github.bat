
@protoc.exe --proto_path=. --go_out=.  *.proto

pause

@rem protoc.exe --proto_path=. --go_out=. a2.proto
@rem protoc.exe --proto_path=. --go_out=. a1.proto

@rem d:\vs2013\run\win32_release\protoc.exe --lua_out=.. logger/logger.proto
@rem d:\vs2013\run\win32_release\protoc.exe --cpp_out=.. gateclient/gateclient.proto
@rem d:\vs2013\run\win32_release\protoc.exe --go_out=.. gate/gate.proto