package asset

import (
	"database/sql"
	"sync/atomic"
)

// ChartData структура для данных графика
type ChartData struct {
	Labels   []string  `json:"labels"`
	Datasets []Dataset `json:"datasets"`
}

// Dataset структура для набора данных
type Dataset struct {
	Label           string   `json:"label"`
	Data            []int    `json:"data"`
	BackgroundColor []string `json:"backgroundColor"`
}

// Глобальный список цветов
var colors = []string{
	"rgba(201, 203, 207, 0.2)",
	"rgba(0, 255, 255, 0.2)",
	"rgba(255, 0, 255, 0.2)",
	"rgba(128, 0, 128, 0.2)",
	"rgba(255, 0, 0, 0.2)",
	"rgba(0, 255, 0, 0.2)",
	"rgba(0, 0, 255, 0.2)",
	"rgba(255, 255, 0, 0.2)",
	"rgba(0, 255, 127, 0.2)",
	"rgba(0, 0, 128, 0.2)",
	"rgba(255, 165, 0, 0.2)",
	"rgba(255, 20, 147, 0.2)",
	"rgba(75, 0, 130, 0.2)",
	"rgba(65, 105, 225, 0.2)",
	"rgba(70, 130, 180, 0.2)",
	"rgba(30, 144, 255, 0.2)",
	"rgba(0, 191, 255, 0.2)",
	"rgba(135, 206, 250, 0.2)",
	"rgba(70, 130, 180, 0.2)",
	"rgba(176, 224, 230, 0.2)",
	"rgba(95, 158, 160, 0.2)",
	"rgba(0, 128, 128, 0.2)",
	"rgba(32, 178, 170, 0.2)",
	"rgba(47, 79, 79, 0.2)",
	"rgba(0, 139, 139, 0.2)",
	"rgba(72, 209, 204, 0.2)",
	"rgba(175, 238, 238, 0.2)",
	"rgba(127, 255, 212, 0.2)",
	"rgba(64, 224, 208, 0.2)",
	"rgba(224, 255, 255, 0.2)",
}

// Атомарный счетчик для отслеживания текущего индекса цвета
var colorIndex int32 = -1

// Функция для выбора следующего цвета из списка
func getNextColor() string {
	// Атомарно увеличиваем счетчик и получаем следующий индекс
	nextIndex := atomic.AddInt32(&colorIndex, 1)
	// Обеспечиваем цикличность выбора цветов
	nextIndex = nextIndex % int32(len(colors))
	return colors[nextIndex]
}

// Функция для заполнения BackgroundColor в зависимости от количества элементов в Data
func fillBackgroundColor(data []int) []string {
	backgroundColors := make([]string, len(data))
	for i := range data {
		backgroundColors[i] = getNextColor()
	}
	return backgroundColors
}

// Функция для создания данных для первой диаграммы
func GetChartDataOne(db *sql.DB, userID int) (ChartData, error) {
	// SQL-запрос для получения типов активов и их процентного соотношения
	query := `
	WITH TotalAssetsValue AS (
		SELECT SUM(amount * price) AS total_value
		FROM assets
		WHERE user_id = $1
	)
	SELECT 
		vt.type,
		round(((SUM(a.amount * a.price) / (SELECT total_value FROM TotalAssetsValue)) * 100), 0) AS asset_percent
	FROM 
		assets a
	LEFT JOIN 
		v_tickers vt 
	ON 
		vt.ticker = a.ticker
	WHERE 
		a.user_id = $1
	GROUP BY 
		vt.type, (SELECT total_value FROM TotalAssetsValue)
	`

	// Выполнение запроса
	rows, err := db.Query(query, userID)
	if err != nil {
		return ChartData{}, err
	}
	defer rows.Close()

	var labels []string
	var data []int
	var backgroundColors []string

	// Обработка результатов запроса
	for rows.Next() {
		var assetType string
		var assetPercent int

		err := rows.Scan(&assetType, &assetPercent)
		if err != nil {
			return ChartData{}, err
		}

		labels = append(labels, assetType)
		data = append(data, assetPercent)
		backgroundColors = fillBackgroundColor(data)
	}

	// Проверка на ошибки при обработке строк
	if err = rows.Err(); err != nil {
		return ChartData{}, err
	}

	// Создание и возврат данных диаграммы
	return ChartData{
		Labels: labels,
		Datasets: []Dataset{
			{
				Label:           "Активы",
				Data:            data,
				BackgroundColor: backgroundColors,
			},
		},
	}, nil
}

// Функция для создания данных для второй диаграммы
func GetChartDataTwo(db *sql.DB, userID int) (ChartData, error) {
	// SQL-запрос для получения активов: Тикер и стоимость
	query := `
	SELECT 
    	a.ticker, 
    	ROUND(SUM(a.amount * a.price)) AS total_cost
	FROM 
    	assets a
	WHERE 
    	a.user_id = $1
	GROUP BY 
    	a.ticker
	`

	// Выполнение запроса
	rows, err := db.Query(query, userID)
	if err != nil {
		return ChartData{}, err
	}
	defer rows.Close()

	var labels []string
	var data []int
	var backgroundColors []string

	// Обработка результатов запроса
	for rows.Next() {
		var tickerName string
		var tickerCost int

		err := rows.Scan(&tickerName, &tickerCost)
		if err != nil {
			return ChartData{}, err
		}

		labels = append(labels, tickerName)
		data = append(data, tickerCost)
		backgroundColors = fillBackgroundColor(data)
	}

	// Проверка на ошибки при обработке строк
	if err = rows.Err(); err != nil {
		return ChartData{}, err
	}

	// Создание и возврат данных диаграммы
	return ChartData{
		Labels: labels,
		Datasets: []Dataset{
			{
				Label:           "Компании",
				Data:            data,
				BackgroundColor: backgroundColors,
			},
		},
	}, nil
}
