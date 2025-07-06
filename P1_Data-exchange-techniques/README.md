在本文中，假设分别采用Python和Golang构建了客户端与服务端，需要实现两个端口之间的数据通信。上述场景在大规模推荐系统中有广泛的应用场景，因为传统的基于协同过滤的推荐方法在基于Go语言构建的后端中部署，但近些年来机器学习、深度学习及大模型赋能的推荐方法受到了越来越多的关注，而Pytorch、Tensorflow、Scikit-learn框架大多基于Python构建，因此将这些模型部署在Golang中，需要实现数据通信。此处给出了常见的数据通信方案，为了简便起见，基于Python的客户端仅构建的一个简单的功能，即接收服务端传递的两个数字，对数字进行求和，之后将数据返回服务端。常见的方法有将客户端包装成Web应用，暴露端口，之后在Golang端对客户端的端口进行访问。

------

- 利用Flask将Python的函数或者模型包装为Web API

**核心是利用Json进行数据交换**，定义了结果result与错误error

客户端运行

```cmd
python Flask_simple.py
```

服务端运行

```cmd
go run main_flask.go
```

- 利用FastAPI将Python函数或者模型包装为Web API

基本代码逻辑与基于Flask的方法类似，需要预先定义输入输出的类模型。

```cmd
python -m unicorn Fastapi:app -- host 0.0.0.0 --port 5000
```

------

- 利用Protobuf与gRPC实现格式统一的多端信息交互框架

利用给出的protoc软件包，将其bin目录加入环境变量中的Path路径。

gRPC接口相当于给出了信息交换过程中的数据输入输出格式。

```protobuf
syntax = 'proto3';
package server;

service Calculator {
   rpc Add(AddRequest) returns (AddResponse) {}
}

message AddRequest {
   int32 a = 1;
   int32 b = 2;
}

message AddResponse {
   int32 result = 1;
}
```

Python 端进行gRPC客户端的程序生成

```powershell
python -m pip install grpcio grpcio-tools
python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. service.proto
```

在工作目录上会生成service_pb2.py和service_pb2_grpc.py程序

Go端进行gRPC服务端的程序生成

安装Go的相关插件

```powershell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go mod init pb1
go get google.golang.org/grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. service.proto
```

注意利用Go语言生成service_grpc.pb.go和service.pb.go时，需要显式指定文件的输出位置，即在Python的Protobuf文件中添加一个option选项，其中pb1与创建的go.mod中的模块名称对应。

```protobuf
option go_package = "pb1/proto";
```

