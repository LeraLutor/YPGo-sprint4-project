package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// строка делится на слайс
	parts := strings.Split(data, ",")

	// проверка на длину
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Неверный формат данных. Ожидается 2 значения: количество шагов и длительность")
	}

	// колво шагов в int
	steps, err := strconv.Atoi(parts[0])
	// проверка на ошибку преобразования колво шагов
	if err != nil {
		return 0, 0, fmt.Errorf("Не удалось прочитать количество шагов: %w", err)
	}

	// проверка на ошибку данных
	if steps <= 0 {
		return 0, 0, fmt.Errorf("Количество шагов должно быть больше нуля")
	}

	// преобразование длительности
	duration, err := time.ParseDuration(parts[1])
	// проверка на ошибку преобразования длительности
	if err != nil {
		return 0, 0, fmt.Errorf("Не удалось прочитать длительность прогулки: %w", err)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
