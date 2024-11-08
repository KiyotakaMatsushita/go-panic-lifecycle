package main

import (
	"fmt"
)

func main() {
	fmt.Println("プログラム開始")
	outerFunction()
	fmt.Println("プログラム終了")
}

func outerFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("outerFunctionでリカバリー:", r)
		}
	}()
	innerFunction()
	fmt.Println("outerFunctionの残りのコード")
}

func innerFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("innerFunctionでリカバリー:", r)
			fmt.Println("innerFunctionのdeferが実行されました")
			panic(r) // パニックを上位に伝播
		}
		fmt.Println("innerFunctionのdeferが実行されました")
	}()
	deeperFunction()
	fmt.Println("innerFunctionの残りのコード")
}

func deeperFunction() {
	defer func() {
		fmt.Println("deeperFunctionのdeferが実行されました")
	}()
	deepestFunction()
	fmt.Println("deeperFunctionの残りのコード")
	panic("deeperFunctionでパニックが発生")
}

func deepestFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("deepestFunctionでリカバリー:", r)
			fmt.Println("deepestFunctionのdeferが実行されました")
			// パニックを上位に伝播させない
			return
		}
		fmt.Println("deepestFunctionのこのdeferが実行されない")
	}()
	panic("deepestFunctionでパニックが発生")
	fmt.Println("deepestFunctionの残りのコードは実行されません")
}
