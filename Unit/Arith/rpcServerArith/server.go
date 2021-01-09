package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

// 聲明算術運算結構體
type Arith struct {
}

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

// 乘法運算
func (a *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Product = req.A * req.B
	return nil
}

// 商和餘數
func (a *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("餘數不能為0")
	}

	// 商
	res.Quotient = req.A / req.B

	// 餘數
	res.Remainder = req.A % req.B

	return nil
}

func main() {
	// 1.註冊服務
	arith := new(Arith)
	rpc.Register(arith)

	// 2.把服務處理綁定到http協議上
	rpc.HandleHTTP()

	// 3.監聽服務，等待客戶端調用求面積和周長的方法
	fmt.Println("服務啟動中，監聽端口8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
