package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", 0, errors.New("Ошибка разделения строки")
	}

	step, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("[Parse int] %w", err)
	}
	str_activ := slice[1]
	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("[Parse Duration] %w", err)
	}

	return step, str_activ, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	return (lenStep * float64(steps)) / float64(mInKm)
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration == 0 {
		return 0
	}

	dist_km := distance(steps)

	return dist_km / duration.Hours()
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, str_activ, duration, err := parseTraining(data)
	if err != nil {
		return err.Error()
	}
	var ccal float64

	switch str_activ {
	case "Ходьба":
		ccal = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		ccal = RunningSpentCalories(steps, weight, duration)
	default:
		return "неизвестный тип тренировки"
	}
	dist_km := distance(steps)
	speed_avg := meanSpeed(steps, duration)

	str1 := fmt.Sprintf("Тип тренировки: %s.\n", str_activ)
	str2 := fmt.Sprintf("Длительность: %.2f ч.\n", duration.Hours())
	str3 := fmt.Sprintf("Дистанция: %.2f км.\n", dist_km)
	str4 := fmt.Sprintf("Скорость: %.2f км/ч\n", speed_avg)
	str5 := fmt.Sprintf("Сожгли калорий: %.2f", ccal)

	return str1 + str2 + str3 + str4 + str5
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	mean_speed := meanSpeed(steps, duration)

	return ((runningCaloriesMeanSpeedMultiplier * mean_speed) - runningCaloriesMeanSpeedShift) * weight
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	mean_speed := meanSpeed(steps, duration)

	return ((walkingCaloriesWeightMultiplier * weight) + (mean_speed*mean_speed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
}
