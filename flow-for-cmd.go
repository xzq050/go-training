package main
import "fmt"
func main(){
	// break
/* 	var x int=0
	for x<5{
		if x==3{
			break
		}
		fmt.Println(x)
		x++
	} */


	// continue

/* 	
	var x int
	for x=0;x<5;x++{
		if x%2==0{ // x 是偶數，不印出來
			continue
		}
		fmt.Println(x)
	}

 */


	// 實際範例，持續讓使用者輸入數字，計算總和。直到使用者輸入 0 爲止
	var result int=0
	var n int
	for true {  // 無窮迴圈
		// var n int
		fmt.Scanln(&n)
		if n==0 {
			break
		}
		result+=n // ressult=result+n
	}
	fmt.Println("總和：",result)


/*  var n int
 var flag int =1
 var sum int = 0
for flag != 0 {
	fmt.Scanln(&n)
	sum += n
	flag = n
}
fmt.Println("總和：",sum) */


}











// https://www.youtube.com/watch?v=MDXGIKkDg1c&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=7



/* 1. 迴圈搭配的命令
1.1 break 強迫中斷迴圈
1.2 continue 強迫回到迴圈的開頭，執行下一次的迴圈

2. Go 語言練習範例
2.1 基本命令的練習展示
2.2 搭配 break 的實際範例：累加使用者輸入的數字，輸入 0 結束程式*/
