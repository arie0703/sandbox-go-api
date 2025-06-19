package main

import (
	"fmt"
	"os"

	"github.com/arie0703/sandbox-go-api/google/cmd/calendar"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用法: go run . <command>")
		fmt.Println("例: go run . calendar")
		return
	}

	switch os.Args[1] {
	case "calendar":
		calendar.Run()
	default:
		fmt.Println("未知のコマンド:", os.Args[1])
	}
}
