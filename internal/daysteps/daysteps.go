package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	splittedData := strings.Split(datastring, ",")
	if len(splittedData) != 2 {
		return errors.New("не хватает элементов в слайсе")
	}

	steps, err := strconv.Atoi(splittedData[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return errors.New("не удалось получить количество шагов")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(splittedData[1])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if duration <= 0 {
		return errors.New("не удалось получить время тренировки")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("invalid calories format: %w", err)
	}

	result := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.`+"\n", ds.Steps, distance, calories)

	return result, nil
}
