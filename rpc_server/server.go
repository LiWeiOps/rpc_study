package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Col struct {
}

type Params struct {
	Num1 int
	Num2 int
}

func (c *Col) ChengFa(p Params, res *int) error {
	*res = p.Num1 * p.Num2
	return nil
}

func (c *Col) ChuFa(p Params, res *[]int) error {
	*res = make([]int, 2)
	(*res)[0] = p.Num1 / p.Num2
	(*res)[1] = p.Num1 % p.Num2
	return nil
}

//func main() {
//	// 声明新对象
//	col := new(Col)
//	// 向rpc进行注册
//	err := rpc.Register(col)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// rpc通过HTTP进行调用
//	rpc.HandleHTTP()
//	err = http.ListenAndServe(":8080", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	// 声明新对象
	col := new(Col)
	// 向rpc进行注册
	err := rpc.Register(col)
	if err != nil {
		log.Fatal(err)
	}
	// 通过jsonrpc可以进行跨语言调用
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal("conn  ", err)
			continue
		}

		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
