package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 聲明接收的參數結構體
type ArithRequest struct {
	A int
	B int
}

// 聲明返回客戶端參數結構體
type ArithResponse struct {
	// 乘積
	Product int

	// 商
	Quotient int

	// 餘數
	Remainder int
}

// 調用服務
func main() {
	// 1.連接遠程RPC服務
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("連接成功...")

	// 2.調用遠程方法
	// 定義接收服務端傳回來的計算結果的變量
	var res ArithResponse

	// 求乘積
	req := ArithRequest{9, 2}
	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("乘法運算: %d * %d = %d  \n", req.A, req.B, res.Product)

	// 求商和餘數
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d / %d , 商 = %d , 餘數 = %d\n  \n", req.A, req.B, res.Quotient, res.Remainder)
}
