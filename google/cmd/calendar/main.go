package calendar

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func getClient(config *oauth2.Config) *http.Client {
	tokFile := tokenCacheFile()
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func tokenCacheFile() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get current user: %v", err)
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir, "calendar-go-quickstart.json")
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("以下のURLをブラウザで開いて、認証コードを入力してください：\n%v\n", authURL)

	var code string
	fmt.Print("認証コード: ")
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("認証コードの読み取りに失敗しました: %v", err)
	}

	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("トークンの取得に失敗しました: %v", err)
	}
	return tok
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("トークンを保存しました: %s\n", path)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("トークンファイルの作成に失敗しました: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Run(year, month int) {
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

	fmt.Printf("%d年%d月の予定:\n", year, month)
	if len(events.Items) == 0 {
		fmt.Println("予定はありません。")
	} else {
		for _, item := range events.Items {
			var start string
			if item.Start.DateTime != "" {
				start = item.Start.DateTime
			} else {
				start = item.Start.Date
			}
			fmt.Printf("%s - %s\n", start, item.Summary)
		}
	}
}
