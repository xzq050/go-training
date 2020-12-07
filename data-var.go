package main // 可執行程式必須使用 main 封包
import "fmt" // 載入内建的 fmt 封包，用來做基本的輸出輸入
func main(){ //建立 main 函式，程式的進入點
	// 執行程式時，從這邊開始
	// 輸出訊息到終端：fmt.Println(資料,資料,...)
	/*
	fmt.Println(3) //整數 int
	fmt.Println(3.1415) // 浮點數 float64
	fmt.Println("測試中文") // 字串 string
	fmt.Println(true) // 布林值
	fmt.Println('a') // 字符 rune
	*/
	// 變數的使用
	var x int // 宣告變數 x
	x=4 //指定資料：把資料 4 放進變數裏
	fmt.Println(x) // 使用變數：用變數的名稱代替資料，做操作
	x=10 // 指定新的資料：把資料 10 放進變數，覆蓋原來的資料
	fmt.Println(x)
	// x="hello"  資料形態不正確
	// var f float64
	// f = 3.1415
	var f float64 = 3.1415926
	fmt.Println(f)
	var s string ="哈嘍 Golang"
	fmt.Println(s)
	var test bool=true
	fmt.Println(test)
	var c rune='b'
	fmt.Println(c)

}





// https://www.youtube.com/watch?v=D6a9RpuL_UA&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=2


/* 1. Golang 資料型態
1.1 整數 int
1.2 浮點數 float64
1.3 字串 string
1.4 布林值 bool
1.5 字符 rune

2. Golang 變數
2.1 宣告變數與變數的資料型態
2.2 指定變數中的資料
2.3 使用變數代替資料做運算 */