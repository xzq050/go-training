package main
//import ("fmt")
import "fmt"
func add(xPtr *int){
	*xPtr=*xPtr+10
	fmt.Println("Add Function",*xPtr)

}

func main(){
	var a int=10
	//var aPtr *int=&a
	//add(aPtr) //Pass by pointer
	add(&a)
	fmt.Println("Main Function",a)
	//和使用者要求輸入，運用到指標參數 Pass by point
	var msg string
	var msgPtr *string=&msg
	fmt.Scanln(msgPtr)
	//fmt.Scanln(&msg)//傳遞字串變數的指標（記憶體位置）
	fmt.Println(msg)

}

/*
func add(x int){
	x=x+10
	fmt.Println("Add Function",x)
}
func main(){
	var a int=10
	add(a) // Pass by value
	fmt.Println("Main Function",a)
}
*/


//Golang 指標參數 - Pass by Pointer 和 Pass by Value 函式參數傳遞 By 彭彭
//https://www.youtube.com/watch?v=k_E9FCehyz4&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=11
/*
喜歡彭彭的教學影片嗎？點擊「加入」按紐取得更多會員服務哦。
加入會員：https://www.youtube.com/channel/UCguZ...

1. 指標基本觀念
1.1 建立指標變數
1.2 反解指標變數

2. 函式的參數傳遞
2.1 傳值 Pass by Value
2.2 傳址 Pass by Pointer

3. 實務的運用
3.1 函式參數與原資料的連動性
3.2 fmt.Scanln 的指標參數運用

更多學習資訊，請到彭彭的課程網站：https://training.pada-x.com/
*/