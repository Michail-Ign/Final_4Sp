package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, errors.New("Ошибка разделения строки.")
	}

	step, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, fmt.Errorf("[Parse int] %w", err)
	}

	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, fmt.Errorf("[Parse Duration] %w", err)
	}

	return step, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps == 0 {
		return ""
	}

	distanse_km := (StepLength * float64(steps)) / 1000
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	str1 := fmt.Sprintf("Количество шагов: %d.\n", steps)
	str2 := fmt.Sprintf("Дистанция составила %.2f км.\n", distanse_km)
	str3 := fmt.Sprintf("Вы сожгли %.2f ккал.", calories)

	return str1 + str2 + str3
}
