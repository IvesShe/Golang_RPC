package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

// 服務端，求矩形面績和週長

// 聲明矩形對象
type Rect struct {
}

// 聲明參數結構體，字段首字母大寫
type Params struct {
	// 長和寬
	Width  int
	Height int
}

// 定義求矩形面積的方法
func (r *Rect) Area(p Params, res *int) error {
	*res = p.Width * p.Height
	return nil
}

// 定義求矩形周長的方法
func (r *Rect) Perimeter(p Params, res *int) error {
	*res = (p.Width + p.Height) * 2
	return nil
}

func main() {
	// 1.註冊服務
	rect := new(Rect)
	rpc.Register(rect)

	// 2.把服務處理綁定到http協議上
	rpc.HandleHTTP()

	// 3.監聽服務，等待客戶端調用求面積和周長的方法
	fmt.Println("服務啟動中，監聽端口8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
