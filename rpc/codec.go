package rpc

import (
	"bytes"
	"encoding/gob"
)

// 定义数据格式和编解码

// 定义rpc交互的数据格式
type RPCData struct {
	// 访问的函数
	Name string
	// 传入的参数
	Args []interface{}
}

// 编码
func encode(data RPCData) ([]byte, error) {
	// rpc默认使用gob进行编解码
	var buf bytes.Buffer
	// 得到字节数组的编码器
	bufEnc := gob.NewEncoder(&buf)
	// 对数据编码
	err := bufEnc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 解码
func decode(b []byte) (RPCData, error) {
	buf := bytes.NewBuffer(b)
	// 返回字节数组解码器
	bufDec := gob.NewDecoder(buf)
	var data RPCData
	// 对数据解码
	err := bufDec.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
