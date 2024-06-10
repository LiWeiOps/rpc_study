package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {
	Num1 int
	Num2 int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	res1 := 0
	err = client.Call("Col.ChengFa", Params{12, 5}, &res1)
	if err != nil {
		log.Fatal("ChengFa  ", err)
		return
	}
	fmt.Println("res1:", res1)
	res2 := make([]int, 2)
	err = client.Call("Col.ChuFa", Params{12, 5}, &res2)
	if err != nil {
		log.Fatal("ChuFa  ", err)
		return
	}
	fmt.Println("res2:", res2)
}
