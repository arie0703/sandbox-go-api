package calendar

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// FoodExpense 食費の情報を格納する構造体
type FoodExpense struct {
	Emoji     string
	PlaceName string
	Amount    int
	Date      string
}

// CalculateFoodExpenses 指定された年月の食費を計算する
func CalculateFoodExpenses(events []*calendar.Event) ([]FoodExpense, int, error) {
	var foodExpenses []FoodExpense
	totalAmount := 0

	// 食費の予定を識別する正規表現
	// 🍚 場所名 1000 の形式を想定
	foodExpensePattern := regexp.MustCompile(`^([^\s]+)\s+([^\s]+(?:\s+[^\s]+)*)\s+(\d+)$`)

	for _, event := range events {
		if event.Summary == "" {
			continue
		}

		matches := foodExpensePattern.FindStringSubmatch(event.Summary)
		if len(matches) == 4 {
			emoji := matches[1]
			placeName := matches[2]
			amountStr := matches[3]

			amount, err := strconv.Atoi(amountStr)
			if err != nil {
				continue // 金額の変換に失敗した場合はスキップ
			}

			// 日付を取得
			var dateStr string
			if event.Start.DateTime != "" {
				// 時刻付きの場合は日付部分のみを取得
				dateStr = event.Start.DateTime[:10] // YYYY-MM-DD形式
			} else {
				// 日付のみの場合はそのまま使用
				dateStr = event.Start.Date
			}

			foodExpense := FoodExpense{
				Emoji:     emoji,
				PlaceName: placeName,
				Amount:    amount,
				Date:      dateStr,
			}

			foodExpenses = append(foodExpenses, foodExpense)
			totalAmount += amount
		}
	}

	return foodExpenses, totalAmount, nil
}

// PrintFoodExpenses 外食費用の詳細を表示する
func PrintFoodExpenses(foodExpenses []FoodExpense, totalAmount int) {
	if len(foodExpenses) == 0 {
		fmt.Println("外食した記録はありません。")
		return
	}

	fmt.Println("\n=== 外食費用の詳細 ===")
	for _, expense := range foodExpenses {
		fmt.Printf("%s %s %s: %d円\n", expense.Date, expense.Emoji, expense.PlaceName, expense.Amount)
	}
	fmt.Printf("\n合計: %d円\n", totalAmount)
}

// RunFoodExpenses 食費計算を実行する
func RunFoodExpenses(year, month int) {
	ctx := context.Background()

	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("credentials.json の読み込みに失敗しました: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("OAuth2 設定の作成に失敗しました: %v", err)
	}

	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("カレンダーサービスの作成に失敗しました: %v", err)
	}

	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	events, err := srv.Events.List("primary").
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(start.Format(time.RFC3339)).
		TimeMax(end.Format(time.RFC3339)).
		OrderBy("startTime").
		Do()
	if err != nil {
		log.Fatalf("予定の取得に失敗しました: %v", err)
	}

	foodExpenses, totalAmount, err := CalculateFoodExpenses(events.Items)
	if err != nil {
		log.Fatalf("食費の計算に失敗しました: %v", err)
	}

	fmt.Printf("%d年%d月の外食費用計算:\n", year, month)
	PrintFoodExpenses(foodExpenses, totalAmount)
}
