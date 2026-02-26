package trainings

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	splittedData := strings.Split(datastring, ",")
	if len(splittedData) != 3 {
		return errors.New("не хватает элементов в слайсе")
	}

	steps, err := strconv.Atoi(splittedData[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return errors.New("не удалось получить количество шагов")
	}
	t.Steps = steps

	trainingType := splittedData[1]
	t.TrainingType = trainingType

	duration, err := time.ParseDuration(splittedData[2])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if duration <= 0 {
		return errors.New("не удалось получить время тренировки")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("invalid calories format: %w", err)
		}

		resultRun := fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`+"\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)

		return resultRun, nil

	case "Ходьба":

		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("invalid calories format: %w", err)
		}

		resultWalk := fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`+"\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories)

		return resultWalk, nil

	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
