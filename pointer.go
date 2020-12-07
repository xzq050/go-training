package main
import "fmt"
func main(){
	// 1.建立存放資料的變數
	var x int=5
	fmt.Println("原來的資料",x)
	//2.取得記憶體位置:&變數名稱
	fmt.Println("資料的記憶體位置",&x)
	//3.存放到指標變數。注意指標資料形態: *資料形態
	var xPtr *int=&x
	fmt.Println("資料的記憶體位置",xPtr)
	//4.反解指標變數:*指標變數名稱
	fmt.Println("反解指標回原來的資料",*xPtr)

	var word string="Hello"
	fmt.Println(word)
	var wordPtr *string=&word
	fmt.Println(wordPtr)
	var word2 string=*wordPtr
	fmt.Println(word2)
	fmt.Println(*wordPtr)
}

//Golang 指標基礎 Pointer - 記憶體位址、指標變數與資料型態、反解指標 By 彭彭
//https://www.youtube.com/watch?v=SC8MPfhh9_8&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=10
/*
喜歡彭彭的教學影片嗎？點擊「加入」按紐取得更多會員服務哦。
加入會員：https://www.youtube.com/channel/UCguZ...

1. 指標的基本操作流程
1.1 建立資料變數
1.2 取得資料變數的記憶體位址
1.3 建立指標變數
1.4 反解指標變數

2. 資料型態的差異
2.1 原始資料型態：int、float64、string
2.2 指標資料型態：*int、*float64、*string

3. 學習方法
3.1 了解指標運作觀念
3.2 反覆練習直到熟練為止

更多學習資訊，請到彭彭的課程網站：https://training.pada-x.com/
*/