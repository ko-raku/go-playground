package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQLドライバ
	"io"
	"net/http"
)

type ResponseData struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Visibility int     `json:"visibility"`
		Pop        float64 `json:"pop"`
		Sys        struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
		Rain  struct {
			H float64 `json:"3h"`
		} `json:"rain,omitempty"`
	} `json:"list"`
	City struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Timezone   int    `json:"timezone"`
		Sunrise    int    `json:"sunrise"`
		Sunset     int    `json:"sunset"`
	} `json:"city"`
}

func main() {
	apiKey := "{API key}"
	apiURL := "https://api.openweathermap.org/data/2.5/forecast/?q=TOKYO,jp&units=metric&lang=ja&appid="
	response, err := http.Get(apiURL + apiKey)

	if err != nil {
		fmt.Println("HTTP GETリクエストエラー:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("レスポンスの読み取りエラー", err)
		return
	}

	var responseData ResponseData
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("JSONデコードエラー", err)
		return
	}

	fmt.Println("City Name:", responseData.City.Name)

	db, err := sql.Open("mysql", "go:go@tcp(mysql-db:3306)/playground")
	if nil != err {
		fmt.Println("データベース接続エラー:", err)
		return
	}
	defer db.Close()

	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("トランザクションの開始エラー:", err)
		return
	}
	defer func() {
		if err != nil {
			// トランザクション中にエラーが発生した場合、ロールバック
			err := tx.Rollback()
			if err != nil {
				return
			}
		} else {
			// エラーがない場合、コミット
			err := tx.Commit()
			if err != nil {
				return
			}
		}
	}()

	// city登録
	_, err = db.Exec("INSERT INTO city (id, country, name) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE id = VALUES(id), country = VALUES(country), name = VALUES(name)", responseData.City.Id, responseData.City.Country, responseData.City.Name)
	if err != nil {
		fmt.Println("データベースエラー:", err)
	}

	// city_detail登録
	_, err = db.Exec("INSERT INTO city_detail (city_id, lat, lon, population, timezone) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE city_id = VALUES(city_id), population = VALUES(population)", responseData.City.Id, responseData.City.Coord.Lat, responseData.City.Coord.Lon, responseData.City.Population, responseData.City.Timezone)
	if err != nil {
		fmt.Println("データベースエラー:", err)
	}

	for _, item := range responseData.List {
		fmt.Printf("日付: %s, 気温: %.2f\n", item.DtTxt, item.Main.Temp)
		// city_temperature登録
		_, err = db.Exec("INSERT INTO city_temperature (city_id, date, temperature, pressure, humidity) VALUES (?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE city_id = VALUES(city_id), date = VALUES(date), temperature = VALUES(temperature), pressure = VALUES(pressure), humidity = VALUES(humidity)", responseData.City.Id, item.DtTxt, item.Main.Temp, item.Main.Pressure, item.Main.Humidity)
		if err != nil {
			fmt.Println("データベースエラー:", err)
		}
	}
}
