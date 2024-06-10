package rpc

import (
	"fmt"
	"net"
	"reflect"
)

type Server struct {
	// 地址
	Addr string
	// 服务端维护的函数名到函数反射值的map
	funcs map[string]reflect.Value
}

// 创建服务端对象
func NewServer(addr string) *Server {
	return &Server{Addr: addr, funcs: make(map[string]reflect.Value)}
}

// 服务端绑定注册方法
// 将函数名与函数真正真正实现对应
// 第一个参数为函数名， 第二个参数传入真正的函数
func (s *Server) Register(rpcName string, f interface{}) {
	if _, ok := s.funcs[rpcName]; ok {
		return
	}
	// map中没有值，将映射添加进map， 方便调用
	s.funcs[rpcName] = reflect.ValueOf(f)
}

// 服务端等待调用
func (s *Server) Run() {
	// 监听
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		fmt.Println("listen err ", err)
		return
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("conn err ", err)
			return
		}
		// 创建会话
		session := NewSession(conn)
		// rpc读取数据
		b, err := session.Read()
		if err != nil {
			return
		}
		// 对数据解码
		var rpcData RPCData
		rpcData, err = decode(b)
		if err != nil {
			fmt.Println("decode err ", err)
			return
		}
		// 根据读取到的数据的函数name， 调用函数
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			fmt.Printf("函数 %s 不存在 err : %v", rpcData.Name, err)
			return
		}
		// 解析遍历客户端传入的参数，放到一个切片中
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			fmt.Println(arg, 11111)
			fmt.Println(reflect.ValueOf(arg), 333333)
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		// 反射调用函数，传入参数
		fmt.Println(rpcData.Args, 2222222)
		fmt.Println(len(inArgs), 11111111)
		out := f.Call(inArgs)
		// 解析遍历结果，放到一个数组中
		outArgs := make([]interface{}, 0, len(out))
		fmt.Println("len out:", len(out))
		for _, o := range out {
			fmt.Println("out: ", o)
			outArgs = append(outArgs, o.Interface())

		}
		fmt.Println("outArgs: ", outArgs)
		// 包装响应数据，返回给客户端
		respRPCData := RPCData{Name: rpcData.Name, Args: outArgs}
		// 编码
		RespBytes, err := encode(respRPCData)
		if err != nil {
			fmt.Println("encode err ", err)
			return
		}
		// 返回数据
		err = session.Write(RespBytes)
		if err != nil {
			fmt.Println("session write err ", err)
			return
		}
	}
}
