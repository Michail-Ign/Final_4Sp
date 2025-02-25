package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, errors.New("Ошибка разделения строки")
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
}
