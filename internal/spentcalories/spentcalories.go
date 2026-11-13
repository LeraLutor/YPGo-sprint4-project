package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

// парсинг тренировки
func parseTraining(data string) (int, string, time.Duration, error) {
	// строка делится на слайс
	parts := strings.Split(data, ",")

	// проверка на длину заданных данных
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("Неверный формат данных: ожидается количество шагов, тип тренировки и длительность")
	}

	// колво шагов в int
	steps, err := strconv.Atoi(parts[0])
	// проверка на ошибку данных
	if err != nil {
		return 0, "", 0, fmt.Errorf("Не удалось прочитать количество шагов: %w", err)
	}

	// преобразование длительности
	duration, err := time.ParseDuration(parts[2])
	// проверка на ошибку преобразования длительности
	if err != nil {
		return 0, "", 0, fmt.Errorf("Не удалось прочитать длительность тренировки: %w", err)
	}

	// приравнивание типа тренировки
	trainingType := parts[1]

	return steps, trainingType, duration, nil
}

// подсчет пройденной дистанции в км
func distance(steps int, height float64) float64 {
	// расчет длины шага
	stepLength := height * stepLengthCoefficient

	// подсчет длины в метрах
	distanceInMeters := float64(steps) * stepLength

	return distanceInMeters / mInKm
}

// подсчет средней скорости в км/ч
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// проверка на длительность тренировки
	if duration <= 0 {
		return 0
	}

	// подсчет дистанции
	distanceInKm := distance(steps, height)

	return distanceInKm / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

// расчет калоррий при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверка колва шагов
	if steps <= 0 {
		return 0, fmt.Errorf("Количество шагов должно быть больше нуля")
	}

	// проверка веса
	if weight <= 0 {
		return 0, fmt.Errorf("Вес должен быть больше нуля")
	}

	// проверка роста
	if height <= 0 {
		return 0, fmt.Errorf("Рост должен быть больше нуля")
	}

	// проверка длительности тренировки
	if duration <= 0 {
		return 0, fmt.Errorf("Длительность тренировки должна быть больше нуля")
	}

	// расчет средней скорости
	speed := meanSpeed(steps, height, duration)
	// длительность тренировки в минуты
	durationInMinutes := duration.Minutes()

	// подсчет калорий
	calories := weight * speed * durationInMinutes / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}
