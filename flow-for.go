package main
import "fmt"
func main(){
	//基本迴圈使用
/* 
	var x int = 0
	for x<3 {
		// fmt.Println("Hello")
		// fmt.Println("Go Lang")
		fmt.Println(x)
		x++
	}
	 */
/* 
	 var x int
	 for x=0;x<5;x+=2 {
		 fmt.Println(x)
	 }
	  */
	// 實際範例，計算 1+2+3+...+50的結果 
	var result int=0
	var x int=1
	for x<=50 {  //測試 1+2+3
		fmt.Println("result:",result,"x:",x)
		result=result+x
		x++
	}
	fmt.Println("last-result:",result)
}




// https://www.youtube.com/watch?v=u3DPQ7tdynw&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=6



/* 1. Golang 流程控制
1.1 for 迴圈的觀念
1.2 for 迴圈基礎語法

2. Golang 練習範例
2.1 無窮迴圈的展示
2.2 for 迴圈基礎語法練習
2.3 for 迴圈範例：等差級數加法 1+2+3+...+50 */