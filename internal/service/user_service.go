package service

import (
	"time"
)

func CalculateAge(dobStr string) (int, error) {
	dob, err := time.Parse("2006-01-02", dobStr)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	years := now.Year() - dob.Year()

	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		years--
	}

	return years, nil
}