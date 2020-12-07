package main
import "fmt"

func show(msg string){
	if msg=="Hello"{
		return
	}
	
	fmt.Println(msg)
}

func add(n1 int,n2 int) int {
	var result int=n1+n2
	//fmt.Println(result)
	return result
}

func test()(int,string){
	return 5,"Test"
}
func main(){
	// 基本的return運作方式
	/*
	show("Hello")
	show("你好")
	show("測試")
	*/
	// 單一回傳值
	/*
	var x int = add(3,4)
	fmt.Println(x*20)
	*/

	//多個回傳值
	var x int
	var s string
	x,s=test()
	fmt.Println(x,s)
	test()
}


//https://www.youtube.com/watch?v=DsH867OY2YA&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=9
/*
喜歡彭彭的教學影片嗎？點擊「加入」按紐取得更多會員服務哦。
加入會員：https://www.youtube.com/channel/UCguZ...

1. 結束函式
1.1 自然結束
1.2 使用 return 強制結束

2. 函式的回傳值
2.1 定義回傳值的資料型態
2.2 使用 return 帶入回傳值
2.3 在函式的外部 ( 呼叫函式的地方 ) 取得回傳值

3. 多個回傳值的定義與使用

更多學習資訊，請到彭彭的課程網站：https://training.pada-x.com/

*/