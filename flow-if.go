package main
import "fmt"
func main(){
	// 基本語法練習
	/* 
	if false{
		fmt.Println("Go")
	}else{
		fmt.Println("Not Go")
	}
 */
	// 簡易情境：ATM 檢測使用者輸入的金額是否適當
	var money int
	fmt.Print("請問想領多少錢？")
	fmt.Scanln(&money)
	fmt.Println("money:",money)
	if money<100{
		fmt.Println("Too Few")
	}else if money<=100000{
		fmt.Println("OK")
	}else{
		fmt.Println("Too Much")
	}
	fmt.Println("執行完畢")
}




// https://www.youtube.com/watch?v=poQYWnS4-i0&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=5




/* 1. Golang 流程控制
1.1 if 判斷式
1.2 if ... else 判斷結構
1.3 if ... else if ... else 多條件判斷結構

2. Golang 練習範例
2.1 判斷式基本語法練習
2.2 ATM 情境範例練習 */