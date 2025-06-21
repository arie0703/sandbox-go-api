package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arie0703/sandbox-go-api/google/cmd/calendar"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用法: go run . <command>")
		fmt.Println("例: go run . calendar 202506")
		fmt.Println("例: go run . food-expenses 202506")
		return
	}

	switch os.Args[1] {
	case "calendar":
		if len(os.Args) < 3 {
			fmt.Println("年月を指定してください（例: 202506）")
			return
		}
		yyyymm := os.Args[2]
		if len(yyyymm) != 6 {
			fmt.Println("YYYYMM形式で入力してください（例: 202506）")
			return
		}

		year, err1 := strconv.Atoi(yyyymm[:4])
		month, err2 := strconv.Atoi(yyyymm[4:])
		if err1 != nil || err2 != nil || month < 1 || month > 12 {
			fmt.Println("無効な年月です")
			return
		}

		calendar.Run(year, month)
	case "food-expenses":
		if len(os.Args) < 3 {
			fmt.Println("年月を指定してください（例: 202506）")
			return
		}
		yyyymm := os.Args[2]
		if len(yyyymm) != 6 {
			fmt.Println("YYYYMM形式で入力してください（例: 202506）")
			return
		}

		year, err1 := strconv.Atoi(yyyymm[:4])
		month, err2 := strconv.Atoi(yyyymm[4:])
		if err1 != nil || err2 != nil || month < 1 || month > 12 {
			fmt.Println("無効な年月です")
			return
		}

		calendar.RunFoodExpenses(year, month)
	default:
		fmt.Println("未知のコマンド:", os.Args[1])
		fmt.Println("利用可能なコマンド: calendar, food-expenses")
	}
}
