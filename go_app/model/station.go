package model

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func GetStations(neLat, neLng, swLat, swLng float64) ([]map[string]interface{}, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, fmt.Errorf("failed to load location: %v", err)
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")

	c := mysql.Config{
		DBName:    dbName,
		User:      dbUser,
		Passwd:    dbPass,
		Addr:      dbAddr,
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	// データベース接続
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// クエリの実行
	query := "SELECT * FROM stations WHERE lat <= ? AND lat >= ? AND lon <= ? AND lon >= ?"
	fmt.Println("Executing query:", query, neLat, swLat, neLng, swLng) // クエリを出力

	rows, err := db.Query(query, neLat, swLat, neLng, swLng)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	// カラム名を取得
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %v", err)
	}

	// データを格納するスライス
	var results []map[string]interface{}

	for rows.Next() {
		// カラム数に応じて値を格納するスライスを作成
		values := make([]sql.RawBytes, len(columns))
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		// 1行分のデータを取得
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		// 1行分のデータをマップに変換
		row := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				row[columns[i]] = nil
			} else {
				row[columns[i]] = string(col)
			}
		}
		results = append(results, row)
	}

	// エラーチェック
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return results, nil
}
