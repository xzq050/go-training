package main // 可執行程式必須使用 main 封包
import "fmt" // 載入内建的 fmt 封包，用來做基本的輸出輸入
func main(){ //建立 main 函式，程式的進入點
	// 基本輸出練習：fmt.Println(資料,資料,...)
	fmt.Println(3,"Hello",true)

	// 基本輸入練習：fmt.Scanln(&變數名稱,&變數名稱,...)
	// &變數名稱：取得變數的指標（Pointer）
/* 	var x int
	fmt.Print("輸入一個數字：")
	fmt.Scanln(&x)
	fmt.Println(x) */

	// 整合練習：輸入兩個數字，并輸出數字相乘的結果
	/* var x int
	var y int
	fmt.Print("輸入第一個數字：")
	fmt.Scanln(&x)
	fmt.Print("輸入第二個數字：")
	fmt.Scanln(&y)
	var result int = x+y
	fmt.Println("result：",result)
 */
var x int
var y int
fmt.Println("輸入兩個數字，用空格隔開：")
fmt.Scanln(&x,&y)
var result=x+y
fmt.Println("result：",result)

}




// https://www.youtube.com/watch?v=iH_fLuaGfaI&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=3

/* 
1. Golang 載入封包
1.1 import 封包名稱
1.2 import "fmt"

2. 使用 fmt 封包做基本輸入
2.1 宣告變數
2.2 fmt.Scanln(&變數名稱)
2.3 取得變數對應的指標

3. 使用 fmt 封包做基本輸出
3.1 輸出換行 fmt.Println(資料)
3.2 輸出不換行 fmt.Print(資料)
 */