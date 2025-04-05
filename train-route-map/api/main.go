package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Neo4jドライバー
var driver neo4j.Driver

// Neo4jに接続する関数
func connectToNeo4j() {
	// .envファイルの読み込み
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から認証情報を取得
	neo4jAuth := os.Getenv("NEO4J_AUTH")
	if neo4jAuth == "" {
		log.Fatal("NEO4J_AUTH environment variable is not set")
	}

	// Neo4jに接続
	driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""))
	if err != nil {
		log.Fatal("Failed to create the driver:", err)
	}
}

// 最短経路を取得するAPI
func getShortestPath(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	startStation := params["start"]
	endStation := params["end"]

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	query := `
		MATCH (start:Station {name: $startName}), (end:Station {name: $endName}),
			  p = shortestPath((start)-[:CONNECTED_TO*]-(end))
		RETURN p
	`

	result, err := session.Run(query, map[string]interface{}{
		"startName": startStation,
		"endName":   endStation,
	})
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}

	if result.Next() {
		path := result.Record().GetByIndex(0).(neo4j.Path)
		json.NewEncoder(w).Encode(path)
	} else {
		http.Error(w, "No path found", http.StatusNotFound)
	}
}

// 駅の情報を取得するAPI
func getStationInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stationName := params["name"]

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	query := `
		MATCH (station:Station {name: $stationName})
		RETURN station
	`

	result, err := session.Run(query, map[string]interface{}{
		"stationName": stationName,
	})
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}

	if result.Next() {
		stationNode := result.Record().GetByIndex(0).(neo4j.Node)
		stationData := map[string]interface{}{
			"name": stationNode.Props["name"],
		}
		json.NewEncoder(w).Encode(stationData)
	} else {
		http.Error(w, "Station not found", http.StatusNotFound)
	}
}

func main() {
	// Neo4jに接続
	connectToNeo4j()
	defer driver.Close()

	// Go HTTPサーバーの設定
	r := mux.NewRouter()

	// 最短経路API
	r.HandleFunc("/shortestpath/{start}/{end}", getShortestPath).Methods("GET")

	// 駅情報API
	r.HandleFunc("/station/{name}", getStationInfo).Methods("GET")

	// サーバー開始
	http.Handle("/", r)
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
