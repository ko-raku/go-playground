# Goを使用して他のAPIを呼び出してレスポンスを取得
* 無料API「[OpenWeatherMap](https://openweathermap.org/price#weather)」を使用
## 事前準備
* 上記サイトでアカウント登録をして、APIキーの取得をしておく
* main.goのapiKeyを発行された自身のAPIキーに置き換える
```http request
// api call
https://api.openweathermap.org/data/2.5/forecast?q={city name}&units=metric&cnt={cnt}&appid={API key}

// response exsample
{
  "city": {
    "id": 2643743,
    "name": "London",
    "coord": {
      "lon": -0.1257,
      "lat": 51.5085
    },
    "country": "GB",
    "population": 1000000,
    "timezone": 3600
  },
  "cod": "200",
  "message": 0.0610271,
  "cnt": 7,
  "list": [
    {
      "dt": 1625140800,
      "sunrise": 1625111254,
      "sunset": 1625170853,
      "temp": {
        "day": 19.55,
        "min": 11.96,
        "max": 19.77,
        "night": 16.09,
        "eve": 19.47,
        "morn": 13.33
      },
      "feels_like": {
        "day": 19.15,
        "night": 15.43,
        "eve": 18.93,
        "morn": 12.7
      },
      "pressure": 1018,
      "humidity": 61,
      "weather": [
        {
          "id": 803,
          "main": "Clouds",
          "description": "broken clouds",
          "icon": "04d"
        }
      ],
      "speed": 2.71,
      "deg": 335,
      "gust": 3.8,
      "clouds": 69,
      "pop": 0.15
    },
    {
      "dt": 1625227200,
      "sunrise": 1625197696,
      "sunset": 1625257232,
      "temp": {
        "day": 21.51,
        "min": 14.46,
        "max": 22.73,
        "night": 16.56,
        "eve": 20.83,
        "morn": 15.21
      },
      "feels_like": {
        "day": 21.36,
        "night": 16.54,
        "eve": 20.69,
        "morn": 14.72
      },
      "pressure": 1016,
      "humidity": 63,
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10d"
        }
      ],
      "speed": 4.01,
      "deg": 216,
      "gust": 6.08,
      "clouds": 58,
      "pop": 0.77,
      "rain": 2.32
    },
    {
      "dt": 1625313600,
      "sunrise": 1625284141,
      "sunset": 1625343609,
      "temp": {
        "day": 17.31,
        "min": 15.75,
        "max": 21.21,
        "night": 16.17,
        "eve": 18.97,
        "morn": 15.81
      },
      "feels_like": {
        "day": 17.34,
        "night": 16.24,
        "eve": 19.09,
        "morn": 15.56
      },
      "pressure": 1014,
      "humidity": 86,
      "weather": [
        {
          "id": 500,
          "main": "Rain",
          "description": "light rain",
          "icon": "10d"
        }
      ],
      "speed": 3.84,
      "deg": 212,
      "gust": 7.93,
      "clouds": 100,
      "pop": 0.81,
      "rain": 3.83
    },
    ........
  ]
}
                   

                  
```

## 実行
```
% docker-compose up
[+] Running 1/0
 ⠿ Container go-dev-go-1  Created                                                                                                                                                                                                  0.0s
Attaching to go-dev-go-1
go-dev-go-1  | City Name: Tokyo
go-dev-go-1  | 日付: 2023-10-10 03:00:00, 気温: 23.87℃
go-dev-go-1  | 日付: 2023-10-10 06:00:00, 気温: 23.79℃
go-dev-go-1  | 日付: 2023-10-10 09:00:00, 気温: 22.89℃
go-dev-go-1  | 日付: 2023-10-10 12:00:00, 気温: 21.32℃
go-dev-go-1  | 日付: 2023-10-10 15:00:00, 気温: 20.45℃
go-dev-go-1 exited with code 0
```

