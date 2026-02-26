package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("не удалось получить количество шагов")
	}
	if weight <= 0 {
		return 0, errors.New("не удалось получить вес пользователя")
	}
	if height <= 0 {
		return 0, errors.New("не удалось получить рост пользователя")
	}
	if duration <= 0 {
		return 0, errors.New("не удалось получить время прогулки")
	}

	calories := (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("не удалось получить количество шагов")
	}
	if weight <= 0 {
		return 0, errors.New("не удалось получить вес пользователя")
	}
	if height <= 0 {
		return 0, errors.New("не удалось получить рост пользователя")
	}
	if duration <= 0 {
		return 0, errors.New("не удалось получить время бега")
	}

	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient * float64(steps)) / mInKm
}
