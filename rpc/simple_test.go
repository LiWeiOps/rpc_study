package rpc

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"testing"
)

// 用户查询

// 用户结构体
// 字段首字母必须大写
type User struct {
	Name string
	Age  int
}

// 用于测试的用户的查询方法
func QueryUser(uid int) (User, error) {
	user := make(map[int]User)
	user[0] = User{Name: "zs", Age: 20}
	user[1] = User{Name: "ls", Age: 21}
	user[2] = User{Name: "ww", Age: 22}

	if u, ok := user[uid]; ok {
		return u, nil
	}
	return User{}, errors.New("query err")
}

func TestRPC(t *testing.T) {
	// 需要对interface{}可能产生的类型进行注册
	gob.Register(User{})
	addr := "127.0.0.1:8080"
	// 创建服务端
	srv := NewServer(addr)
	// 将方法注册到服务端
	srv.Register("queryUser", QueryUser)
	// 服务端等待调用
	go srv.Run()
	// 客户端获取连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error(err)
	}
	// 创建客户端
	cli := NewClient(conn)
	// 创建函数原型
	var query func(int) (User, error)
	cli.callRPC("queryUser", &query)
	user1, err1 := query(1)
	if err1 != nil {
		t.Fatal(err)
	}
	fmt.Println(user1)
}
