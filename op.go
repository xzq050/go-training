package main
import "fmt"
func main(){
	// 算數運算：+，-，*，/
	var x int
	x = 3*3+10
	fmt.Println(x)
	// 指定運算：=，+=，-=，*=，/=
	x=5
	// x=x+2
	// x+=2 // x=x+2
	x-=2 //x=x-2
	fmt.Println(x)
	// 單元運算：++，--
	x++
	fmt.Println(x)

	//比較運算：>,<,>=,<=,==,!=
	var result bool=4==4
	fmt.Println(result)

	//邏輯運算：!,&&,||
	// result = !true
	// result = true&&false // and,且，&&：兩邊都是 true，結果就是true
	result=false||false  // or,或，||：只要有一邊是 true，結果就是 true
	fmt.Println(result)
}


// https://www.youtube.com/watch?v=cMq1aVvHZKo&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=4

/* 
0. 基本運算符號 ( 運算子 )

1. 運算符號的種類
1-1. 算數運算 +, -, *, /
1-2. 指定運算 =, +=, -=, *=, /=
1-3. 單元運算 ++, --
1-4. 比較運算
1-5. 邏輯運算 !, &&, ||

 */