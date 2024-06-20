package asset

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

// Функция для создания данных для первой диаграммы
func GetChartDataOne() ChartData {
	return ChartData{
		Labels: []string{"Red", "Blue", "Yellow", "Green", "Purple", "Orange"},
		Datasets: []Dataset{
			{
				Label: "My First Dataset",
				Data:  []int{300, 50, 100, 200, 150, 250},
				BackgroundColor: []string{
					"rgba(255, 99, 132, 0.2)",
					"rgba(54, 162, 235, 0.2)",
					"rgba(255, 206, 86, 0.2)",
					"rgba(75, 192, 192, 0.2)",
					"rgba(153, 102, 255, 0.2)",
					"rgba(255, 159, 64, 0.2)",
				},
			},
		},
	}
}

// Функция для создания данных для второй диаграммы
func GetChartDataTwo() ChartData {
	return ChartData{
		Labels: []string{"ETF", "Акции", "Облигации"},
		Datasets: []Dataset{
			{
				Label: "My Second Dataset",
				Data:  []int{30, 50, 20},
				BackgroundColor: []string{
					"rgba(54, 162, 235, 0.2)",
					"rgba(255, 206, 86, 0.2)",
					"rgba(75, 192, 192, 0.2)",
				},
			},
		},
	}
}
